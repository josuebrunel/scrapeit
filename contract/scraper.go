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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"TaskFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"}],\"name\":\"complete\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"fail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"},{\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611b718061005b5f395ff3fe608060405260043610610054575f3560e01c8063198e2b8a1461005857806358709cf21461007457806378122f3a146100b65780638da5cb5b146100de578063eca4cbe614610108578063f3fe12c914610124575b5f5ffd5b610072600480360381019061006d9190611324565b61014c565b005b34801561007f575f5ffd5b5061009a6004803603810190610095919061139a565b610380565b6040516100ad979695949392919061150b565b60405180910390f35b3480156100c1575f5ffd5b506100dc60048036038101906100d7919061139a565b6105b3565b005b3480156100e9575f5ffd5b506100f2610ade565b6040516100ff919061158d565b60405180910390f35b610122600480360381019061011d9190611324565b610b02565b005b34801561012f575f5ffd5b5061014a6004803603810190610145919061139a565b61103e565b005b5f341161018e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610185906115f0565b60405180910390fd5b5f6040518060e001604052808481526020013373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160405180602001604052805f81525081526020015f60038111156101e9576101e8611480565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001348152509050806001825f01516040516102249190611648565b90815260200160405180910390205f820151815f019081610245919061185b565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190816102a1919061185b565b5060608201518160030190816102b7919061185b565b506080820151816004015f6101000a81548160ff021916908360038111156102e2576102e1611480565b5b021790555060a08201518160040160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c082015181600501559050507fe04b238473f76560b59321a44de66c30d1805378858d062d20174c0ce97d40a8815f0151826040015160405161037392919061192a565b60405180910390a1505050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546103b79061168b565b80601f01602080910402602001604051908101604052809291908181526020018280546103e39061168b565b801561042e5780601f106104055761010080835404028352916020019161042e565b820191905f5260205f20905b81548152906001019060200180831161041157829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020180546104689061168b565b80601f01602080910402602001604051908101604052809291908181526020018280546104949061168b565b80156104df5780601f106104b6576101008083540402835291602001916104df565b820191905f5260205f20905b8154815290600101906020018083116104c257829003601f168201915b5050505050908060030180546104f49061168b565b80601f01602080910402602001604051908101604052809291908181526020018280546105209061168b565b801561056b5780601f106105425761010080835404028352916020019161056b565b820191905f5260205f20905b81548152906001019060200180831161054e57829003601f168201915b505050505090806004015f9054906101000a900460ff16908060040160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b805f6001826040516105c59190611648565b90815260200160405180910390206040518060e00160405290815f820180546105ed9061168b565b80601f01602080910402602001604051908101604052809291908181526020018280546106199061168b565b80156106645780601f1061063b57610100808354040283529160200191610664565b820191905f5260205f20905b81548152906001019060200180831161064757829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546106d29061168b565b80601f01602080910402602001604051908101604052809291908181526020018280546106fe9061168b565b80156107495780601f1061072057610100808354040283529160200191610749565b820191905f5260205f20905b81548152906001019060200180831161072c57829003601f168201915b505050505081526020016003820180546107629061168b565b80601f016020809104026020016040519081016040528092919081815260200182805461078e9061168b565b80156107d95780601f106107b0576101008083540402835291602001916107d9565b820191905f5260205f20905b8154815290600101906020018083116107bc57829003601f168201915b50505050508152602001600482015f9054906101000a900460ff16600381111561080657610805611480565b5b600381111561081857610817611480565b5b81526020016004820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152505090503373ffffffffffffffffffffffffffffffffffffffff168160a0015173ffffffffffffffffffffffffffffffffffffffff16146108f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e7906119a9565b60405180910390fd5b6001600381111561090457610903611480565b5b8160800151600381111561091b5761091a611480565b5b14610924575f5ffd5b6001600381111561093857610937611480565b5b6001846040516109489190611648565b90815260200160405180910390206004015f9054906101000a900460ff16600381111561097857610977611480565b5b14610981575f5ffd5b3373ffffffffffffffffffffffffffffffffffffffff166001846040516109a89190611648565b908152602001604051809103902060040160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036109f9575f5ffd5b6003600184604051610a0b9190611648565b90815260200160405180910390206004015f6101000a81548160ff02191690836003811115610a3d57610a3c611480565b5b02179055505f600184604051610a539190611648565b908152602001604051809103902060040160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fde839f5aa8eca016f92c21edc678f15cf438ee8c44a2ba3c8775b94d96b004b583604051610ad191906119c7565b60405180910390a1505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b815f600182604051610b149190611648565b90815260200160405180910390206040518060e00160405290815f82018054610b3c9061168b565b80601f0160208091040260200160405190810160405280929190818152602001828054610b689061168b565b8015610bb35780601f10610b8a57610100808354040283529160200191610bb3565b820191905f5260205f20905b815481529060010190602001808311610b9657829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054610c219061168b565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4d9061168b565b8015610c985780601f10610c6f57610100808354040283529160200191610c98565b820191905f5260205f20905b815481529060010190602001808311610c7b57829003601f168201915b50505050508152602001600382018054610cb19061168b565b80601f0160208091040260200160405190810160405280929190818152602001828054610cdd9061168b565b8015610d285780601f10610cff57610100808354040283529160200191610d28565b820191905f5260205f20905b815481529060010190602001808311610d0b57829003601f168201915b50505050508152602001600482015f9054906101000a900460ff166003811115610d5557610d54611480565b5b6003811115610d6757610d66611480565b5b81526020016004820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152505090503373ffffffffffffffffffffffffffffffffffffffff168160a0015173ffffffffffffffffffffffffffffffffffffffff1614610e3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e36906119a9565b60405180910390fd5b60016003811115610e5357610e52611480565b5b81608001516003811115610e6a57610e69611480565b5b14610e73575f5ffd5b6002600185604051610e859190611648565b90815260200160405180910390206004015f6101000a81548160ff02191690836003811115610eb757610eb6611480565b5b021790555082600185604051610ecd9190611648565b90815260200160405180910390206003019081610eea919061185b565b507f1e6cd60eb3a95fa354fda400577a90e42e85f27a8031114340936da2b89c532f843385604051610f1e939291906119e7565b60405180910390a15f600185604051610f379190611648565b908152602001604051809103902060040160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600186604051610f8f9190611648565b908152602001604051809103902060050154604051610fad90611a57565b5f6040518083038185875af1925050503d805f8114610fe7576040519150601f19603f3d011682016040523d82523d5f602084013e610fec565b606091505b505090506001151581151514611037576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102e90611ab5565b60405180910390fd5b5050505050565b5f600381111561105157611050611480565b5b6001826040516110619190611648565b90815260200160405180910390206004015f9054906101000a900460ff16600381111561109157611090611480565b5b14806110ed57506003808111156110ab576110aa611480565b5b6001826040516110bb9190611648565b90815260200160405180910390206004015f9054906101000a900460ff1660038111156110eb576110ea611480565b5b145b61112c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112390611b1d565b60405180910390fd5b6001808260405161113d9190611648565b90815260200160405180910390206004015f6101000a81548160ff0219169083600381111561116f5761116e611480565b5b0217905550336001826040516111859190611648565b908152602001604051809103902060040160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611236826111f0565b810181811067ffffffffffffffff8211171561125557611254611200565b5b80604052505050565b5f6112676111d7565b9050611273828261122d565b919050565b5f67ffffffffffffffff82111561129257611291611200565b5b61129b826111f0565b9050602081019050919050565b828183375f83830152505050565b5f6112c86112c384611278565b61125e565b9050828152602081018484840111156112e4576112e36111ec565b5b6112ef8482856112a8565b509392505050565b5f82601f83011261130b5761130a6111e8565b5b813561131b8482602086016112b6565b91505092915050565b5f5f6040838503121561133a576113396111e0565b5b5f83013567ffffffffffffffff811115611357576113566111e4565b5b611363858286016112f7565b925050602083013567ffffffffffffffff811115611384576113836111e4565b5b611390858286016112f7565b9150509250929050565b5f602082840312156113af576113ae6111e0565b5b5f82013567ffffffffffffffff8111156113cc576113cb6111e4565b5b6113d8848285016112f7565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f611413826113e1565b61141d81856113eb565b935061142d8185602086016113fb565b611436816111f0565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61146a82611441565b9050919050565b61147a81611460565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600481106114be576114bd611480565b5b50565b5f8190506114ce826114ad565b919050565b5f6114dd826114c1565b9050919050565b6114ed816114d3565b82525050565b5f819050919050565b611505816114f3565b82525050565b5f60e0820190508181035f830152611523818a611409565b90506115326020830189611471565b81810360408301526115448188611409565b905081810360608301526115588187611409565b905061156760808301866114e4565b61157460a0830185611471565b61158160c08301846114fc565b98975050505050505050565b5f6020820190506115a05f830184611471565b92915050565b7f496e76616c696420616d6f756e740000000000000000000000000000000000005f82015250565b5f6115da600e836113eb565b91506115e5826115a6565b602082019050919050565b5f6020820190508181035f830152611607816115ce565b9050919050565b5f81905092915050565b5f611622826113e1565b61162c818561160e565b935061163c8185602086016113fb565b80840191505092915050565b5f6116538284611618565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806116a257607f821691505b6020821081036116b5576116b461165e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026117177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826116dc565b61172186836116dc565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61175c611757611752846114f3565b611739565b6114f3565b9050919050565b5f819050919050565b61177583611742565b61178961178182611763565b8484546116e8565b825550505050565b5f5f905090565b6117a0611791565b6117ab81848461176c565b505050565b5b818110156117ce576117c35f82611798565b6001810190506117b1565b5050565b601f821115611813576117e4816116bb565b6117ed846116cd565b810160208510156117fc578190505b611810611808856116cd565b8301826117b0565b50505b505050565b5f82821c905092915050565b5f6118335f1984600802611818565b1980831691505092915050565b5f61184b8383611824565b9150826002028217905092915050565b611864826113e1565b67ffffffffffffffff81111561187d5761187c611200565b5b611887825461168b565b6118928282856117d2565b5f60209050601f8311600181146118c3575f84156118b1578287015190505b6118bb8582611840565b865550611922565b601f1984166118d1866116bb565b5f5b828110156118f8578489015182556001820191506020850194506020810190506118d3565b868310156119155784890151611911601f891682611824565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f8301526119428185611409565b905081810360208301526119568184611409565b90509392505050565b7f4e6f74207375626d6974746572730000000000000000000000000000000000005f82015250565b5f611993600e836113eb565b915061199e8261195f565b602082019050919050565b5f6020820190508181035f8301526119c081611987565b9050919050565b5f6020820190508181035f8301526119df8184611409565b905092915050565b5f6060820190508181035f8301526119ff8186611409565b9050611a0e6020830185611471565b8181036040830152611a208184611409565b9050949350505050565b5f81905092915050565b50565b5f611a425f83611a2a565b9150611a4d82611a34565b5f82019050919050565b5f611a6182611a37565b9150819050919050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f611a9f600f836113eb565b9150611aaa82611a6b565b602082019050919050565b5f6020820190508181035f830152611acc81611a93565b9050919050565b7f4e6f742070656e64696e67206f72206661696c656400000000000000000000005f82015250565b5f611b076015836113eb565b9150611b1282611ad3565b602082019050919050565b5f6020820190508181035f830152611b3481611afb565b905091905056fea264697066735822122070301daad9b150d1db99aeaa8549073eb1e66c0f08ac0d29ec3a6fcfec54ff2b64736f6c634300081e0033",
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

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, address submitter, string input, string output, uint8 status, address assignee, uint256 amount)
func (_Contract *ContractCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	Uuid      string
	Submitter common.Address
	Input     string
	Output    string
	Status    uint8
	Assignee  common.Address
	Amount    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Uuid      string
		Submitter common.Address
		Input     string
		Output    string
		Status    uint8
		Assignee  common.Address
		Amount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Uuid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Submitter = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Input = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Output = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Assignee = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, address submitter, string input, string output, uint8 status, address assignee, uint256 amount)
