package server

import (
	"log"
	"scrapeit/pkg/env"
	"scrapeit/pkg/task"

	"github.com/google/uuid"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Api struct {
	ListenAddr string
	Publisher  *Server
	Manager    *task.Manager
}

type Request struct {
	Url string `json:"url"`
}

type Response struct {
	Task  task.Task `json:"task"`
	Error string    `json:"error"`
}

func NewApi() *Api {
	return &Api{
		ListenAddr: env.Get(env.AppAddrN),
		Publisher:  New(),
		Manager:    task.NewManager(),
	}
}

func (api *Api) Start() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/v1/scrape", api.Scrape())
	e.GET("/v1/result/:id", api.Result())
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
		tskID := uuid.New().String()
		tsk := task.NewTask(tskID, c.Request().RemoteAddr, req.Url)
		if err := a.Manager.Create(c.Request().Context(), tsk.ID, tsk); err != nil {
			xlog.Error("failed to create task", "error", err)
			return c.JSON(500, Response{Error: err.Error()})
		}
		a.Publisher.Publish(c.Request().Context(), &tsk)
		return c.JSON(200, Response{Task: tsk})
	}
}

func (a Api) Result() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")
		tsk, err := a.Manager.Get(c.Request().Context(), id)
		if err != nil {
			xlog.Error("failed to get task", "error", err)
			return c.JSON(500, Response{Error: err.Error()})
		}
		if tsk.Status != task.Completed {
			return c.JSON(200, Response{Task: tsk})
		}
		return c.JSON(200, Response{Task: tsk})
	}
}
