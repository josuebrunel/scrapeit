package task

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"scrapeit/contract"
	"scrapeit/pkg/etherutil"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/stephenafamo/bob"
)

var Status = map[uint8]TaskStatus{
	0: Pending,
	1: Active,
	2: Completed,
	3: Failed,
}

type EtherManager struct {
	client       *ethclient.Client
	chainID      *big.Int
	contractAddr common.Address
	db           bob.DB
}

func NewEtherManager(network, contractAddr string) *EtherManager {
	cli, err := ethclient.Dial(network)
	if err != nil {
		xlog.Error("Error while connecting to Ethereum node", "error", err)
		panic(err)
	}
	chainID, err := cli.NetworkID(context.Background())
	if err != nil {
		xlog.Error("Error while getting chain ID", "error", err)
		panic(err)
	}
	return &EtherManager{
		client:       cli,
		chainID:      chainID,
		contractAddr: common.HexToAddress(contractAddr),
	}
}

func (e *EtherManager) SetDB(db bob.DB) {
	e.db = db
}

func (e *EtherManager) Close() {
	e.client.Close()
}

func (e EtherManager) getAuth(ctx context.Context, pk *ecdsa.PrivateKey) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(pk, e.chainID)
	if err != nil {
		xlog.Error("Failed to create transaction", "error", err)
		return nil
	}
	auth.GasPrice = etherutil.MustGetGasPrice(ctx, e.client)
	auth.Nonce = big.NewInt(int64(etherutil.MustGetNonce(ctx, e.client, pk)))
	auth.GasLimit = uint64(300000)
	return auth
}

func (e EtherManager) Create(ctx context.Context, pk *ecdsa.PrivateKey, tsk Task) error {
	auth := e.getAuth(ctx, pk)
	if auth == nil {
		return errors.New("failed to create transaction")
	}
	auth.Value = big.NewInt(150000)
	instance, err := contract.NewContract(e.contractAddr, e.client)
	if err != nil {
		xlog.Error("Failed to instantiate a Token contract", "error", err)
		return err
	}
	tx, err := instance.Create(auth, tsk.ID, tsk.UserID, tsk.Input)
	if err != nil {
		xlog.Error("Failed to create task", "error", err)
		return err
	}
	xlog.Info("Task created", "tx", tx.Hash().Hex())
	return nil
}

func (e EtherManager) Get(ctx context.Context, id string, pk *ecdsa.PrivateKey) (Task, error) {
	var tsk Task
	auth := etherutil.MustGetAuth(ctx, e.client, pk)
	instance, err := contract.NewContract(e.contractAddr, e.client)
	if err != nil {
		xlog.Error("Failed to instantiate a Token contract", "error", err)
		return tsk, err
	}
	tx, err := instance.Tasks(&bind.CallOpts{
		From:    auth.From,
		Context: ctx,
	}, id)
	if err != nil {
		xlog.Error("Failed to get task", "error", err)
		return tsk, err
	}
	tsk = Task{
		ID:        tx.Uuid,
		Assignee:  tx.Assignee.String(),
		Submitter: tx.Submitter.String(),
		Input:     tx.Input,
		Output:    tx.Output,
		Status:    TaskStatus(Status[tx.Status]),
		CreatedAt: time.Unix(tx.CreatedAt.Int64(), 0).UTC(),
		UpdatedAt: time.Unix(tx.UpdatedAt.Int64(), 0).UTC(),
	}
	return tsk, nil
}

func (e EtherManager) Claim(ctx context.Context, id string, pk *ecdsa.PrivateKey) error {
	auth := e.getAuth(ctx, pk)
	if auth == nil {
		return errors.New("failed to create transaction")
	}
	instance, err := contract.NewContract(e.contractAddr, e.client)
	if err != nil {
		xlog.Error("Failed to instantiate a Token contract", "error", err)
		return err
	}
	tx, err := instance.Claim(auth, id)
	if err != nil {
		xlog.Error("Failed to claim task", "error", err)
		return err
	}

	xlog.Info("Task claimed", "tx", tx.Hash().Hex(), "id", id, "assignee", pk.PublicKey)
	return nil
}

func (e EtherManager) Complete(ctx context.Context, id string, output string, pk *ecdsa.PrivateKey) error {
	auth := e.getAuth(ctx, pk)
	if auth == nil {
		return errors.New("failed to create transaction")
	}
	instance, err := contract.NewContract(e.contractAddr, e.client)
	if err != nil {
		xlog.Error("Failed to instantiate a Token contract", "error", err)
		return err
	}
	tx, err := instance.Complete(auth, id, output)
	if err != nil {
		xlog.Error("Failed to complete task", "error", err)
		return err
	}
	xlog.Info("Task completed", "tx", tx.Hash().Hex(), "id", id, "output", output, "assignee", pk.PublicKey)
	return nil
}

func (e EtherManager) Fail(ctx context.Context, id string, pk *ecdsa.PrivateKey) error {
	auth := e.getAuth(ctx, pk)
	if auth == nil {
		return errors.New("failed to create transaction")
	}
	instance, err := contract.NewContract(e.contractAddr, e.client)
	if err != nil {
		xlog.Error("Failed to instantiate a Token contract", "error", err)
		return err
	}
	tx, err := instance.Fail(auth, id)
	if err != nil {
		xlog.Error("Failed to fail task", "error", err)
		return err
	}
	xlog.Info("Task failed", "tx", tx.Hash().Hex(), "id", id)
	return nil
}
