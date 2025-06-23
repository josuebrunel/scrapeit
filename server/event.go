package server

import (
	"context"
	"scrapeit/contract"
	"scrapeit/pkg/env"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/josuebrunel/gopkg/xlog"
)

func (s Server) ListenContractEvents(ctx context.Context) {
	ctAddr := common.HexToAddress(env.Get(env.EtherContractAddrN))
	xlog.Info("listening for contract events", "contract", ctAddr.Hex())
	wsUrl := env.Get(env.EtherWSAddrN)

	evCli, err := ethclient.Dial(wsUrl)
	if err != nil {
		xlog.Error("Error while connecting to Ethereum node", "error", err)
		panic(err)
	}
	defer evCli.Close()
	filterer, err := contract.NewContractFilterer(ctAddr, evCli)
	if err != nil {
		xlog.Error("Error while creating filterer", "error", err)
		panic(err)
	}
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	taskCreatedEvents := make(chan *contract.ContractTaskCreated)
	subCreated, err := filterer.WatchTaskCreated(watchOpts, taskCreatedEvents, nil, nil)
	if err != nil {
		xlog.Error("Error while watching task created events", "error", err)
		panic(err)
	}
	defer subCreated.Unsubscribe()

	taskProcessedEvents := make(chan *contract.ContractTaskProcessed)
	subProcessed, err := filterer.WatchTaskProcessed(watchOpts, taskProcessedEvents, nil, nil)
	if err != nil {
		xlog.Error("Error while watching task processed events", "error", err)
		panic(err)
	}
	defer subProcessed.Unsubscribe()

	xlog.Info("Sucessfully subscribed to contract events", "contract", ctAddr.Hex())

	for {
		select {
		case event := <-taskCreatedEvents:
			xlog.Info("Task created event", "event", event)
		case err := <-subCreated.Err():
			xlog.Error("TaskCreated subscription error", "error", err)
		case event := <-taskProcessedEvents:
			xlog.Info("Task processed event", "event", event)
		case err := <-subProcessed.Err():
			xlog.Error("TaskProcessed subscription error", "error", err)
		case <-ctx.Done():
			xlog.Info("Context done")
			return
		}
	}
}
