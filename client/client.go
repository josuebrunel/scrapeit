package client

import (
	"context"
	"scrapeit/pkg/common"
	"scrapeit/pkg/env"
	"scrapeit/pkg/task"
	"scrapeit/scraper"

	"github.com/josuebrunel/gopkg/xlog"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type Client struct {
	Address   string
	TopicName string
	Node      *common.Node
	Manager   *task.Manager
}

func New() *Client {
	return &Client{
		Address:   env.Get(env.PubAddrN),
		TopicName: env.Get(env.PubTopicN),
		Manager:   task.NewManager(),
	}
}

func (c *Client) SetNode(node *common.Node) {
	c.Node = node
}

func (c Client) Run() error {
	ctx := context.Background()
	node := common.NewNode(ctx, c.TopicName, "/ip4/0.0.0.0/tcp/0", common.PeersRoot(), false)
	c.SetNode(node)
	defer node.Host.Close()

	subscriber, err := node.Topic.Subscribe()
	if err != nil {
		panic(err)
	}
	c.Subscribe(ctx, subscriber)
	return nil
}

func (c Client) Subscribe(ctx context.Context, subscriber *pubsub.Subscription) {
	for {
		msg, err := subscriber.Next(ctx)
		if err != nil {
			panic(err)
		}
		if msg.ReceivedFrom == c.Node.ID {
			continue
		}
		var tsk task.Task
		tsk.JSONUnmarshal(msg.Data)
		xlog.Info("go message", "from", msg.ReceivedFrom, "topic", msg.Topic, "task", tsk)
		err = c.Manager.Claim(ctx, tsk.ID, c.Node.ID.String())
		if err != nil {
			xlog.Error("Failed to claim task", "err", err, "task", tsk.ID)
			continue
		}
		xlog.Info("Task claimed", "task", tsk.ID, "assignee", c.Node.ID.String())
		content, scrpErr := scraper.Scrape(ctx, tsk.Input)
		if scrpErr != nil {
			xlog.Error("Failed to scrape", "err", scrpErr, "task", tsk.ID)
			tsk.Output = scrpErr.Error()
		} else {
			tsk.Output, err = common.IPFSClient.UploadFile(ctx, content)
			if err != nil {
				xlog.Error("Failed to upload file", "err", err, "task", tsk.ID)
				scrpErr = err
			}
			xlog.Debug("Task result uploaded", "task", tsk.ID, "cid", tsk.Output)
		}

		if err := c.Manager.Complete(ctx, tsk.ID, tsk.Output, scrpErr); err != nil {
			xlog.Error("Failed to complete task", "err", err, "task", tsk.ID)
			continue
		}
		xlog.Info("Task completed", "task", tsk.ID)
	}
}
