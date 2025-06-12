package app

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"scrapeit/pkg/common"
	"scrapeit/pkg/env"
	"scrapeit/pkg/models"
	"scrapeit/task"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/labstack/echo/v5"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
)

const (
	Status200 = http.StatusOK
	Status400 = http.StatusBadRequest
	Status500 = http.StatusInternalServerError
)

func (a App) Scrape() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			req Request
			ctx = c.Request().Context()
		)

		if err := c.Bind(&req); err != nil {
			xlog.Error("failed to bind request", "error", err)
			return c.JSON(http.StatusBadRequest, Response[task.Task]{Error: err.Error()})
		}
		xlog.Debug("received request", "url", req.Url)
		var (
			userID   = c.Get(RequestContextUserID).(string)
			ethAddr  = c.Get(RequestContextEthAddr).(string)
			contract = env.Get(env.EtherContractAddrN)
			tsk      = task.NewTask(userID, ethAddr, req.Url)
		)
		tsk.Contract = contract
		tsk.Reward = task.RewardDefault
		if err := a.Manager.Upsert(ctx, tsk); err != nil {
			xlog.Error("failed to upsert task", "error", err)
			return c.JSON(Status500, Response[task.Task]{Error: err.Error()})
		}
		if err := a.Manager.Create(ctx, a.privateKey, tsk); err != nil {
			xlog.Error("failed to create task", "error", err)
			return c.JSON(Status500, Response[task.Task]{Error: err.Error()})
		}
		a.Publisher.Publish(ctx, &tsk)
		return c.JSON(http.StatusOK, Response[task.Task]{Data: tsk})
	}
}

func (a App) Submit(filters ...bob.Expression) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx     = c.Request().Context()
			id      = c.PathParam("id")
			reqTask task.Task
		)
		if err := c.Bind(&reqTask); err != nil {
			xlog.Error("failed to bind request", "error", err)
			return c.JSON(http.StatusBadRequest, Response[task.Task]{Error: err.Error()})
		}
		reqTask.UpdatedAt = time.Now().UTC()
		tsk, err := models.FindTask(ctx, a.bob, uuid.Must(uuid.FromString(id)))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				xlog.Error("task not found", "id", id)
				return c.JSON(http.StatusNotFound, Response[task.Task]{Error: "task not found"})
			}
			return c.JSON(http.StatusInternalServerError, Response[task.Task]{Error: err.Error()})
		}
		switch task.TaskStatus(tsk.Status.V) {
		case task.Pending:
			if reqTask.Status != task.Active {
				xlog.Error("invalid status", "status", reqTask.Status)
				return c.JSON(http.StatusBadRequest, Response[task.Task]{Error: "invalid status"})
			}
		case task.Active:
			if reqTask.Status != task.Completed && reqTask.Status != task.Failed {
				xlog.Error("invalid status", "status", reqTask.Status)
				return c.JSON(http.StatusBadRequest, Response[task.Task]{Error: "invalid status"})
			}
		}
		if err := a.Manager.Upsert(ctx, reqTask); err != nil {
			xlog.Error("failed to update task", "error", err)
			return c.JSON(Status500, Response[task.Task]{Error: err.Error()})
		}
		if err := a.Manager.Complete(ctx, id, tsk.Assignee.V, reqTask.Output, reqTask.Status == task.Completed, a.privateKey); err != nil {
			xlog.Error("failed to complete task", "error", err)
			return c.JSON(Status500, Response[task.Task]{Error: err.Error()})
		}
		return c.JSON(http.StatusOK, Response[task.Task]{Data: reqTask})
	}
}

func (a App) Claim() echo.HandlerFunc {
	return a.Submit(psql.Quote("status").EQ(psql.Arg(task.Pending)))
}

func (a App) Complete() echo.HandlerFunc {
	return a.Submit(psql.Quote("status").EQ(psql.Arg(task.Active)))
}

func (a App) Result() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx = c.Request().Context()
			id  = c.PathParam("id")
		)
		tsk, err := a.Manager.DBGet(ctx, id)
		if err != nil {
			xlog.Error("failed to get task", "error", err)
			return c.JSON(Status500, Response[task.Task]{Error: err.Error()})
		}
		if tsk.Status != task.Completed {
			return c.JSON(http.StatusOK, Response[task.Task]{Data: tsk})
		}
		ipfsCli, err := common.NewIPFSClient(env.Get(env.IpfsAddrN))
		if err != nil {
			xlog.Error("failed to create ipfs client", "error", err)
			return c.JSON(Status500, Response[task.Task]{Error: err.Error()})
		}
		content, err := ipfsCli.DownloadFile(ctx, tsk.Output)
		if err != nil {
			xlog.Error("failed to download file", "error", err)
			return c.JSON(Status500, Response[task.Task]{Error: err.Error()})
		}
		tsk.Output = content
		return c.JSON(http.StatusOK, Response[task.Task]{Data: tsk})
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
		return c.JSON(http.StatusOK, Response[env.Settings]{Data: env.Settings{
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
