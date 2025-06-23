// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"uuid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"uuid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"name\":\"TaskProcessed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_uuid\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"_assignee\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"name\":\"complete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_uuid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"assignee\",\"type\":\"address\"},{\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"inputHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"outputHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506113658061005b5f395ff3fe608060405260043610610054575f3560e01c806316c38b3c146100585780635c975abb146100805780638da5cb5b146100aa578063e579f500146100d4578063ec9c5f8b14610115578063f29a6d8814610131575b5f5ffd5b348015610063575f5ffd5b5061007e60048036038101906100799190610a02565b610159565b005b34801561008b575f5ffd5b50610094610203565b6040516100a19190610a3c565b60405180910390f35b3480156100b5575f5ffd5b506100be610215565b6040516100cb9190610a94565b60405180910390f35b3480156100df575f5ffd5b506100fa60048036038101906100f59190610ae0565b610239565b60405161010c96959493929190610bc5565b60405180910390f35b61012f600480360381019061012a9190610d60565b6102bc565b005b34801561013c575f5ffd5b5061015760048036038101906101529190610de4565b6105e6565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101de90610ebe565b60405180910390fd5b805f60146101000a81548160ff02191690831515021790555050565b5f60149054906101000a900460ff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900460ff16908060020154908060030154908060040154905086565b5f60149054906101000a900460ff161561030b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030290610f26565b60405180910390fd5b5f5f1b820361034f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034690610f8e565b60405180910390fd5b5f3411610391576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038890610ff6565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660015f8481526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610431576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042890611084565b60405180910390fd5b6040518060c001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f600281111561048757610486610b2b565b5b8152602001348152602001828051906020012081526020015f5f1b81525060015f8481526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff0219169083600281111561056b5761056a610b2b565b5b0217905550606082015181600201556080820151816003015560a082015181600401559050503373ffffffffffffffffffffffffffffffffffffffff16827fe6ed5f53ccab7a2f0d7f4018b0e6d22f351364b376253cd99e3e3d6016d260ce83346040516105da9291906110f2565b60405180910390a35050565b5f60149054906101000a900460ff1615610635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062c90610f26565b60405180910390fd5b5f60015f8681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036106d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d09061116a565b60405180910390fd5b5f60028111156106ec576106eb610b2b565b5b8160010160149054906101000a900460ff1660028111156107105761070f610b2b565b5b14610750576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610747906111d2565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036107be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b59061123a565b60405180910390fd5b83816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081156109025760018160010160146101000a81548160ff0219169083600281111561082e5761082d610b2b565b5b0217905550828051906020012081600401819055505f816002015411156108fd575f8473ffffffffffffffffffffffffffffffffffffffff16826002015460405161087890611285565b5f6040518083038185875af1925050503d805f81146108b2576040519150601f19603f3d011682016040523d82523d5f602084013e6108b7565b606091505b50509050806108fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f2906112e3565b60405180910390fd5b505b610930565b60028160010160146101000a81548160ff0219169083600281111561092a57610929610b2b565b5b02179055505b806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16857f597ae4c5fa5347b68d6cb82bc163309f77b5d3dde38e33574ebff0919c8ba2df8360010160149054906101000a900460ff16866040516109ad929190611301565b60405180910390a35050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f8115159050919050565b6109e1816109cd565b81146109eb575f5ffd5b50565b5f813590506109fc816109d8565b92915050565b5f60208284031215610a1757610a166109c5565b5b5f610a24848285016109ee565b91505092915050565b610a36816109cd565b82525050565b5f602082019050610a4f5f830184610a2d565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a7e82610a55565b9050919050565b610a8e81610a74565b82525050565b5f602082019050610aa75f830184610a85565b92915050565b5f819050919050565b610abf81610aad565b8114610ac9575f5ffd5b50565b5f81359050610ada81610ab6565b92915050565b5f60208284031215610af557610af46109c5565b5b5f610b0284828501610acc565b91505092915050565b5f610b1582610a55565b9050919050565b610b2581610b0b565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60038110610b6957610b68610b2b565b5b50565b5f819050610b7982610b58565b919050565b5f610b8882610b6c565b9050919050565b610b9881610b7e565b82525050565b5f819050919050565b610bb081610b9e565b82525050565b610bbf81610aad565b82525050565b5f60c082019050610bd85f830189610b1c565b610be56020830188610b1c565b610bf26040830187610b8f565b610bff6060830186610ba7565b610c0c6080830185610bb6565b610c1960a0830184610bb6565b979650505050505050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c7282610c2c565b810181811067ffffffffffffffff82111715610c9157610c90610c3c565b5b80604052505050565b5f610ca36109bc565b9050610caf8282610c69565b919050565b5f67ffffffffffffffff821115610cce57610ccd610c3c565b5b610cd782610c2c565b9050602081019050919050565b828183375f83830152505050565b5f610d04610cff84610cb4565b610c9a565b905082815260208101848484011115610d2057610d1f610c28565b5b610d2b848285610ce4565b509392505050565b5f82601f830112610d4757610d46610c24565b5b8135610d57848260208601610cf2565b91505092915050565b5f5f60408385031215610d7657610d756109c5565b5b5f610d8385828601610acc565b925050602083013567ffffffffffffffff811115610da457610da36109c9565b5b610db085828601610d33565b9150509250929050565b610dc381610b0b565b8114610dcd575f5ffd5b50565b5f81359050610dde81610dba565b92915050565b5f5f5f5f60808587031215610dfc57610dfb6109c5565b5b5f610e0987828801610acc565b9450506020610e1a87828801610dd0565b935050604085013567ffffffffffffffff811115610e3b57610e3a6109c9565b5b610e4787828801610d33565b9250506060610e58878288016109ee565b91505092959194509250565b5f82825260208201905092915050565b7f43616c6c6572206973206e6f74206f776e6572000000000000000000000000005f82015250565b5f610ea8601383610e64565b9150610eb382610e74565b602082019050919050565b5f6020820190508181035f830152610ed581610e9c565b9050919050565b7f436f6e74726163742069732070617573656400000000000000000000000000005f82015250565b5f610f10601283610e64565b9150610f1b82610edc565b602082019050919050565b5f6020820190508181035f830152610f3d81610f04565b9050919050565b7f555549442063616e6e6f7420626520656d7074790000000000000000000000005f82015250565b5f610f78601483610e64565b9150610f8382610f44565b602082019050919050565b5f6020820190508181035f830152610fa581610f6c565b9050919050565b7f496e76616c696420616d6f756e740000000000000000000000000000000000005f82015250565b5f610fe0600e83610e64565b9150610feb82610fac565b602082019050919050565b5f6020820190508181035f83015261100d81610fd4565b9050919050565b7f5461736b20776974682074686973205555494420616c726561647920657869735f8201527f7473000000000000000000000000000000000000000000000000000000000000602082015250565b5f61106e602283610e64565b915061107982611014565b604082019050919050565b5f6020820190508181035f83015261109b81611062565b9050919050565b5f81519050919050565b8281835e5f83830152505050565b5f6110c4826110a2565b6110ce8185610e64565b93506110de8185602086016110ac565b6110e781610c2c565b840191505092915050565b5f6040820190508181035f83015261110a81856110ba565b90506111196020830184610ba7565b9392505050565b7f496e76616c6964205555494400000000000000000000000000000000000000005f82015250565b5f611154600c83610e64565b915061115f82611120565b602082019050919050565b5f6020820190508181035f83015261118181611148565b9050919050565b7f5461736b206e6f742070656e64696e67000000000000000000000000000000005f82015250565b5f6111bc601083610e64565b91506111c782611188565b602082019050919050565b5f6020820190508181035f8301526111e9816111b0565b9050919050565b7f41737369676e656520616464726573732063616e6e6f74206265207a65726f005f82015250565b5f611224601f83610e64565b915061122f826111f0565b602082019050919050565b5f6020820190508181035f83015261125181611218565b9050919050565b5f81905092915050565b50565b5f6112705f83611258565b915061127b82611262565b5f82019050919050565b5f61128f82611265565b9150819050919050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f6112cd600f83610e64565b91506112d882611299565b602082019050919050565b5f6020820190508181035f8301526112fa816112c1565b9050919050565b5f6040820190506113145f830185610b8f565b818103602083015261132681846110ba565b9050939250505056fea2646970667358221220c92156876b3a82f3dfb04c71ef09d1332bf5ac3cb6344bf5f5ce281fc4ce028c64736f6c634300081e0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractCallerSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0xe579f500.
//
// Solidity: function tasks(bytes32 ) view returns(address submitter, address assignee, uint8 status, uint256 amount, bytes32 inputHash, bytes32 outputHash)
func (_Contract *ContractCaller) Tasks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Submitter  common.Address
	Assignee   common.Address
	Status     uint8
	Amount     *big.Int
	InputHash  [32]byte
	OutputHash [32]byte
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Submitter  common.Address
		Assignee   common.Address
		Status     uint8
		Amount     *big.Int
		InputHash  [32]byte
		OutputHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Submitter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Assignee = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.InputHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.OutputHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0xe579f500.
