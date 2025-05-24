package client

import (
	"context"
	"scrapeit/pkg/common"
	"time"

	"github.com/josuebrunel/gopkg/xlog"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const (
	DiscoveryInterval = time.Hour
	DiscoveryTag      = "scrapeit"
)

type Client struct {
	Address   string
	TopicName string
	Node      *common.Node
}

func New(topic string) *Client {
	return &Client{
		TopicName: topic,
	}
}

func (c *Client) SetNode(node *common.Node) {
	c.Node = node
}

func (c Client) Run() error {
	ctx := context.Background()
	node := common.NewNode(ctx, c.TopicName, "/ip4/0.0.0.0/tcp/0")
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
		xlog.Info("go message", "from", msg.ReceivedFrom, "topic", msg.Topic, "data", string(msg.Data))
	}
}
