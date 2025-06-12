package etherutil

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/josuebrunel/gopkg/xlog"
)

func PrivateKeyFromHex(hexKey string) (*ecdsa.PrivateKey, error) {
	d, err := hexutil.Decode(hexKey)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(d)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func MustPrivateKeyFromHex(hexKey string) *ecdsa.PrivateKey {
	privateKey, err := PrivateKeyFromHex(hexKey)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func GetAddress(pk *ecdsa.PrivateKey) string {
	address := crypto.PubkeyToAddress(pk.PublicKey)
	xlog.Debug("Address", "address", address.Hex())
	return address.Hex()
}

func GetAddressFromHexKey(hexKey string) string {
	pk, err := PrivateKeyFromHex(hexKey)
	if err != nil {
		xlog.Error("Failed to get address from hex key", "error", err)
		return ""
	}
	return GetAddress(pk)
}

func GetAccountBalance(ctx context.Context, cli *ethclient.Client, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := cli.BalanceAt(ctx, account, nil)
	if err != nil {
		xlog.Error("Failed to get account balance", "error", err)
		return nil, err
	}
	return balance, nil
}

func AsEther(wei *big.Int) *big.Float {
	fBalance := new(big.Float)
	fBalance.SetString(wei.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	xlog.Debug("Account balance", "balance", ethValue.String(), "unit", "ether")
	return ethValue
}

func GetNonce(ctx context.Context, cli *ethclient.Client, pk *ecdsa.PrivateKey) (uint64, error) {
	address := crypto.PubkeyToAddress(pk.PublicKey)
	nonce, err := cli.PendingNonceAt(ctx, address)
	if err != nil {
		xlog.Error("Failed to get nonce", "error", err)
		return uint64(0), err
	}
	xlog.Debug("Nonce", "nonce", nonce)
	return nonce, nil
}

func MustGetNonce(ctx context.Context, cli *ethclient.Client, pk *ecdsa.PrivateKey) uint64 {
	nonce, err := GetNonce(ctx, cli, pk)
	if err != nil {
		panic(err)
	}
	return nonce
}

func GetGasPrice(ctx context.Context, cli *ethclient.Client) (*big.Int, error) {
	gasPrice, err := cli.SuggestGasPrice(ctx)
	if err != nil {
		xlog.Error("Failed to get gas price", "error", err)
		return nil, err
	}
	xlog.Debug("Gas price", "gasPrice", gasPrice)
	return gasPrice, nil
}

func MustGetGasPrice(ctx context.Context, cli *ethclient.Client) *big.Int {
	gasPrice, err := GetGasPrice(ctx, cli)
	if err != nil {
		panic(err)
	}
	return gasPrice
}

func GetChainID(ctx context.Context, cli *ethclient.Client) (*big.Int, error) {
	chainID, err := cli.NetworkID(ctx)
	if err != nil {
		xlog.Error("Failed to get chain ID", "error", err)
		return nil, err
	}
	xlog.Debug("Chain ID", "chainID", chainID)
	return chainID, nil
}

func GetAuth(ctx context.Context, cli *ethclient.Client, pk *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := GetChainID(ctx, cli)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		xlog.Error("Failed to create transaction", "error", err)
		return auth, nil
	}
	auth.GasPrice = MustGetGasPrice(ctx, cli)
	auth.Nonce = big.NewInt(int64(MustGetNonce(ctx, cli, pk)))
	auth.GasLimit = uint64(3000000)
	auth.Value = big.NewInt(0)
	return auth, nil
}

func MustGetAuth(ctx context.Context, cli *ethclient.Client, pk *ecdsa.PrivateKey) *bind.TransactOpts {
	auth, err := GetAuth(ctx, cli, pk)
	if err != nil {
		panic(err)
	}
	return auth
}

func BigIntFromString(s string) *big.Int {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil
	}
	return i
}

func WeiToEtherString(wei *big.Int, dec int) string {
	if wei == nil || wei.Sign() == 0 {
		return "0.0"
	}
	weiFloat := new(big.Float).SetInt(wei)
	ethFloat := new(big.Float).Quo(weiFloat, big.NewFloat(math.Pow10(dec)))
	return ethFloat.Text('f', dec)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
