package app

import (
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"scrapeit/account"
	"scrapeit/pkg/common"
	"scrapeit/pkg/env"
	"scrapeit/pkg/etherutil"
	"scrapeit/pkg/task"
	"scrapeit/pkg/view"
	"scrapeit/server"

	_ "github.com/lib/pq"
	"github.com/stephenafamo/bob"
	"github.com/supabase-community/auth-go"

	"github.com/google/uuid"
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
	Database    string
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
	return &App{
		AuthBackend: env.Get(env.AppAuthBackendN),
		Database:    env.Get(env.AppDatabaseN),
		Publisher:   server.New(),
		privateKey:  etherutil.MustPrivateKeyFromHex(env.Get(env.EtherPrivateKeyN)),
		Manager: task.NewEtherManager(
			env.Get(env.EtherAddrN),
			env.Get(env.EtherContractAddrN),
		),
		ListenAddr: env.Get(env.AppAddrN),
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

func (a *App) Start() error {
	db := getDBConnection(a.Database)
	defer db.Close()
	a.bob = bob.NewDB(db)
	defer a.bob.Close()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(xsession.LoadAndSave(xsession.SessionManager))
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error { return view.Render(c, http.StatusOK, Index(), nil) })
	e.POST("/v1/task", a.Scrape())
	e.GET("/v1/task/:id", a.Result())
	e.GET("/settings", a.Settings())
	// get Auth backend client
	authCli := auth.New("scrapeit", "").WithCustomAuthURL(a.AuthBackend)
	// mount account
	account.Mount(e, db, authCli)
	e.GET("/account/apikey", a.GenerateApiKey(), xsession.LoginRequired("/account/signin"))
	e.PUT("/account/apikey", a.GenerateApiKey(), xsession.LoginRequired("/account/signin"))
	xlog.Info("starting server", "addr", a.ListenAddr)
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
		pk := etherutil.MustPrivateKeyFromHex(req.PrivateKey)
		tskID := uuid.New().String()
		tsk := task.NewTask(tskID, etherutil.GetAddress(pk), req.Url)
		if err := a.Manager.Create(ctx, pk, tsk); err != nil {
			xlog.Error("failed to create task", "error", err)
			return c.JSON(500, Response[task.Task]{Error: err.Error()})
		}
		a.Publisher.Publish(ctx, &tsk)
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
