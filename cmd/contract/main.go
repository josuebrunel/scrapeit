package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"os"
	"scrapeit/contract"
	"scrapeit/pkg/env"
	"scrapeit/pkg/etherutil"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/josuebrunel/gopkg/xlog"
)

func main() {
	var (
		privateKey string
		wallet     string
		err        error
		pk         *ecdsa.PrivateKey
		ctx        = context.Background()
	)
	flag.StringVar(&privateKey, "private-key", "", "Account private key")
	flag.StringVar(&wallet, "wallet", "", "Account wallet")
	flag.Parse()

	if privateKey == "" && wallet == "" {
		panic("private key or wallet is required")
	}

	if privateKey != "" && wallet != "" {
		panic("only one of private key or wallet can be provided")
	}

	if privateKey != "" {
		pk, err = etherutil.PrivateKeyFromHex(privateKey)
		if err != nil {
			xlog.Error("Invalid private key", "error", err)
			os.Exit(1)
		}
	}

	if err := deploy(ctx, pk); err != nil {
		xlog.Error("Failed to deploy", "error", err)
		os.Exit(1)
	}
}

func deploy(ctx context.Context, pk *ecdsa.PrivateKey) error {
	client, err := ethclient.Dial(env.Get(env.EtherAddrN))
	if err != nil {
		xlog.Error("Failed to connect to Ethereum node", "error", err)
		return err
	}
	defer client.Close()
	auth := etherutil.MustGetAuth(ctx, client, pk)
	ctAddr, ctTrx, ct, err := contract.DeployContract(auth, client)
	if err != nil {
		xlog.Error("Failed to deploy contract", "error", err)
		return err
	}
	xlog.Info("Contract deployed", "address", ctAddr.Hex(), "transaction", ctTrx.Hash().Hex(), "contract", ct)
	return nil
}
