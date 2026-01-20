package app

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"log"
	"net/http"
	"scrapeit/account"
	"scrapeit/pkg/env"
	"scrapeit/pkg/etherutil"
	"scrapeit/pkg/view"
	"scrapeit/server"
	"scrapeit/task"

	_ "github.com/lib/pq"
	"github.com/stephenafamo/bob"
	"github.com/supabase-community/auth-go"

	"github.com/josuebrunel/gopkg/xlog"
	"github.com/josuebrunel/gopkg/xsession"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type App struct {
	ListenAddr  string
	Publisher   *server.Server
	Manager     *task.EtherManager
	privateKey  *ecdsa.PrivateKey
	AuthBackend string
	DatabaseURL string
	db          *sql.DB
	bob         bob.DB
}

type Request struct {
	Url        string `json:"url"`
	PrivateKey string `json:"private_key"`
}

type Response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}

func NewApp() *App {
	db := getDBConnection(env.Get(env.AppDatabaseN))
	bdb := bob.NewDB(db)
	m := task.NewEtherManager(
		env.Get(env.EtherAddrN),
		env.Get(env.EtherContractAddrN),
	)
	m.SetDB(bdb)
	return &App{
		db:          db,
		bob:         bdb,
		AuthBackend: env.Get(env.AppAuthBackendN),
		DatabaseURL: env.Get(env.AppDatabaseN),
		Publisher:   server.New(*m),
		privateKey:  etherutil.MustPrivateKeyFromHex(env.Get(env.EtherPrivateKeyN)),
		Manager:     m,
		ListenAddr:  env.Get(env.AppAddrN),
	}
}

func getDBConnection(dsn string) *sql.DB {
	xlog.Debug("connecting to database", "dsn", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)

	}
	return db
}

func (a *App) Close() {
	a.db.Close()
	a.bob.Close()
	a.Manager.Close()
}

func (a *App) Start() error {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(xsession.LoadAndSave(xsession.SessionManager))
	e.Static("/static", "static")
	e.GET("/", func(c *echo.Context) error { return view.Render(c, http.StatusOK, Index(), nil) })
	// TaskV1
	taskV1 := e.Group("/api/v1/task", ApiKeyAuthMiddleware(a.bob))
	taskV1.POST("", a.Scrape())
	taskV1.PATCH("/:id/claim", a.Claim())
	taskV1.PATCH("/:id/complete", a.Complete())
	taskV1.GET("/:id", a.Result())
	e.GET("/settings", a.Settings())
	// get Auth backend client
	authCli := auth.New("scrapeit", "").WithCustomAuthURL(a.AuthBackend)
	ethCli := a.Manager.GetClient()
	// mount account
	xlog.Info("mounting account")
	account.Mount(e, a.db, authCli, &ethCli)
	e.GET("/account/apikey", a.GenerateApiKey(), xsession.LoginRequired("/account/signin"))
	e.PUT("/account/apikey", a.GenerateApiKey(), xsession.LoginRequired("/account/signin"))
	// mount task
	xlog.Info("mounting task")
	task.Mount(context.Background(), e, a.bob)
	xlog.Info("starting server", "addr", a.ListenAddr)
	go a.Publisher.ListenContractEvents(context.Background())
	log.Fatal(e.Start(a.ListenAddr))
	return nil
}
