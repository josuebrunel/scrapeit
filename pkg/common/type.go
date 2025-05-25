package common

import (
	"scrapeit/pkg/env"
	"strings"

	"github.com/josuebrunel/gopkg/xlog"
	"github.com/multiformats/go-multiaddr"
)

const (
	ActionScrape     = "scrape"
	ActionScreenshot = "screenshot"
)

var (
	PeersRootNil = []multiaddr.Multiaddr{}
	PeersRoot    = func() []multiaddr.Multiaddr {
		var peers []multiaddr.Multiaddr
		peersEnv := env.Get(env.PubPeersRoot)
		if peersEnv == "" {
			return PeersRootNil
		}
		for _, peer := range strings.Split(peersEnv, ",") {
			addr, err := multiaddr.NewMultiaddr(peer)
			if err != nil {
				xlog.Error("Failed to parse peer address", "error", err)
				continue
			}
			peers = append(peers, addr)
		}
		return peers
	}
)
