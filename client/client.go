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
	Manager    Manager
	PrivateKey string
	Settings   env.Settings
}

func New() *Client {
	settings := env.NewSetting()
	xlog.Debug("settings from root peer", "settings", settings)
	return &Client{
		TopicName: settings.Peer.Topic,
		Manager: NewManager(
			env.Get(env.PubPeersRoot),
			etherutil.GetAddressFromHexKey(env.Get(env.EtherPrivateKeyN)),
			env.Get(env.ApiKey),
		),
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
		if err = c.Manager.Claim(ctx, tsk); err != nil {
			xlog.Error("Failed to claim task", "err", err, "task", tsk.ID)
			continue
		}
		xlog.Info("Task claimed", "task", tsk.ID, "assignee", c.Node.ID.String())
		content, err := scraper.Scrape(ctx, tsk.Input)
		if err != nil {
			xlog.Error("Failed to scrape", "err", err, "task", tsk.ID)
			tsk.Output = err.Error()
		} else {
			tsk.Output, err = ipfsCli.UploadFile(ctx, content)
			if err != nil {
				xlog.Error("Failed to upload file", "err", err, "task", tsk.ID)
				err = err
			}
			xlog.Debug("Task result uploaded", "task", tsk.ID, "cid", tsk.Output)
		}
		if err := c.Manager.Complete(ctx, tsk, err); err != nil {
			xlog.Error("Failed to complete task", "err", err, "task", tsk.ID)
			continue
		}
		xlog.Info("Task completed", "task", tsk.ID)
	}
}
