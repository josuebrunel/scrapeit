package task

import (
	"context"
	"net/http"
	"scrapeit/pkg/models"
	"scrapeit/pkg/view"

	"github.com/gofrs/uuid/v5"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/josuebrunel/gopkg/xsession"
	"github.com/labstack/echo/v5"
	"github.com/stephenafamo/bob"
)

const (
	HandlerNameList   = "task.list"
	HandlerNameDetail = "task.detail"
)

type Handler struct {
	db bob.DB
}

func NewHandler(db bob.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func Mount(ctx context.Context, e *echo.Echo, db bob.DB) {
	handler := NewHandler(db)
	g := e.Group("/task", xsession.LoginRequired("/account/signin"))
	g.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "",
		Name:    HandlerNameList,
		Handler: handler.List(),
	})
	g.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/:id",
		Name:    HandlerNameDetail,
		Handler: handler.Detail(),
	})
}

func (h Handler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx  = c.Request().Context()
			user = xsession.GetUser(ctx)
		)
		tt, err := QueryListByUser(ctx, h.db, user.ID)
		if err != nil {
			xlog.Error("failed to query tasks", "error", err)
			return err
		}
		return view.Render(c, http.StatusOK, ListView(tt), nil)
	}
}

func (h Handler) Detail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx = c.Request().Context()
			id  = uuid.Must(uuid.FromString(c.PathParam("id")))
		)
		t, err := models.FindTask(ctx, h.db, id)
		if err != nil {
			xlog.Error("failed to find task", "error", err, "id", id)
			return err
		}
		return view.Render(c, http.StatusOK, DetailView(*t), nil)
	}
}
