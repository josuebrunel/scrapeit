package task

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"scrapeit/contract"
	"scrapeit/pkg/etherutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/stephenafamo/bob"
)

var Status = map[uint8]TaskStatus{
	0: Pending,
	1: Completed,
	2: Failed,
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

func (e EtherManager) GetClient() ethclient.Client {
	return *e.client
}

func (e EtherManager) GetDB() bob.DB {
	return e.db
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

func (e EtherManager) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, e.client, tx)
	if err != nil {
		xlog.Error("Failed to mine transaction", "error", err)
		return receipt, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		xlog.Error("Transaction failed",
			"error", err,
			"transaction", tx.Hash().Hex(),
			"status", receipt.Status,
			"gasUser", receipt.GasUsed,
		)
		return receipt, err
	}
	return receipt, nil
}

func (e EtherManager) Create(ctx context.Context, pk *ecdsa.PrivateKey, tsk Task) (*types.Transaction, error) {
	auth := e.getAuth(ctx, pk)
	if auth == nil {
		return nil, errors.New("failed to create transaction")
	}
	auth.Value = etherutil.BigIntFromString(tsk.Reward)
	instance, err := contract.NewContract(e.contractAddr, e.client)
	if err != nil {
		xlog.Error("Failed to instantiate a Token contract", "error", err)
		return nil, err
	}
	tx, err := instance.Create(auth, s2b32(tsk.ID), tsk.Input)
	if err != nil {
		xlog.Error("Failed to create task", "error", err)
		return nil, err
	}
	xlog.Info("Task created", "tx", tx.Hash().Hex())
	return tx, nil
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
	}, s2b32(id))
	if err != nil {
		xlog.Error("Failed to get task", "error", err)
		return tsk, err
	}
	tsk = Task{
		Assignee:  tx.Assignee.String(),
		Submitter: tx.Submitter.String(),
		Input:     b322s(tx.InputHash),
		Output:    b322s(tx.OutputHash),
		Status:    TaskStatus(Status[tx.Status]),
	}
	return tsk, nil
}

func (e EtherManager) Complete(ctx context.Context, id string, assignee string, output string, success bool, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
	auth := e.getAuth(ctx, pk)
	if auth == nil {
		return nil, errors.New("failed to create transaction")
	}
	instance, err := contract.NewContract(e.contractAddr, e.client)
	if err != nil {
		xlog.Error("Failed to instantiate a Token contract", "error", err)
		return nil, err
	}
	xlog.Info("Completing task", "id", id, "assignee", assignee, "output", output)
	tx, err := instance.Complete(auth, s2b32(id), common.HexToAddress(assignee), output, success)
	if err != nil {
		xlog.Error("Failed to complete task", "error", err)
		return nil, err
	}
	xlog.Info("Task completed", "tx", tx.Hash().Hex(), "id", id, "output", output, "assignee", pk.PublicKey)
	return tx, nil
}

func str2byte32(s string) [32]byte {
	var b [32]byte
	copy(b[:], s)
	return b
}

func byte322str(b [32]byte) string {
	return string(b[:])
}

var s2b32 = str2byte32
var b322s = byte322str
