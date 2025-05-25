package common

import (
	"context"
	"sync"
	"time"

	"github.com/josuebrunel/gopkg/xlog"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	routingdiscovery "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	discoveryutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/multiformats/go-multiaddr"
)

func NewDHT(ctx context.Context, host host.Host, bootstrappedPeers []multiaddr.Multiaddr, serverMode bool) (*dht.IpfsDHT, error) {
	var options []dht.Option

	if serverMode {
		options = append(options, dht.Mode(dht.ModeAutoServer))
	}

	khdt, err := dht.New(ctx, host, options...)
	if err != nil {
		return nil, err
	}

	if err := khdt.Bootstrap(ctx); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	for _, peerAddr := range bootstrappedPeers {
		peerInfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerInfo); err != nil {
				xlog.Error("Error while connecting to bootstrap node:", "peer", peerInfo.ID, "error", err)
			} else {
				xlog.Info("Connection established with bootstrap node")
			}
		}()
	}
	wg.Wait()

	return khdt, nil
}

func DHTDiscover(ctx context.Context, h host.Host, dht *dht.IpfsDHT, rendezvous string) {
	var routingDiscover = routingdiscovery.NewRoutingDiscovery(dht)
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C: // Periodically discover peers
			peers, err := discoveryutil.FindPeers(ctx, routingDiscover, rendezvous)
			if err != nil {
				xlog.Error("Error while discovering peers:", err)
			}
			for _, p := range peers {
				if p.ID == h.ID() {
					continue
				}
				xlog.Info("Found peer:", p)
				if h.Network().Connectedness(p.ID) != network.Connected {
					_, err := h.Network().DialPeer(ctx, p.ID)
					xlog.Info("Connected to peer", "peer", p.ID, "err", err)
					if err != nil {
						continue
					}
				}
			}
		}
	}
}
