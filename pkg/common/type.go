package common

import (
	"encoding/json"
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

type Message struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Data   string `json:"data"`
	Type   string `json:"type"`
}

func (m Message) JSONMarshal() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}

func (m *Message) JSONUnmarshal(b []byte) {
	if err := json.Unmarshal(b, m); err != nil {
		panic(err)
	}
}
