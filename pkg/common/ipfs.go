package common

import (
	"context"
	"errors"
	"io"
	"scrapeit/pkg/env"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/josuebrunel/gopkg/xlog"
)

var (
	IPFSClient *ipfsClient
)

func init() {
	var err error
	IPFSClient, err = NewIPFSClient()
	if err != nil {
		xlog.Error("Failed to create IPFS client", "error", err)
	} else {
		xlog.Info("IPFS client initialized successfully")
	}
}

type ipfsClient struct {
	sh *shell.Shell
}

func NewIPFSClient() (*ipfsClient, error) {
	ipfsAddr := env.Get(env.IpfsAddrN)
	if ipfsAddr == "" {
		xlog.Error("Invalid IPFS address", "address", ipfsAddr)
		return nil, errors.New("Invalid IPFS address")
	}
	sh := shell.NewShell(ipfsAddr)
	if !sh.IsUp() {
		xlog.Warn("IPFS shell is not up", "address", ipfsAddr)
	} else {
		xlog.Info("IPFS shell is up", "address", ipfsAddr)
	}
	return &ipfsClient{sh: sh}, nil
}

func (c ipfsClient) UploadFile(ctx context.Context, content string) (string, error) {
	cid, err := c.sh.Add(strings.NewReader(content))
	if err != nil {
		xlog.Error("Error uploading file to IPFS", "error", err)
		return "", err
	}
	xlog.Info("File uploaded to IPFS", "cid", cid)
	return cid, nil
}

func (c ipfsClient) DownloadFile(ctx context.Context, cid string) (string, error) {
	reader, err := c.sh.Cat(cid)
	if err != nil {
		xlog.Error("Error downloading file from IPFS", "error", err)
		return "", err
	}
	defer reader.Close()
	content, err := io.ReadAll(reader)
	if err != nil {
		xlog.Error("Error reading file from IPFS", "error", err)
		return "", err
	}
	xlog.Info("File downloaded from IPFS", "cid", cid)
	return string(content), nil
}
