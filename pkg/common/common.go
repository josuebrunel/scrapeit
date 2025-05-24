package common

import (
	"context"
	"time"

	"github.com/josuebrunel/gopkg/xlog"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/multiformats/go-multiaddr"
)

const (
	DiscoveryInterval = time.Hour
	DiscoveryTag      = "scrapeit"
)

type DiscoveryNotifee struct {
	h host.Host
}

func NewDiscoveryNotifee(h host.Host) *DiscoveryNotifee {
	return &DiscoveryNotifee{
		h: h,
	}
}

func (d *DiscoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	xlog.Info("discovered new peer", "peer", pi.ID)
	if err := d.h.Connect(context.Background(), pi); err != nil {
		xlog.Error("error while connecting to peer", "err", err)
	}
}

func SetupDiscover(h host.Host) error {
	m := mdns.NewMdnsService(h, DiscoveryTag, NewDiscoveryNotifee(h))
	return m.Start()
}

type Node struct {
	ID        peer.ID
	Address   string
	Host      host.Host
	Topic     *pubsub.Topic
	TopicName string
}

func NewNode(ctx context.Context, topicName, address string, bootStrappingPeers []multiaddr.Multiaddr) *Node {
	host, err := libp2p.New(libp2p.ListenAddrStrings(address))
	if err != nil {
		xlog.Error("error while creating host", "err", err)
		panic(err)
	}
	xlog.Info("host created", "host", host.ID())
	for id, add := range host.Addrs() {
		xlog.Info("host address", "no", id, "addr", add)
	}
	xlog.Info("creating pubsub service")
	gossip, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		panic(err)
	}
	// dht discovery
	dht, err := NewDHT(ctx, host, bootStrappingPeers)
	if err != nil {
		panic(err)
	}
	go DHTDiscover(ctx, host, dht, DiscoveryTag)
	// local discovery
	xlog.Info("setup local discovery")
	if err := SetupDiscover(host); err != nil {
		panic(err)
	}
	topic, err := gossip.Join(topicName)
	if err != nil {
		panic(err)
	}
	return &Node{
		ID:        host.ID(),
		Address:   address,
		Host:      host,
		Topic:     topic,
		TopicName: topicName,
	}
}
