package env

import "os"

const (
	AppAddrN     = "APP_ADDR"
	AppAddrV     = ":8080"
	PubAddrN     = "PUB_ADDR"
	PubAddrV     = "/ip4/0.0.0.0/tcp/7654"
	PubTopicN    = "PUB_TOPIC"
	PubTopicV    = "scraping"
	PubPeersRoot = "PUB_ROOT_PEERS"
)

var DefaultValues = map[string]string{
	AppAddrN:  AppAddrV,
	PubAddrN:  PubAddrV,
	PubTopicN: PubTopicV,
}

func Get(key string) string {
	v, ok := DefaultValues[key]
	if !ok {
		return os.Getenv(key)
	}
	return v
}
