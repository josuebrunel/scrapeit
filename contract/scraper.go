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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"user_id\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"name\":\"TaskProcessed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_assignee\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"name\":\"complete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_user_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"user_id\",\"type\":\"string\"},{\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"assignee\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506119df8061005b5f395ff3fe608060405260043610610054575f3560e01c806316c38b3c1461005857806358709cf2146100805780635c975abb146100c55780635d28560a146100ef5780638da5cb5b1461010b578063da5ae48214610135575b5f5ffd5b348015610063575f5ffd5b5061007e60048036038101906100799190610cf2565b61015d565b005b34801561008b575f5ffd5b506100a660048036038101906100a19190610e59565b610207565b6040516100bc9a99989796959493929190610fca565b60405180910390f35b3480156100d0575f5ffd5b506100d96104d2565b6040516100e6919061108f565b60405180910390f35b610109600480360381019061010491906110a8565b6104e4565b005b348015610116575f5ffd5b5061011f6108bb565b60405161012c919061116c565b60405180910390f35b348015610140575f5ffd5b5061015b600480360381019061015691906111af565b6108df565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e290611295565b60405180910390fd5b805f60146101000a81548160ff02191690831515021790555050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461023e906112e0565b80601f016020809104026020016040519081016040528092919081815260200182805461026a906112e0565b80156102b55780601f1061028c576101008083540402835291602001916102b5565b820191905f5260205f20905b81548152906001019060200180831161029857829003601f168201915b5050505050908060010180546102ca906112e0565b80601f01602080910402602001604051908101604052809291908181526020018280546102f6906112e0565b80156103415780601f1061031857610100808354040283529160200191610341565b820191905f5260205f20905b81548152906001019060200180831161032457829003601f168201915b505050505090806002015f9054906101000a900460ff16908060020160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040180546103b3906112e0565b80601f01602080910402602001604051908101604052809291908181526020018280546103df906112e0565b801561042a5780601f106104015761010080835404028352916020019161042a565b820191905f5260205f20905b81548152906001019060200180831161040d57829003601f168201915b50505050509080600501805461043f906112e0565b80601f016020809104026020016040519081016040528092919081815260200182805461046b906112e0565b80156104b65780601f1061048d576101008083540402835291602001916104b6565b820191905f5260205f20905b81548152906001019060200180831161049957829003601f168201915b505050505090806006015490806007015490806008015490508a565b5f60149054906101000a900460ff1681565b5f60149054906101000a900460ff1615610533576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052a9061135a565b60405180910390fd5b5f835111610576576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056d906113c2565b60405180910390fd5b5f34116105b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105af9061142a565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff166001846040516105df9190611482565b90815260200160405180910390206003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610666576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065d90611508565b60405180910390fd5b5f6040518061014001604052808581526020018481526020015f600281111561069257610691610f00565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160405180602001604052805f8152508152602001348152602001428152602001428152509050806001825f015160405161070f9190611482565b90815260200160405180910390205f820151815f01908161073091906116c6565b50602082015181600101908161074691906116c6565b506040820151816002015f6101000a81548160ff0219169083600281111561077157610770610f00565b5b021790555060608201518160020160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600401908161081891906116c6565b5060c082015181600501908161082e91906116c6565b5060e0820151816006015561010082015181600701556101208201518160080155905050806080015173ffffffffffffffffffffffffffffffffffffffff167f62b751e0cbad6fdb8b8e78702ad231910ed198c3d2e6c6b2a2a93f1a02dc78f2825f015183602001518460a00151346040516108ad9493929190611795565b60405180910390a250505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60149054906101000a900460ff161561092e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109259061135a565b60405180910390fd5b5f845111610971576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610968906113c2565b60405180910390fd5b5f6001856040516109829190611482565b90815260200160405180910390209050600160028111156109a6576109a5610f00565b5b816002015f9054906101000a900460ff1660028111156109c9576109c8610f00565b5b03610a09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0090611837565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6e9061189f565b60405180910390fd5b838160020160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081610af0576002816002015f6101000a81548160ff02191690836002811115610ae657610ae5610f00565b5b0217905550610b2f565b6001816002015f6101000a81548160ff02191690836002811115610b1757610b16610f00565b5b021790555082816005019081610b2d91906116c6565b505b428160080181905550818015610b4857505f8160060154115b15610c1f575f8160020160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168260060154604051610b9a906118ea565b5f6040518083038185875af1925050503d805f8114610bd4576040519150601f19603f3d011682016040523d82523d5f602084013e610bd9565b606091505b5050905080610c1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1490611948565b60405180910390fd5b505b8060020160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff2d02e58fd2bc0370968819f376a61ab21c0ff70786f41b9247053e81dd2f97786836002015f9054906101000a900460ff1686604051610c9d93929190611966565b60405180910390a25050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f8115159050919050565b610cd181610cbd565b8114610cdb575f5ffd5b50565b5f81359050610cec81610cc8565b92915050565b5f60208284031215610d0757610d06610cb5565b5b5f610d1484828501610cde565b91505092915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d6b82610d25565b810181811067ffffffffffffffff82111715610d8a57610d89610d35565b5b80604052505050565b5f610d9c610cac565b9050610da88282610d62565b919050565b5f67ffffffffffffffff821115610dc757610dc6610d35565b5b610dd082610d25565b9050602081019050919050565b828183375f83830152505050565b5f610dfd610df884610dad565b610d93565b905082815260208101848484011115610e1957610e18610d21565b5b610e24848285610ddd565b509392505050565b5f82601f830112610e4057610e3f610d1d565b5b8135610e50848260208601610deb565b91505092915050565b5f60208284031215610e6e57610e6d610cb5565b5b5f82013567ffffffffffffffff811115610e8b57610e8a610cb9565b5b610e9784828501610e2c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610ed282610ea0565b610edc8185610eaa565b9350610eec818560208601610eba565b610ef581610d25565b840191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60038110610f3e57610f3d610f00565b5b50565b5f819050610f4e82610f2d565b919050565b5f610f5d82610f41565b9050919050565b610f6d81610f53565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f9c82610f73565b9050919050565b610fac81610f92565b82525050565b5f819050919050565b610fc481610fb2565b82525050565b5f610140820190508181035f830152610fe3818d610ec8565b90508181036020830152610ff7818c610ec8565b9050611006604083018b610f64565b611013606083018a610fa3565b6110206080830189610fa3565b81810360a08301526110328188610ec8565b905081810360c08301526110468187610ec8565b905061105560e0830186610fbb565b611063610100830185610fbb565b611071610120830184610fbb565b9b9a5050505050505050505050565b61108981610cbd565b82525050565b5f6020820190506110a25f830184611080565b92915050565b5f5f5f606084860312156110bf576110be610cb5565b5b5f84013567ffffffffffffffff8111156110dc576110db610cb9565b5b6110e886828701610e2c565b935050602084013567ffffffffffffffff81111561110957611108610cb9565b5b61111586828701610e2c565b925050604084013567ffffffffffffffff81111561113657611135610cb9565b5b61114286828701610e2c565b9150509250925092565b5f61115682610f73565b9050919050565b6111668161114c565b82525050565b5f60208201905061117f5f83018461115d565b92915050565b61118e81610f92565b8114611198575f5ffd5b50565b5f813590506111a981611185565b92915050565b5f5f5f5f608085870312156111c7576111c6610cb5565b5b5f85013567ffffffffffffffff8111156111e4576111e3610cb9565b5b6111f087828801610e2c565b94505060206112018782880161119b565b935050604085013567ffffffffffffffff81111561122257611221610cb9565b5b61122e87828801610e2c565b925050606061123f87828801610cde565b91505092959194509250565b7f43616c6c6572206973206e6f74206f776e6572000000000000000000000000005f82015250565b5f61127f601383610eaa565b915061128a8261124b565b602082019050919050565b5f6020820190508181035f8301526112ac81611273565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806112f757607f821691505b60208210810361130a576113096112b3565b5b50919050565b7f436f6e74726163742069732070617573656400000000000000000000000000005f82015250565b5f611344601283610eaa565b915061134f82611310565b602082019050919050565b5f6020820190508181035f83015261137181611338565b9050919050565b7f555549442063616e6e6f7420626520656d7074790000000000000000000000005f82015250565b5f6113ac601483610eaa565b91506113b782611378565b602082019050919050565b5f6020820190508181035f8301526113d9816113a0565b9050919050565b7f496e76616c696420616d6f756e740000000000000000000000000000000000005f82015250565b5f611414600e83610eaa565b915061141f826113e0565b602082019050919050565b5f6020820190508181035f83015261144181611408565b9050919050565b5f81905092915050565b5f61145c82610ea0565b6114668185611448565b9350611476818560208601610eba565b80840191505092915050565b5f61148d8284611452565b915081905092915050565b7f5461736b20776974682074686973205555494420616c726561647920657869735f8201527f7473000000000000000000000000000000000000000000000000000000000000602082015250565b5f6114f2602283610eaa565b91506114fd82611498565b604082019050919050565b5f6020820190508181035f83015261151f816114e6565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026115827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611547565b61158c8683611547565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6115c76115c26115bd84610fb2565b6115a4565b610fb2565b9050919050565b5f819050919050565b6115e0836115ad565b6115f46115ec826115ce565b848454611553565b825550505050565b5f5f905090565b61160b6115fc565b6116168184846115d7565b505050565b5b818110156116395761162e5f82611603565b60018101905061161c565b5050565b601f82111561167e5761164f81611526565b61165884611538565b81016020851015611667578190505b61167b61167385611538565b83018261161b565b50505b505050565b5f82821c905092915050565b5f61169e5f1984600802611683565b1980831691505092915050565b5f6116b6838361168f565b9150826002028217905092915050565b6116cf82610ea0565b67ffffffffffffffff8111156116e8576116e7610d35565b5b6116f282546112e0565b6116fd82828561163d565b5f60209050601f83116001811461172e575f841561171c578287015190505b61172685826116ab565b86555061178d565b601f19841661173c86611526565b5f5b828110156117635784890151825560018201915060208501945060208101905061173e565b86831015611780578489015161177c601f89168261168f565b8355505b6001600288020188555050505b505050505050565b5f6080820190508181035f8301526117ad8187610ec8565b905081810360208301526117c18186610ec8565b905081810360408301526117d58185610ec8565b90506117e46060830184610fbb565b95945050505050565b7f5461736b20616c726561647920636f6d706c65746564000000000000000000005f82015250565b5f611821601683610eaa565b915061182c826117ed565b602082019050919050565b5f6020820190508181035f83015261184e81611815565b9050919050565b7f41737369676e656520616464726573732063616e6e6f74206265207a65726f005f82015250565b5f611889601f83610eaa565b915061189482611855565b602082019050919050565b5f6020820190508181035f8301526118b68161187d565b9050919050565b5f81905092915050565b50565b5f6118d55f836118bd565b91506118e0826118c7565b5f82019050919050565b5f6118f4826118ca565b9150819050919050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f611932600f83610eaa565b915061193d826118fe565b602082019050919050565b5f6020820190508181035f83015261195f81611926565b9050919050565b5f6060820190508181035f83015261197e8186610ec8565b905061198d6020830185610f64565b818103604083015261199f8184610ec8565b905094935050505056fea2646970667358221220d564b97bb0bded76dd73b0ba7437f99f6c278fb9c0d6e9e88a0f1ce7eb88724b64736f6c634300081e0033",
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

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, string user_id, uint8 status, address assignee, address submitter, string input, string output, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	Uuid      string
	UserId    string
	Status    uint8
	Assignee  common.Address
	Submitter common.Address
	Input     string
	Output    string
	Amount    *big.Int
	CreatedAt *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Uuid      string
		UserId    string
		Status    uint8
		Assignee  common.Address
		Submitter common.Address
		Input     string
		Output    string
		Amount    *big.Int
		CreatedAt *big.Int
		UpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Uuid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.UserId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Assignee = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Submitter = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Input = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Output = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, string user_id, uint8 status, address assignee, address submitter, string input, string output, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractSession) Tasks(arg0 string) (struct {
	Uuid      string
	UserId    string
	Status    uint8
	Assignee  common.Address
	Submitter common.Address
	Input     string
	Output    string
	Amount    *big.Int
	CreatedAt *big.Int
	UpdatedAt *big.Int
}, error) {
	return _Contract.Contract.Tasks(&_Contract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, string user_id, uint8 status, address assignee, address submitter, string input, string output, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractCallerSession) Tasks(arg0 string) (struct {
	Uuid      string
	UserId    string
	Status    uint8
	Assignee  common.Address
	Submitter common.Address
	Input     string
	Output    string
	Amount    *big.Int
	CreatedAt *big.Int
	UpdatedAt *big.Int
}, error) {
	return _Contract.Contract.Tasks(&_Contract.CallOpts, arg0)
}

// Complete is a paid mutator transaction binding the contract method 0xda5ae482.
//
// Solidity: function complete(string _uuid, address _assignee, string _output, bool _success) returns()
func (_Contract *ContractTransactor) Complete(opts *bind.TransactOpts, _uuid string, _assignee common.Address, _output string, _success bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "complete", _uuid, _assignee, _output, _success)
}