func (_Contract *ContractSession) Tasks(arg0 string) (struct {
	Uuid      string
	Submitter common.Address
	Input     string
	Output    string
	Status    uint8
	Assignee  common.Address
	Amount    *big.Int
}, error) {
	return _Contract.Contract.Tasks(&_Contract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, address submitter, string input, string output, uint8 status, address assignee, uint256 amount)
func (_Contract *ContractCallerSession) Tasks(arg0 string) (struct {
	Uuid      string
	Submitter common.Address
	Input     string
	Output    string
	Status    uint8
	Assignee  common.Address
	Amount    *big.Int
}, error) {
	return _Contract.Contract.Tasks(&_Contract.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0xf3fe12c9.
//
// Solidity: function claim(string _uuid) returns()
func (_Contract *ContractTransactor) Claim(opts *bind.TransactOpts, _uuid string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claim", _uuid)
}

// Claim is a paid mutator transaction binding the contract method 0xf3fe12c9.
//
// Solidity: function claim(string _uuid) returns()
func (_Contract *ContractSession) Claim(_uuid string) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, _uuid)
}

// Claim is a paid mutator transaction binding the contract method 0xf3fe12c9.
//
// Solidity: function claim(string _uuid) returns()
func (_Contract *ContractTransactorSession) Claim(_uuid string) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, _uuid)
}

