package env

import "os"

const (
	AppAddrN           = "APP_ADDR"
	AppAddrV           = ":8080"
	PubAddrN           = "PUB_ADDR"
	PubAddrV           = "/ip4/0.0.0.0/tcp/7654"
	PubTopicN          = "PUB_TOPIC"
	PubTopicV          = "scraping"
	PubPeersRoot       = "PUB_ROOT_PEERS"
	RedisAddrN         = "REDIS_ADDR"
	RedisAddrV         = "localhost:6379"
	CDPRemoteN         = "CDP_REMOTE"
	IpfsAddrN          = "IPFS_ADDR"
	EtherAddrN         = "ETH_NETWORK"
	EtherContractAddrN = "ETH_CONTRACT_ADDR"
	EtherPrivateKeyN   = "ETH_PRIVATE_KEY"
)

var DefaultValues = map[string]string{
	AppAddrN:   AppAddrV,
	PubAddrN:   PubAddrV,
	PubTopicN:  PubTopicV,
	RedisAddrN: RedisAddrV,
}

func Get(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return DefaultValues[key]
	}
	return v
}
