package common

import (
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/multiformats/go-multiaddr"
)

const (
	ActionScrape     = "scrape"
	ActionScreenshot = "screenshot"
)

func MultiAddrFromStrings(ss []string) []multiaddr.Multiaddr {
	var r = []multiaddr.Multiaddr{}
	for _, a := range ss {
		addr, err := multiaddr.NewMultiaddr(a)
		if err != nil {
			xlog.Error("failed to parse multiaddr", "error", err)
			continue
		}
		r = append(r, addr)
	}
	return r
}
