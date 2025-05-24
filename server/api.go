package server

import (
	"log"
	"scrapeit/pkg/common"
	"scrapeit/pkg/env"

	"github.com/josuebrunel/gopkg/xlog"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Api struct {
	ListenAddr string
	Publisher  *Server
}

type Request struct {
	Url string `json:"url"`
}

type Response struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

func NewApi() *Api {
	return &Api{
		ListenAddr: env.Get(env.AppAddrN),
		Publisher:  New(),
	}
}

func (api *Api) Start() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/v1/scrape", api.Scrape())
	xlog.Info("starting server", "addr", api.ListenAddr)
	log.Fatal(e.Start(api.ListenAddr))
	return nil
}

func (a Api) Scrape() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req Request
		if err := c.Bind(&req); err != nil {
			xlog.Error("failed to bind request", "error", err)
			return c.JSON(400, Response{Error: err.Error()})
		}
		xlog.Debug("received request", "url", req.Url)
		a.Publisher.Publish(c.Request().Context(), common.Message{Data: req.Url})
		return c.JSON(200, Response{Data: "received"})
	}
}