// Complete is a paid mutator transaction binding the contract method 0xeca4cbe6.
//
// Solidity: function complete(string _uuid, string _output) payable returns()
func (_Contract *ContractTransactor) Complete(opts *bind.TransactOpts, _uuid string, _output string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "complete", _uuid, _output)
}

// Complete is a paid mutator transaction binding the contract method 0xeca4cbe6.
//
// Solidity: function complete(string _uuid, string _output) payable returns()
func (_Contract *ContractSession) Complete(_uuid string, _output string) (*types.Transaction, error) {
	return _Contract.Contract.Complete(&_Contract.TransactOpts, _uuid, _output)
}

// Complete is a paid mutator transaction binding the contract method 0xeca4cbe6.
//
// Solidity: function complete(string _uuid, string _output) payable returns()
func (_Contract *ContractTransactorSession) Complete(_uuid string, _output string) (*types.Transaction, error) {
	return _Contract.Contract.Complete(&_Contract.TransactOpts, _uuid, _output)
}

// Create is a paid mutator transaction binding the contract method 0x198e2b8a.
//
// Solidity: function create(string _uuid, string _input) payable returns()
func (_Contract *ContractTransactor) Create(opts *bind.TransactOpts, _uuid string, _input string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "create", _uuid, _input)
}

// Create is a paid mutator transaction binding the contract method 0x198e2b8a.
//
// Solidity: function create(string _uuid, string _input) payable returns()
func (_Contract *ContractSession) Create(_uuid string, _input string) (*types.Transaction, error) {
	return _Contract.Contract.Create(&_Contract.TransactOpts, _uuid, _input)
}

