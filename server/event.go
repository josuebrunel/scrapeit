package server

import (
	"context"
	"scrapeit/contract"
	"scrapeit/pkg/env"
	"scrapeit/pkg/models"
	"scrapeit/task"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofrs/uuid/v5"
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
		case err := <-subCreated.Err():
			xlog.Error("TaskCreated subscription error", "error", err)
		case event := <-taskProcessedEvents:
			xlog.Info("Task processed event", "event", event)
			tsk, err := models.FindTask(ctx, s.Manager.GetDB(), uuid.Must(uuid.FromString(event.Uuid)))
			if err != nil {
				xlog.Error("Error while finding task", "error", err, "uuid", event.Uuid)
				continue
			}
			if err := s.Manager.Upsert(ctx, task.Task{
				ID:          tsk.ID.String(),
				Assignee:    tsk.Assignee.V,
				Output:      event.Output,
				Status:      task.TaskStatus(task.Status[event.Status]),
				Paid:        task.Status[event.Status] == task.Completed,
				Transaction: event.Raw.TxHash.Hex(),
				UpdatedAt:   tsk.UpdatedAt.V,
			}); err != nil {
				xlog.Error("Error while updating task", "error", err, "uuid", event.Uuid)
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
