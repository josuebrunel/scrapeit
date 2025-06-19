package client

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/url"
	"scrapeit/task"

	"github.com/davesavic/clink"
	"github.com/josuebrunel/gopkg/xlog"
)

const (
	URITask         = "/api/v1/task"
	URITaskClaim    = "/claim"
	URITaskComplete = "/complete"
)

var (
	ErrTaskClaim    = errors.New("failed to claim task")
	ErrTaskComplete = errors.New("failed to complete task")
)

type Manager struct {
	BaseURL string
	Address string
	ApiKey  string
	client  *clink.Client
}

func NewManager(url, address, apiKey string) Manager {
	cli := clink.NewClient(clink.WithHeader("Content-Type", "application/json"), clink.WithHeader("Authorization", apiKey))
	return Manager{
		BaseURL: url,
		Address: address,
		ApiKey:  apiKey,
		client:  cli,
	}
}

func (m Manager) GetEndpoint(id ...string) string {
	url, err := url.JoinPath(m.BaseURL, append([]string{URITask}, id...)...)
	if err != nil {
		xlog.Error("Error while getting endpoint", "error", err)
		return ""
	}
	return url
}

func (m Manager) Update(ctx context.Context, url string, tsk task.Task) error {
	tsk.Assignee = m.Address
	payload := bytes.NewBuffer(tsk.JSONMarshal())
	resp, err := m.client.Patch(url, payload)
	xlog.Info("Updating task", "url", url, "task", tsk)
	if err != nil {
		xlog.Error("Error while updating task", "task", tsk, "error", err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		xlog.Error("Error while updating task", "task", tsk, "status", resp.StatusCode)
		return ErrTaskClaim
	}
	return nil
}

func (m Manager) Claim(ctx context.Context, tsk task.Task) error {
	tsk.Status = task.Active
	return m.Update(ctx, m.GetEndpoint(tsk.ID, URITaskClaim), tsk)
}

func (m Manager) Complete(ctx context.Context, tsk task.Task, err error) error {
	tsk.Status = task.Completed
	if err != nil {
		tsk.Status = task.Failed
		tsk.Output = err.Error()
	}
	return m.Update(ctx, m.GetEndpoint(tsk.ID, URITaskComplete), tsk)
}