// Create is a paid mutator transaction binding the contract method 0x198e2b8a.
//
// Solidity: function create(string _uuid, string _input) payable returns()
func (_Contract *ContractTransactorSession) Create(_uuid string, _input string) (*types.Transaction, error) {
	return _Contract.Contract.Create(&_Contract.TransactOpts, _uuid, _input)
}

// Fail is a paid mutator transaction binding the contract method 0x78122f3a.
//
// Solidity: function fail(string _uuid) returns()
func (_Contract *ContractTransactor) Fail(opts *bind.TransactOpts, _uuid string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fail", _uuid)
}

// Fail is a paid mutator transaction binding the contract method 0x78122f3a.
//
// Solidity: function fail(string _uuid) returns()
func (_Contract *ContractSession) Fail(_uuid string) (*types.Transaction, error) {
	return _Contract.Contract.Fail(&_Contract.TransactOpts, _uuid)
}

// Fail is a paid mutator transaction binding the contract method 0x78122f3a.
//
// Solidity: function fail(string _uuid) returns()
func (_Contract *ContractTransactorSession) Fail(_uuid string) (*types.Transaction, error) {
	return _Contract.Contract.Fail(&_Contract.TransactOpts, _uuid)
}

// ContractTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the Contract contract.
type ContractTaskCompletedIterator struct {
	Event *ContractTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskCompleted)
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
		it.Event = new(ContractTaskCompleted)
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
func (it *ContractTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskCompleted represents a TaskCompleted event raised by the Contract contract.
type ContractTaskCompleted struct {
	Uuid     string
	Assignee common.Address
	Output   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x1e6cd60eb3a95fa354fda400577a90e42e85f27a8031114340936da2b89c532f.
//
// Solidity: event TaskCompleted(string uuid, address assignee, string output)
func (_Contract *ContractFilterer) FilterTaskCompleted(opts *bind.FilterOpts) (*ContractTaskCompletedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskCompleted")
	if err != nil {
		return nil, err
	}
	return &ContractTaskCompletedIterator{contract: _Contract.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x1e6cd60eb3a95fa354fda400577a90e42e85f27a8031114340936da2b89c532f.
//
// Solidity: event TaskCompleted(string uuid, address assignee, string output)
func (_Contract *ContractFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractTaskCompleted) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskCompleted)
				if err := _Contract.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x1e6cd60eb3a95fa354fda400577a90e42e85f27a8031114340936da2b89c532f.
//
// Solidity: event TaskCompleted(string uuid, address assignee, string output)
func (_Contract *ContractFilterer) ParseTaskCompleted(log types.Log) (*ContractTaskCompleted, error) {
	event := new(ContractTaskCompleted)
	if err := _Contract.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Uuid  string
	Input string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0xe04b238473f76560b59321a44de66c30d1805378858d062d20174c0ce97d40a8.
//
// Solidity: event TaskCreated(string uuid, string input)
func (_Contract *ContractFilterer) FilterTaskCreated(opts *bind.FilterOpts) (*ContractTaskCreatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return &ContractTaskCreatedIterator{contract: _Contract.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0xe04b238473f76560b59321a44de66c30d1805378858d062d20174c0ce97d40a8.
//
// Solidity: event TaskCreated(string uuid, string input)
func (_Contract *ContractFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractTaskCreated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskCreated")
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

// ParseTaskCreated is a log parse operation binding the contract event 0xe04b238473f76560b59321a44de66c30d1805378858d062d20174c0ce97d40a8.
//
// Solidity: event TaskCreated(string uuid, string input)
func (_Contract *ContractFilterer) ParseTaskCreated(log types.Log) (*ContractTaskCreated, error) {
	event := new(ContractTaskCreated)
	if err := _Contract.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskFailedIterator is returned from FilterTaskFailed and is used to iterate over the raw logs and unpacked data for TaskFailed events raised by the Contract contract.
type ContractTaskFailedIterator struct {
	Event *ContractTaskFailed // Event containing the contract specifics and raw log

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
func (it *ContractTaskFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskFailed)
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
		it.Event = new(ContractTaskFailed)
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
func (it *ContractTaskFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskFailed represents a TaskFailed event raised by the Contract contract.
type ContractTaskFailed struct {
	Uuid string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTaskFailed is a free log retrieval operation binding the contract event 0xde839f5aa8eca016f92c21edc678f15cf438ee8c44a2ba3c8775b94d96b004b5.
//
// Solidity: event TaskFailed(string uuid)
func (_Contract *ContractFilterer) FilterTaskFailed(opts *bind.FilterOpts) (*ContractTaskFailedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskFailed")
	if err != nil {
		return nil, err
	}
	return &ContractTaskFailedIterator{contract: _Contract.contract, event: "TaskFailed", logs: logs, sub: sub}, nil
}

// WatchTaskFailed is a free log subscription operation binding the contract event 0xde839f5aa8eca016f92c21edc678f15cf438ee8c44a2ba3c8775b94d96b004b5.
//
// Solidity: event TaskFailed(string uuid)
func (_Contract *ContractFilterer) WatchTaskFailed(opts *bind.WatchOpts, sink chan<- *ContractTaskFailed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskFailed)
				if err := _Contract.contract.UnpackLog(event, "TaskFailed", log); err != nil {
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

// ParseTaskFailed is a log parse operation binding the contract event 0xde839f5aa8eca016f92c21edc678f15cf438ee8c44a2ba3c8775b94d96b004b5.
//
// Solidity: event TaskFailed(string uuid)
func (_Contract *ContractFilterer) ParseTaskFailed(log types.Log) (*ContractTaskFailed, error) {
	event := new(ContractTaskFailed)
	if err := _Contract.contract.UnpackLog(event, "TaskFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