//
// Solidity: function tasks(bytes32 ) view returns(address submitter, address assignee, uint8 status, uint256 amount, bytes32 inputHash, bytes32 outputHash)
func (_Contract *ContractSession) Tasks(arg0 [32]byte) (struct {
	Submitter  common.Address
	Assignee   common.Address
	Status     uint8
	Amount     *big.Int
	InputHash  [32]byte
	OutputHash [32]byte
}, error) {
	return _Contract.Contract.Tasks(&_Contract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0xe579f500.
//
// Solidity: function tasks(bytes32 ) view returns(address submitter, address assignee, uint8 status, uint256 amount, bytes32 inputHash, bytes32 outputHash)
func (_Contract *ContractCallerSession) Tasks(arg0 [32]byte) (struct {
	Submitter  common.Address
	Assignee   common.Address
	Status     uint8
	Amount     *big.Int
	InputHash  [32]byte
	OutputHash [32]byte
}, error) {
	return _Contract.Contract.Tasks(&_Contract.CallOpts, arg0)
}

// Complete is a paid mutator transaction binding the contract method 0xf29a6d88.
//
// Solidity: function complete(bytes32 _uuid, address _assignee, string _output, bool _success) returns()
func (_Contract *ContractTransactor) Complete(opts *bind.TransactOpts, _uuid [32]byte, _assignee common.Address, _output string, _success bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "complete", _uuid, _assignee, _output, _success)
}

// Complete is a paid mutator transaction binding the contract method 0xf29a6d88.
//
// Solidity: function complete(bytes32 _uuid, address _assignee, string _output, bool _success) returns()
func (_Contract *ContractSession) Complete(_uuid [32]byte, _assignee common.Address, _output string, _success bool) (*types.Transaction, error) {
	return _Contract.Contract.Complete(&_Contract.TransactOpts, _uuid, _assignee, _output, _success)
}

// Complete is a paid mutator transaction binding the contract method 0xf29a6d88.
//
// Solidity: function complete(bytes32 _uuid, address _assignee, string _output, bool _success) returns()
func (_Contract *ContractTransactorSession) Complete(_uuid [32]byte, _assignee common.Address, _output string, _success bool) (*types.Transaction, error) {
	return _Contract.Contract.Complete(&_Contract.TransactOpts, _uuid, _assignee, _output, _success)
}

// Create is a paid mutator transaction binding the contract method 0xec9c5f8b.
//
// Solidity: function create(bytes32 _uuid, string _input) payable returns()
func (_Contract *ContractTransactor) Create(opts *bind.TransactOpts, _uuid [32]byte, _input string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "create", _uuid, _input)
}

