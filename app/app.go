package app

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"scrapeit/account"
	"scrapeit/pkg/common"
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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(xsession.LoadAndSave(xsession.SessionManager))
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error { return view.Render(c, http.StatusOK, Index(), nil) })
	e.POST("/v1/task", a.Scrape(), ApiKeyAuthMiddleware(a.bob))
	e.GET("/v1/task/:id", a.Result(), ApiKeyAuthMiddleware(a.bob))
	e.GET("/settings", a.Settings())
	// get Auth backend client
	authCli := auth.New("scrapeit", "").WithCustomAuthURL(a.AuthBackend)
	// mount account
	account.Mount(e, a.db, authCli)
	e.GET("/account/apikey", a.GenerateApiKey(), xsession.LoginRequired("/account/signin"))
	e.PUT("/account/apikey", a.GenerateApiKey(), xsession.LoginRequired("/account/signin"))
	xlog.Info("starting server", "addr", a.ListenAddr)
	go a.Publisher.ListenContractEvents(context.Background())
	log.Fatal(e.Start(a.ListenAddr))
	return nil
}

func (a App) Scrape() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			req Request
			ctx = c.Request().Context()
		)

		if err := c.Bind(&req); err != nil {
			xlog.Error("failed to bind request", "error", err)
			return c.JSON(400, Response[task.Task]{Error: err.Error()})
		}
		xlog.Debug("received request", "url", req.Url)
		var (
			pk       = etherutil.MustPrivateKeyFromHex(req.PrivateKey)
			userID   = c.Get(RequestContextUserID).(string)
			contract = env.Get(env.EtherContractAddrN)
			tsk      = task.NewTask(userID, etherutil.GetAddress(pk), req.Url)
		)
		tsk.Contract = contract
		if err := a.Manager.Upsert(ctx, tsk); err != nil {
			xlog.Error("failed to upsert task", "error", err)
			return c.JSON(500, Response[task.Task]{Error: err.Error()})
		}
		if err := a.Manager.Create(ctx, pk, tsk); err != nil {
			xlog.Error("failed to create task", "error", err)
			return c.JSON(500, Response[task.Task]{Error: err.Error()})
		}
		return c.JSON(200, Response[task.Task]{Data: tsk})
	}
}

func (a App) Result() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx = c.Request().Context()
			id  = c.PathParam("id")
		)
		tsk, err := a.Manager.Get(ctx, id, a.privateKey)
		if err != nil {
			xlog.Error("failed to get task", "error", err)
			return c.JSON(500, Response[task.Task]{Error: err.Error()})
		}
		if tsk.Status != task.Completed {
			return c.JSON(200, Response[task.Task]{Data: tsk})
		}
		ipfsCli, err := common.NewIPFSClient(env.Get(env.IpfsAddrN))
		if err != nil {
			xlog.Error("failed to create ipfs client", "error", err)
			return c.JSON(500, Response[task.Task]{Error: err.Error()})
		}
		content, err := ipfsCli.DownloadFile(ctx, tsk.Output)
		if err != nil {
			xlog.Error("failed to download file", "error", err)
			return c.JSON(500, Response[task.Task]{Error: err.Error()})
		}
		tsk.Output = content
		return c.JSON(200, Response[task.Task]{Data: tsk})
	}
}

func (a App) Settings() echo.HandlerFunc {
	return func(c echo.Context) error {
		peerID := a.Publisher.Node.ID.String()
		addr := func() []string {
			var r = []string{}
			for _, a := range a.Publisher.Node.Host.Addrs() {
				r = append(r, fmt.Sprintf("%s/p2p/%s", a.String(), peerID))
			}
			return r
		}
		return c.JSON(200, Response[env.Settings]{Data: env.Settings{
			Peer: env.Peer{
				ID:    peerID,
				Addrs: addr(),
				Ipfs:  env.Get(env.IpfsAddrN),
				Topic: env.Get(env.PubTopicN),
			},
			Ethereum: env.Ethereum{
				Network:  env.Get(env.EtherAddrN),
				Contract: env.Get(env.EtherContractAddrN),
			},
		}})
	}
}
