package env

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/josuebrunel/gopkg/xlog"
)

const (
	AppAddrN           = "APP_ADDR"
	AppAddrV           = ":8080"
	AppAuthBackendN    = "APP_AUTH_BACKEND"
	AppAuthBackendV    = "http://auth:9999"
	AppDatabaseN       = "APP_DATABASE"
	PubAddrN           = "PUB_ADDR"
	PubAddrV           = "/ip4/0.0.0.0/tcp/7654"
	PubTopicN          = "PUB_TOPIC"
	PubTopicV          = "scraping"
	PubPeersRoot       = "PUB_ROOT_PEERS"
	CDPRemoteN         = "CDP_REMOTE"
	IpfsAddrN          = "IPFS_ADDR"
	EtherAddrN         = "ETH_NETWORK"
	EtherWSAddrN       = "ETH_NETWORK_WS"
	EtherContractAddrN = "ETH_CONTRACT_ADDR"
	EtherPrivateKeyN   = "ETH_PRIVATE_KEY"
)

var DefaultValues = map[string]string{
	AppAddrN:        AppAddrV,
	AppAuthBackendN: AppAuthBackendV,
	PubAddrN:        PubAddrV,
	PubTopicN:       PubTopicV,
}

func Get(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return DefaultValues[key]
	}
	return v
}

type Settings struct {
	Peer     Peer     `json:"peer"`
	Ethereum Ethereum `json:"ethereum"`
}

type Peer struct {
	ID    string   `json:"id"`
	Addrs []string `json:"addrs"`
	Ipfs  string   `json:"ipfs"`
	Topic string   `json:"topic"`
}

type Ethereum struct {
	Network  string `json:"network"`
	Contract string `json:"contract"`
}

func NewSetting() Settings {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	url := Get(PubPeersRoot)
	xlog.Info("getting settings from root peer", "url", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		xlog.Error("failed to create http request", "error", err)
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		xlog.Error("failed to get peers root", "error", err)
		panic(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		xlog.Error("failed to read http response", "error", err)
		panic(err)
	}
	var s = struct {
		Data  Settings `json:"data"`
		Error string   `json:"error"`
	}{}
	if err = json.Unmarshal(data, &s); err != nil {
		xlog.Error("failed to unmarshal peers root", "error", err)
		panic(err)
	}
	xlog.Debug("settings from root peer", "settings", s)
	return s.Data
}
