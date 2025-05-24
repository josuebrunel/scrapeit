package server

import (
	"bufio"
	"context"
	"os"
	"scrapeit/pkg/common"

	"github.com/josuebrunel/gopkg/xlog"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type Server struct {
	TopicName string
	Node      *common.Node
}

func New(topicName string) *Server {
	return &Server{TopicName: topicName}
}

func (s Server) Publish(ctx context.Context, topic *pubsub.Topic) {
	xlog.Info("enter message")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := scanner.Text()
			xlog.Info("publishing message", "msg", msg)
			if len(msg) != 0 {
				topic.Publish(ctx, []byte(msg))
			}

		}
	}
}

func (s *Server) SetNode(node *common.Node) {
	s.Node = node
}

func (s Server) Run() error {
	ctx := context.Background()
	node := common.NewNode(ctx, s.TopicName, "/ip4/0.0.0.0/tcp/7777")
	s.SetNode(node)
	defer s.Node.Host.Close()
	s.Publish(ctx, node.Topic)
	return nil
}
