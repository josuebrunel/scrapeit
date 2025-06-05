package common

import (
	"context"
	"errors"
	"io"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/josuebrunel/gopkg/xlog"
)

type IPFSClient struct {
	sh *shell.Shell
}

func NewIPFSClient(addr string) (*IPFSClient, error) {
	ipfsAddr := addr
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
	return &IPFSClient{sh: sh}, nil
}

func (c IPFSClient) UploadFile(ctx context.Context, content string) (string, error) {
	cid, err := c.sh.Add(strings.NewReader(content))
	if err != nil {
		xlog.Error("Error uploading file to IPFS", "error", err)
		return "", err
	}
	xlog.Info("File uploaded to IPFS", "cid", cid)
	return cid, nil
}

func (c IPFSClient) DownloadFile(ctx context.Context, cid string) (string, error) {
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
