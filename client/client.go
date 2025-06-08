package client

import (
	"context"
	"scrapeit/pkg/common"
	"scrapeit/pkg/env"
	"scrapeit/pkg/etherutil"
	"scrapeit/scraper"
	"scrapeit/task"

	"github.com/josuebrunel/gopkg/xlog"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type Client struct {
	TopicName  string
	Node       *common.Node
	Manager    *task.EtherManager
	PrivateKey string
	Settings   env.Settings
}

func New() *Client {
	settings := env.NewSetting()
	xlog.Debug("settings from root peer", "settings", settings)
	return &Client{
		TopicName:  settings.Peer.Topic,
		Manager:    task.NewEtherManager(settings.Ethereum.Network, settings.Ethereum.Contract),
		PrivateKey: env.Get(env.EtherPrivateKeyN),
		Settings:   settings,
	}
}

func (c *Client) SetNode(node *common.Node) {
	c.Node = node
}

func (c Client) Run() error {
	ctx := context.Background()
	peersRoot := common.MultiAddrFromStrings(c.Settings.Peer.Addrs)
	node := common.NewNode(ctx, c.TopicName, "/ip4/0.0.0.0/tcp/0", peersRoot, false)
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
	pk := etherutil.MustPrivateKeyFromHex(c.PrivateKey)
	ipfsCli, err := common.NewIPFSClient(c.Settings.Peer.Ipfs)
	if err != nil {
		panic(err)
	}
	xlog.Info("subscribing to topic", "topic", c.TopicName)
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
		err = c.Manager.Claim(ctx, tsk.ID, pk)
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
			tsk.Output, err = ipfsCli.UploadFile(ctx, content)
			if err != nil {
				xlog.Error("Failed to upload file", "err", err, "task", tsk.ID)
				scrpErr = err
			}
			xlog.Debug("Task result uploaded", "task", tsk.ID, "cid", tsk.Output)
		}
		if scrpErr != nil {
			if err := c.Manager.Fail(ctx, tsk.ID, pk); err != nil {
				xlog.Error("Failed to fail task", "err", err, "task", tsk.ID)
				continue
			}
		}
		if err := c.Manager.Complete(ctx, tsk.ID, tsk.Output, pk); err != nil {
			xlog.Error("Failed to complete task", "err", err, "task", tsk.ID)
			continue
		}
		xlog.Info("Task completed", "task", tsk.ID)
	}
}
