package server

import (
	"context"
	"scrapeit/pkg/common"
	"scrapeit/pkg/env"
	"scrapeit/pkg/task"

	"github.com/josuebrunel/gopkg/xlog"
	"github.com/multiformats/go-multiaddr"
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
	node := common.NewNode(ctx, s.TopicName, s.Address, []multiaddr.Multiaddr{}, true)
	s.SetNode(node)
	return s
}

func (s Server) Publish(ctx context.Context, tsk *task.Task) {
	xlog.Info("publishing task", "task", *tsk, "topic", s.Node.TopicName)
	s.Node.Topic.Publish(ctx, tsk.JSONMarshal())
	xlog.Info("task published")
}

func (s *Server) SetNode(node *common.Node) {
	s.Node = node
}