// Complete is a paid mutator transaction binding the contract method 0xda5ae482.
//
// Solidity: function complete(string _uuid, address _assignee, string _output, bool _success) returns()
func (_Contract *ContractSession) Complete(_uuid string, _assignee common.Address, _output string, _success bool) (*types.Transaction, error) {
	return _Contract.Contract.Complete(&_Contract.TransactOpts, _uuid, _assignee, _output, _success)
}

// Complete is a paid mutator transaction binding the contract method 0xda5ae482.
//
// Solidity: function complete(string _uuid, address _assignee, string _output, bool _success) returns()
func (_Contract *ContractTransactorSession) Complete(_uuid string, _assignee common.Address, _output string, _success bool) (*types.Transaction, error) {
	return _Contract.Contract.Complete(&_Contract.TransactOpts, _uuid, _assignee, _output, _success)
}

// Create is a paid mutator transaction binding the contract method 0x5d28560a.
//
// Solidity: function create(string _uuid, string _user_id, string _input) payable returns()
func (_Contract *ContractTransactor) Create(opts *bind.TransactOpts, _uuid string, _user_id string, _input string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "create", _uuid, _user_id, _input)
}

// Create is a paid mutator transaction binding the contract method 0x5d28560a.
//
// Solidity: function create(string _uuid, string _user_id, string _input) payable returns()
func (_Contract *ContractSession) Create(_uuid string, _user_id string, _input string) (*types.Transaction, error) {
	return _Contract.Contract.Create(&_Contract.TransactOpts, _uuid, _user_id, _input)
}

