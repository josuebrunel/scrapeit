package server

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"scrapeit/pkg/common"
	"scrapeit/pkg/env"
	"scrapeit/pkg/etherutil"
	"scrapeit/pkg/task"

	"github.com/google/uuid"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Api struct {
	ListenAddr string
	Publisher  *Server
	Manager    *task.EtherManager
	privateKey *ecdsa.PrivateKey
}

type Request struct {
	Url        string `json:"url"`
	PrivateKey string `json:"private_key"`
}

type Response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}

func NewApi() *Api {
	return &Api{
		ListenAddr: env.Get(env.AppAddrN),
		Publisher:  New(),
		Manager: task.NewEtherManager(
			env.Get(env.EtherAddrN),
			env.Get(env.EtherContractAddrN),
		),
		privateKey: etherutil.MustPrivateKeyFromHex(env.Get(env.EtherPrivateKeyN)),
	}
}

func (api *Api) Start() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/v1/task", api.Scrape())
	e.GET("/v1/task/:id", api.Result())
	e.GET("/settings", api.Settings())
	xlog.Info("starting server", "addr", api.ListenAddr)
	log.Fatal(e.Start(api.ListenAddr))
	return nil
}

func (a Api) Scrape() echo.HandlerFunc {
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

func (a Api) Result() echo.HandlerFunc {
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

func (a Api) Settings() echo.HandlerFunc {
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