// Create is a paid mutator transaction binding the contract method 0xec9c5f8b.
//
// Solidity: function create(bytes32 _uuid, string _input) payable returns()
func (_Contract *ContractSession) Create(_uuid [32]byte, _input string) (*types.Transaction, error) {
	return _Contract.Contract.Create(&_Contract.TransactOpts, _uuid, _input)
}

// Create is a paid mutator transaction binding the contract method 0xec9c5f8b.
//
// Solidity: function create(bytes32 _uuid, string _input) payable returns()
func (_Contract *ContractTransactorSession) Create(_uuid [32]byte, _input string) (*types.Transaction, error) {
	return _Contract.Contract.Create(&_Contract.TransactOpts, _uuid, _input)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Contract *ContractTransactor) SetPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPaused", _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Contract *ContractSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPaused(&_Contract.TransactOpts, _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Contract *ContractTransactorSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPaused(&_Contract.TransactOpts, _paused)
}

// ContractTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Contract contract.
type ContractTaskCreatedIterator struct {
	Event *ContractTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskCreated represents a TaskCreated event raised by the Contract contract.
type ContractTaskCreated struct {
	Uuid      [32]byte
	Submitter common.Address
	Input     string
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0xe6ed5f53ccab7a2f0d7f4018b0e6d22f351364b376253cd99e3e3d6016d260ce.
//
// Solidity: event TaskCreated(bytes32 indexed uuid, address indexed submitter, string input, uint256 amount)
func (_Contract *ContractFilterer) FilterTaskCreated(opts *bind.FilterOpts, uuid [][32]byte, submitter []common.Address) (*ContractTaskCreatedIterator, error) {

	var uuidRule []interface{}
	for _, uuidItem := range uuid {
		uuidRule = append(uuidRule, uuidItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskCreated", uuidRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskCreatedIterator{contract: _Contract.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0xe6ed5f53ccab7a2f0d7f4018b0e6d22f351364b376253cd99e3e3d6016d260ce.
//
// Solidity: event TaskCreated(bytes32 indexed uuid, address indexed submitter, string input, uint256 amount)
func (_Contract *ContractFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractTaskCreated, uuid [][32]byte, submitter []common.Address) (event.Subscription, error) {

	var uuidRule []interface{}
	for _, uuidItem := range uuid {
		uuidRule = append(uuidRule, uuidItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskCreated", uuidRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskCreated)
				if err := _Contract.contract.UnpackLog(event, "TaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCreated is a log parse operation binding the contract event 0xe6ed5f53ccab7a2f0d7f4018b0e6d22f351364b376253cd99e3e3d6016d260ce.
//
// Solidity: event TaskCreated(bytes32 indexed uuid, address indexed submitter, string input, uint256 amount)
func (_Contract *ContractFilterer) ParseTaskCreated(log types.Log) (*ContractTaskCreated, error) {
	event := new(ContractTaskCreated)
	if err := _Contract.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskProcessedIterator is returned from FilterTaskProcessed and is used to iterate over the raw logs and unpacked data for TaskProcessed events raised by the Contract contract.
type ContractTaskProcessedIterator struct {
	Event *ContractTaskProcessed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTaskProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskProcessed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTaskProcessed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTaskProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskProcessed represents a TaskProcessed event raised by the Contract contract.
type ContractTaskProcessed struct {
	Uuid     [32]byte
	Assignee common.Address
	Status   uint8
	Output   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskProcessed is a free log retrieval operation binding the contract event 0x597ae4c5fa5347b68d6cb82bc163309f77b5d3dde38e33574ebff0919c8ba2df.
//
// Solidity: event TaskProcessed(bytes32 indexed uuid, address indexed assignee, uint8 status, string output)
func (_Contract *ContractFilterer) FilterTaskProcessed(opts *bind.FilterOpts, uuid [][32]byte, assignee []common.Address) (*ContractTaskProcessedIterator, error) {

	var uuidRule []interface{}
	for _, uuidItem := range uuid {
		uuidRule = append(uuidRule, uuidItem)
	}
	var assigneeRule []interface{}
	for _, assigneeItem := range assignee {
		assigneeRule = append(assigneeRule, assigneeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskProcessed", uuidRule, assigneeRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskProcessedIterator{contract: _Contract.contract, event: "TaskProcessed", logs: logs, sub: sub}, nil
}

// WatchTaskProcessed is a free log subscription operation binding the contract event 0x597ae4c5fa5347b68d6cb82bc163309f77b5d3dde38e33574ebff0919c8ba2df.
//
// Solidity: event TaskProcessed(bytes32 indexed uuid, address indexed assignee, uint8 status, string output)
func (_Contract *ContractFilterer) WatchTaskProcessed(opts *bind.WatchOpts, sink chan<- *ContractTaskProcessed, uuid [][32]byte, assignee []common.Address) (event.Subscription, error) {

	var uuidRule []interface{}
	for _, uuidItem := range uuid {
		uuidRule = append(uuidRule, uuidItem)
	}
	var assigneeRule []interface{}
	for _, assigneeItem := range assignee {
		assigneeRule = append(assigneeRule, assigneeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskProcessed", uuidRule, assigneeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskProcessed)
				if err := _Contract.contract.UnpackLog(event, "TaskProcessed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskProcessed is a log parse operation binding the contract event 0x597ae4c5fa5347b68d6cb82bc163309f77b5d3dde38e33574ebff0919c8ba2df.
//
// Solidity: event TaskProcessed(bytes32 indexed uuid, address indexed assignee, uint8 status, string output)
func (_Contract *ContractFilterer) ParseTaskProcessed(log types.Log) (*ContractTaskProcessed, error) {
	event := new(ContractTaskProcessed)
	if err := _Contract.contract.UnpackLog(event, "TaskProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