// Create is a paid mutator transaction binding the contract method 0x5d28560a.
//
// Solidity: function create(string _uuid, string _user_id, string _input) payable returns()
func (_Contract *ContractTransactorSession) Create(_uuid string, _user_id string, _input string) (*types.Transaction, error) {
	return _Contract.Contract.Create(&_Contract.TransactOpts, _uuid, _user_id, _input)
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
	Uuid      string
	UserId    string
	Submitter common.Address
	Input     string
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x62b751e0cbad6fdb8b8e78702ad231910ed198c3d2e6c6b2a2a93f1a02dc78f2.
//
// Solidity: event TaskCreated(string uuid, string user_id, address indexed submitter, string input, uint256 amount)
func (_Contract *ContractFilterer) FilterTaskCreated(opts *bind.FilterOpts, submitter []common.Address) (*ContractTaskCreatedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskCreated", submitterRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskCreatedIterator{contract: _Contract.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x62b751e0cbad6fdb8b8e78702ad231910ed198c3d2e6c6b2a2a93f1a02dc78f2.
//
// Solidity: event TaskCreated(string uuid, string user_id, address indexed submitter, string input, uint256 amount)
func (_Contract *ContractFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractTaskCreated, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskCreated", submitterRule)
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

// ParseTaskCreated is a log parse operation binding the contract event 0x62b751e0cbad6fdb8b8e78702ad231910ed198c3d2e6c6b2a2a93f1a02dc78f2.
//
// Solidity: event TaskCreated(string uuid, string user_id, address indexed submitter, string input, uint256 amount)
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
	Uuid     string
	Assignee common.Address
	Status   uint8
	Output   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskProcessed is a free log retrieval operation binding the contract event 0xf2d02e58fd2bc0370968819f376a61ab21c0ff70786f41b9247053e81dd2f977.
//
// Solidity: event TaskProcessed(string uuid, address indexed assignee, uint8 status, string output)
func (_Contract *ContractFilterer) FilterTaskProcessed(opts *bind.FilterOpts, assignee []common.Address) (*ContractTaskProcessedIterator, error) {

	var assigneeRule []interface{}
	for _, assigneeItem := range assignee {
		assigneeRule = append(assigneeRule, assigneeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskProcessed", assigneeRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskProcessedIterator{contract: _Contract.contract, event: "TaskProcessed", logs: logs, sub: sub}, nil
}

// WatchTaskProcessed is a free log subscription operation binding the contract event 0xf2d02e58fd2bc0370968819f376a61ab21c0ff70786f41b9247053e81dd2f977.
//
// Solidity: event TaskProcessed(string uuid, address indexed assignee, uint8 status, string output)
func (_Contract *ContractFilterer) WatchTaskProcessed(opts *bind.WatchOpts, sink chan<- *ContractTaskProcessed, assignee []common.Address) (event.Subscription, error) {

	var assigneeRule []interface{}
	for _, assigneeItem := range assignee {
		assigneeRule = append(assigneeRule, assigneeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskProcessed", assigneeRule)
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

// ParseTaskProcessed is a log parse operation binding the contract event 0xf2d02e58fd2bc0370968819f376a61ab21c0ff70786f41b9247053e81dd2f977.
//
// Solidity: event TaskProcessed(string uuid, address indexed assignee, uint8 status, string output)
func (_Contract *ContractFilterer) ParseTaskProcessed(log types.Log) (*ContractTaskProcessed, error) {
	event := new(ContractTaskProcessed)
	if err := _Contract.contract.UnpackLog(event, "TaskProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
