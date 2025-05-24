package server

import (
	"context"
	"scrapeit/pkg/common"
	"scrapeit/pkg/env"

	"github.com/josuebrunel/gopkg/xlog"
)

type Server struct {
	Address   string
	TopicName string
	Node      *common.Node
}

func New() *Server {
	s := &Server{
		Address:   env.Get(env.PubAddrN),
		TopicName: env.Get(env.PubTopicN),
	}
	ctx := context.Background()
	node := common.NewNode(ctx, s.TopicName, s.Address, common.PeersRootNil)
	s.SetNode(node)
	return s
}

func (s Server) Publish(ctx context.Context, msg common.Message) {
	xlog.Info("enter message")
	xlog.Info("publishing message", "msg", msg)
	s.Node.Topic.Publish(ctx, msg.JSONMarshal())
}

func (s *Server) SetNode(node *common.Node) {
	s.Node = node
}
