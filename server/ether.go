package server

import (
	"context"
	"scrapeit/contract"
	"scrapeit/pkg/env"
	"scrapeit/task"

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
	subCreated, err := filterer.WatchTaskCreated(watchOpts, taskCreatedEvents, nil)
	if err != nil {
		xlog.Error("Error while watching task created events", "error", err)
		panic(err)
	}
	defer subCreated.Unsubscribe()

	taskClaimedEvents := make(chan *contract.ContractTaskClaimed)
	subClaimed, err := filterer.WatchTaskClaimed(watchOpts, taskClaimedEvents, nil)
	if err != nil {
		xlog.Error("Error while watching task claimed events", "error", err)
		panic(err)
	}
	defer subClaimed.Unsubscribe()

	taskProcessedEvents := make(chan *contract.ContractTaskProcessed)
	subProcessed, err := filterer.WatchTaskProcessed(watchOpts, taskProcessedEvents, nil)
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
			tsk := task.Task{
				ID:          event.Uuid,
				UserID:      event.UserId,
				Submitter:   event.Submitter.Hex(),
				Input:       event.Input,
				Status:      task.Pending,
				Transaction: event.Raw.TxHash.Hex(),
			}
			s.Publish(ctx, &tsk)
		case err := <-subCreated.Err():
			xlog.Error("TaskCreated subscription error", "error", err)
		case event := <-taskClaimedEvents:
			xlog.Info("Task claimed event", "event", event)
			tsk := task.Task{
				ID:       event.Uuid,
				Assignee: event.Assignee.Hex(),
				Status:   task.Active,
			}
			if err := s.Manager.Upsert(ctx, tsk); err != nil {
				xlog.Error("Error while upserting task", "error", err)
				continue
			}
		case err := <-subClaimed.Err():
			xlog.Error("TaskClaimed subscription error", "error", err)
		case event := <-taskProcessedEvents:
			xlog.Info("Task processed event", "event", event)
			tsk := task.Task{
				ID:     event.Uuid,
				Output: event.Output,
				Status: task.Status[event.Status],
			}
			if err := s.Manager.Upsert(ctx, tsk); err != nil {
				xlog.Error("Error while upserting task", "error", err)
				continue
			}
		case err := <-subProcessed.Err():
			xlog.Error("TaskProcessed subscription error", "error", err)
		case <-ctx.Done():
			xlog.Info("Context done")
			return
		}
	}
}
