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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"}],\"name\":\"TaskClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assigneeWhoFailed\",\"type\":\"address\"}],\"name\":\"TaskFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"}],\"name\":\"complete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"fail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"},{\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061207d8061005b5f395ff3fe608060405260043610610054575f3560e01c8063198e2b8a1461005857806358709cf21461007457806378122f3a146100b85780638da5cb5b146100e0578063eca4cbe61461010a578063f3fe12c914610132575b5f5ffd5b610072600480360381019061006d9190611626565b61015a565b005b34801561007f575f5ffd5b5061009a6004803603810190610095919061169c565b6104ba565b6040516100af9998979695949392919061180d565b60405180910390f35b3480156100c3575f5ffd5b506100de60048036038101906100d9919061169c565b6106f9565b005b3480156100eb575f5ffd5b506100f4610c4c565b60405161010191906118ad565b60405180910390f35b348015610115575f5ffd5b50610130600480360381019061012b9190611626565b610c70565b005b34801561013d575f5ffd5b506101586004803603810190610153919061169c565b611266565b005b5f82511161019d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019490611910565b60405180910390fd5b5f34116101df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d690611978565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660018360405161020691906119d0565b90815260200160405180910390206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461028d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028490611a56565b60405180910390fd5b5f6040518061012001604052808481526020013373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160405180602001604052805f81525081526020015f60038111156102e9576102e8611782565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001348152602001428152602001428152509050806001825f015160405161033091906119d0565b90815260200160405180910390205f820151815f0190816103519190611c71565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190816103ad9190611c71565b5060608201518160030190816103c39190611c71565b506080820151816004015f6101000a81548160ff021916908360038111156103ee576103ed611782565b5b021790555060a08201518160040160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816005015560e0820151816006015561010082015181600701559050503373ffffffffffffffffffffffffffffffffffffffff167f9bc0fc01301fa260e9f61a4f2be24187c8a9f7cebcaf3a83251af30a355f62d2825f01518360400151346040516104ad93929190611d40565b60405180910390a2505050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546104f190611aa1565b80601f016020809104026020016040519081016040528092919081815260200182805461051d90611aa1565b80156105685780601f1061053f57610100808354040283529160200191610568565b820191905f5260205f20905b81548152906001019060200180831161054b57829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020180546105a290611aa1565b80601f01602080910402602001604051908101604052809291908181526020018280546105ce90611aa1565b80156106195780601f106105f057610100808354040283529160200191610619565b820191905f5260205f20905b8154815290600101906020018083116105fc57829003601f168201915b50505050509080600301805461062e90611aa1565b80601f016020809104026020016040519081016040528092919081815260200182805461065a90611aa1565b80156106a55780601f1061067c576101008083540402835291602001916106a5565b820191905f5260205f20905b81548152906001019060200180831161068857829003601f168201915b505050505090806004015f9054906101000a900460ff16908060040160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154908060060154908060070154905089565b805f60018260405161070b91906119d0565b9081526020016040518091039020604051806101200160405290815f8201805461073490611aa1565b80601f016020809104026020016040519081016040528092919081815260200182805461076090611aa1565b80156107ab5780601f10610782576101008083540402835291602001916107ab565b820191905f5260205f20905b81548152906001019060200180831161078e57829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461081990611aa1565b80601f016020809104026020016040519081016040528092919081815260200182805461084590611aa1565b80156108905780601f1061086757610100808354040283529160200191610890565b820191905f5260205f20905b81548152906001019060200180831161087357829003601f168201915b505050505081526020016003820180546108a990611aa1565b80601f01602080910402602001604051908101604052809291908181526020018280546108d590611aa1565b80156109205780601f106108f757610100808354040283529160200191610920565b820191905f5260205f20905b81548152906001019060200180831161090357829003601f168201915b50505050508152602001600482015f9054906101000a900460ff16600381111561094d5761094c611782565b5b600381111561095f5761095e611782565b5b81526020016004820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582015481526020016006820154815260200160078201548152505090503373ffffffffffffffffffffffffffffffffffffffff16600183604051610a0091906119d0565b908152602001604051809103902060040160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a88576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7f90611dcd565b60405180910390fd5b60016003811115610a9c57610a9b611782565b5b600183604051610aac91906119d0565b90815260200160405180910390206004015f9054906101000a900460ff166003811115610adc57610adb611782565b5b14610b1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1390611e35565b60405180910390fd5b5f835111610b5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5690611910565b60405180910390fd5b5f600184604051610b7091906119d0565b908152602001604051809103902090506003816004015f6101000a81548160ff02191690836003811115610ba757610ba6611782565b5b02179055505f8160040160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504281600701819055503373ffffffffffffffffffffffffffffffffffffffff167f9ac4bafed1f4d0b87ae4f9fed31daee22a04893aa69e866007d5bc926f0ff17d85604051610c3e9190611e53565b60405180910390a250505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b815f600182604051610c8291906119d0565b9081526020016040518091039020604051806101200160405290815f82018054610cab90611aa1565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd790611aa1565b8015610d225780601f10610cf957610100808354040283529160200191610d22565b820191905f5260205f20905b815481529060010190602001808311610d0557829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054610d9090611aa1565b80601f0160208091040260200160405190810160405280929190818152602001828054610dbc90611aa1565b8015610e075780601f10610dde57610100808354040283529160200191610e07565b820191905f5260205f20905b815481529060010190602001808311610dea57829003601f168201915b50505050508152602001600382018054610e2090611aa1565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4c90611aa1565b8015610e975780601f10610e6e57610100808354040283529160200191610e97565b820191905f5260205f20905b815481529060010190602001808311610e7a57829003601f168201915b50505050508152602001600482015f9054906101000a900460ff166003811115610ec457610ec3611782565b5b6003811115610ed657610ed5611782565b5b81526020016004820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582015481526020016006820154815260200160078201548152505090503373ffffffffffffffffffffffffffffffffffffffff16600183604051610f7791906119d0565b908152602001604051809103902060040160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610fff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff690611dcd565b60405180910390fd5b6001600381111561101357611012611782565b5b60018360405161102391906119d0565b90815260200160405180910390206004015f9054906101000a900460ff16600381111561105357611052611782565b5b14611093576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108a90611e35565b60405180910390fd5b5f8451116110d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110cd90611910565b60405180910390fd5b5f6001856040516110e791906119d0565b908152602001604051809103902090506002816004015f6101000a81548160ff0219169083600381111561111e5761111d611782565b5b0217905550838160030190816111349190611c71565b504281600701819055503373ffffffffffffffffffffffffffffffffffffffff167f1e6cd60eb3a95fa354fda400577a90e42e85f27a8031114340936da2b89c532f8686604051611186929190611e73565b60405180910390a25f8160040160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682600501546040516111db90611ed5565b5f6040518083038185875af1925050503d805f8114611215576040519150601f19603f3d011682016040523d82523d5f602084013e61121a565b606091505b505090508061125e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125590611f33565b60405180910390fd5b505050505050565b5f8151116112a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a090611910565b60405180910390fd5b5f6001826040516112ba91906119d0565b908152602001604051809103902090505f60038111156112dd576112dc611782565b5b816004015f9054906101000a900460ff166003811115611300576112ff611782565b5b148061133f575060038081111561131a57611319611782565b5b816004015f9054906101000a900460ff16600381111561133d5761133c611782565b5b145b61137e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137590611f9b565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16816001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361140f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161140690612029565b60405180910390fd5b6001816004015f6101000a81548160ff0219169083600381111561143657611435611782565b5b0217905550338160040160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504281600701819055503373ffffffffffffffffffffffffffffffffffffffff167f54e973d1f7ca7aab25c55b07f06ce112eaa3d4c1563914ac7047182ccaaef4f7836040516114cd9190611e53565b60405180910390a25050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611538826114f2565b810181811067ffffffffffffffff8211171561155757611556611502565b5b80604052505050565b5f6115696114d9565b9050611575828261152f565b919050565b5f67ffffffffffffffff82111561159457611593611502565b5b61159d826114f2565b9050602081019050919050565b828183375f83830152505050565b5f6115ca6115c58461157a565b611560565b9050828152602081018484840111156115e6576115e56114ee565b5b6115f18482856115aa565b509392505050565b5f82601f83011261160d5761160c6114ea565b5b813561161d8482602086016115b8565b91505092915050565b5f5f6040838503121561163c5761163b6114e2565b5b5f83013567ffffffffffffffff811115611659576116586114e6565b5b611665858286016115f9565b925050602083013567ffffffffffffffff811115611686576116856114e6565b5b611692858286016115f9565b9150509250929050565b5f602082840312156116b1576116b06114e2565b5b5f82013567ffffffffffffffff8111156116ce576116cd6114e6565b5b6116da848285016115f9565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f611715826116e3565b61171f81856116ed565b935061172f8185602086016116fd565b611738816114f2565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61176c82611743565b9050919050565b61177c81611762565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600481106117c0576117bf611782565b5b50565b5f8190506117d0826117af565b919050565b5f6117df826117c3565b9050919050565b6117ef816117d5565b82525050565b5f819050919050565b611807816117f5565b82525050565b5f610120820190508181035f830152611826818c61170b565b9050611835602083018b611773565b8181036040830152611847818a61170b565b9050818103606083015261185b818961170b565b905061186a60808301886117e6565b61187760a0830187611773565b61188460c08301866117fe565b61189160e08301856117fe565b61189f6101008301846117fe565b9a9950505050505050505050565b5f6020820190506118c05f830184611773565b92915050565b7f555549442063616e6e6f7420626520656d7074790000000000000000000000005f82015250565b5f6118fa6014836116ed565b9150611905826118c6565b602082019050919050565b5f6020820190508181035f830152611927816118ee565b9050919050565b7f496e76616c696420616d6f756e740000000000000000000000000000000000005f82015250565b5f611962600e836116ed565b915061196d8261192e565b602082019050919050565b5f6020820190508181035f83015261198f81611956565b9050919050565b5f81905092915050565b5f6119aa826116e3565b6119b48185611996565b93506119c48185602086016116fd565b80840191505092915050565b5f6119db82846119a0565b915081905092915050565b7f5461736b20776974682074686973205555494420616c726561647920657869735f8201527f7473000000000000000000000000000000000000000000000000000000000000602082015250565b5f611a406022836116ed565b9150611a4b826119e6565b604082019050919050565b5f6020820190508181035f830152611a6d81611a34565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611ab857607f821691505b602082108103611acb57611aca611a74565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611b2d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611af2565b611b378683611af2565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611b72611b6d611b68846117f5565b611b4f565b6117f5565b9050919050565b5f819050919050565b611b8b83611b58565b611b9f611b9782611b79565b848454611afe565b825550505050565b5f5f905090565b611bb6611ba7565b611bc1818484611b82565b505050565b5b81811015611be457611bd95f82611bae565b600181019050611bc7565b5050565b601f821115611c2957611bfa81611ad1565b611c0384611ae3565b81016020851015611c12578190505b611c26611c1e85611ae3565b830182611bc6565b50505b505050565b5f82821c905092915050565b5f611c495f1984600802611c2e565b1980831691505092915050565b5f611c618383611c3a565b9150826002028217905092915050565b611c7a826116e3565b67ffffffffffffffff811115611c9357611c92611502565b5b611c9d8254611aa1565b611ca8828285611be8565b5f60209050601f831160018114611cd9575f8415611cc7578287015190505b611cd18582611c56565b865550611d38565b601f198416611ce786611ad1565b5f5b82811015611d0e57848901518255600182019150602085019450602081019050611ce9565b86831015611d2b5784890151611d27601f891682611c3a565b8355505b6001600288020188555050505b505050505050565b5f6060820190508181035f830152611d58818661170b565b90508181036020830152611d6c818561170b565b9050611d7b60408301846117fe565b949350505050565b7f43616c6c6572206973206e6f74207461736b2061737369676e656500000000005f82015250565b5f611db7601b836116ed565b9150611dc282611d83565b602082019050919050565b5f6020820190508181035f830152611de481611dab565b9050919050565b7f5461736b206e6f7420696e20436c61696d6564207374617465000000000000005f82015250565b5f611e1f6019836116ed565b9150611e2a82611deb565b602082019050919050565b5f6020820190508181035f830152611e4c81611e13565b9050919050565b5f6020820190508181035f830152611e6b818461170b565b905092915050565b5f6040820190508181035f830152611e8b818561170b565b90508181036020830152611e9f818461170b565b90509392505050565b5f81905092915050565b50565b5f611ec05f83611ea8565b9150611ecb82611eb2565b5f82019050919050565b5f611edf82611eb5565b9150819050919050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f611f1d600f836116ed565b9150611f2882611ee9565b602082019050919050565b5f6020820190508181035f830152611f4a81611f11565b9050919050565b7f4e6f742070656e64696e67206f72206661696c656400000000000000000000005f82015250565b5f611f856015836116ed565b9150611f9082611f51565b602082019050919050565b5f6020820190508181035f830152611fb281611f79565b9050919050565b7f5375626d69747465722063616e6e6f7420636c61696d207468656972206f776e5f8201527f207461736b000000000000000000000000000000000000000000000000000000602082015250565b5f6120136025836116ed565b915061201e82611fb9565b604082019050919050565b5f6020820190508181035f83015261204081612007565b905091905056fea2646970667358221220a803b36f62aee9963d0749ab08f9de52c093bbe0141ef099f449bc2fe9ffb86764736f6c634300081e0033",
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
// Solidity: function tasks(string ) view returns(string uuid, address submitter, string input, string output, uint8 status, address assignee, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	Uuid      string
	Submitter common.Address
	Input     string
	Output    string
	Status    uint8
	Assignee  common.Address
	Amount    *big.Int
	CreatedAt *big.Int
	UpdatedAt *big.Int
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
		CreatedAt *big.Int
		UpdatedAt *big.Int
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
	outstruct.CreatedAt = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, address submitter, string input, string output, uint8 status, address assignee, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractSession) Tasks(arg0 string) (struct {
	Uuid      string
	Submitter common.Address
	Input     string
	Output    string
	Status    uint8
	Assignee  common.Address
	Amount    *big.Int
	CreatedAt *big.Int
	UpdatedAt *big.Int
}, error) {
	return _Contract.Contract.Tasks(&_Contract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, address submitter, string input, string output, uint8 status, address assignee, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractCallerSession) Tasks(arg0 string) (struct {
	Uuid      string
	Submitter common.Address
	Input     string
	Output    string
	Status    uint8
	Assignee  common.Address
	Amount    *big.Int
	CreatedAt *big.Int
	UpdatedAt *big.Int
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
// Solidity: function complete(string _uuid, string _output) returns()
func (_Contract *ContractTransactor) Complete(opts *bind.TransactOpts, _uuid string, _output string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "complete", _uuid, _output)
}

// Complete is a paid mutator transaction binding the contract method 0xeca4cbe6.
//
// Solidity: function complete(string _uuid, string _output) returns()
func (_Contract *ContractSession) Complete(_uuid string, _output string) (*types.Transaction, error) {
	return _Contract.Contract.Complete(&_Contract.TransactOpts, _uuid, _output)
}

// Complete is a paid mutator transaction binding the contract method 0xeca4cbe6.
//
// Solidity: function complete(string _uuid, string _output) returns()
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

// ContractTaskClaimedIterator is returned from FilterTaskClaimed and is used to iterate over the raw logs and unpacked data for TaskClaimed events raised by the Contract contract.
type ContractTaskClaimedIterator struct {
	Event *ContractTaskClaimed // Event containing the contract specifics and raw log

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
func (it *ContractTaskClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskClaimed)
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
		it.Event = new(ContractTaskClaimed)
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
func (it *ContractTaskClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskClaimed represents a TaskClaimed event raised by the Contract contract.
type ContractTaskClaimed struct {
	Uuid     string
	Assignee common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskClaimed is a free log retrieval operation binding the contract event 0x54e973d1f7ca7aab25c55b07f06ce112eaa3d4c1563914ac7047182ccaaef4f7.
//
// Solidity: event TaskClaimed(string uuid, address indexed assignee)
func (_Contract *ContractFilterer) FilterTaskClaimed(opts *bind.FilterOpts, assignee []common.Address) (*ContractTaskClaimedIterator, error) {

	var assigneeRule []interface{}
	for _, assigneeItem := range assignee {
		assigneeRule = append(assigneeRule, assigneeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskClaimed", assigneeRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskClaimedIterator{contract: _Contract.contract, event: "TaskClaimed", logs: logs, sub: sub}, nil
}

// WatchTaskClaimed is a free log subscription operation binding the contract event 0x54e973d1f7ca7aab25c55b07f06ce112eaa3d4c1563914ac7047182ccaaef4f7.
//
// Solidity: event TaskClaimed(string uuid, address indexed assignee)
func (_Contract *ContractFilterer) WatchTaskClaimed(opts *bind.WatchOpts, sink chan<- *ContractTaskClaimed, assignee []common.Address) (event.Subscription, error) {

	var assigneeRule []interface{}
	for _, assigneeItem := range assignee {
		assigneeRule = append(assigneeRule, assigneeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskClaimed", assigneeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskClaimed)
				if err := _Contract.contract.UnpackLog(event, "TaskClaimed", log); err != nil {
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

// ParseTaskClaimed is a log parse operation binding the contract event 0x54e973d1f7ca7aab25c55b07f06ce112eaa3d4c1563914ac7047182ccaaef4f7.
//
// Solidity: event TaskClaimed(string uuid, address indexed assignee)
func (_Contract *ContractFilterer) ParseTaskClaimed(log types.Log) (*ContractTaskClaimed, error) {
	event := new(ContractTaskClaimed)
	if err := _Contract.contract.UnpackLog(event, "TaskClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
// Solidity: event TaskCompleted(string uuid, address indexed assignee, string output)
func (_Contract *ContractFilterer) FilterTaskCompleted(opts *bind.FilterOpts, assignee []common.Address) (*ContractTaskCompletedIterator, error) {

	var assigneeRule []interface{}
	for _, assigneeItem := range assignee {
		assigneeRule = append(assigneeRule, assigneeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskCompleted", assigneeRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskCompletedIterator{contract: _Contract.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x1e6cd60eb3a95fa354fda400577a90e42e85f27a8031114340936da2b89c532f.
//
// Solidity: event TaskCompleted(string uuid, address indexed assignee, string output)
func (_Contract *ContractFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractTaskCompleted, assignee []common.Address) (event.Subscription, error) {

	var assigneeRule []interface{}
	for _, assigneeItem := range assignee {
		assigneeRule = append(assigneeRule, assigneeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskCompleted", assigneeRule)
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
// Solidity: event TaskCompleted(string uuid, address indexed assignee, string output)
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
	Uuid      string
	Submitter common.Address
	Input     string
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x9bc0fc01301fa260e9f61a4f2be24187c8a9f7cebcaf3a83251af30a355f62d2.
//
// Solidity: event TaskCreated(string uuid, address indexed submitter, string input, uint256 amount)
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

// WatchTaskCreated is a free log subscription operation binding the contract event 0x9bc0fc01301fa260e9f61a4f2be24187c8a9f7cebcaf3a83251af30a355f62d2.
//
// Solidity: event TaskCreated(string uuid, address indexed submitter, string input, uint256 amount)
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

// ParseTaskCreated is a log parse operation binding the contract event 0x9bc0fc01301fa260e9f61a4f2be24187c8a9f7cebcaf3a83251af30a355f62d2.
//
// Solidity: event TaskCreated(string uuid, address indexed submitter, string input, uint256 amount)
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
	Uuid              string
	AssigneeWhoFailed common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTaskFailed is a free log retrieval operation binding the contract event 0x9ac4bafed1f4d0b87ae4f9fed31daee22a04893aa69e866007d5bc926f0ff17d.
//
// Solidity: event TaskFailed(string uuid, address indexed assigneeWhoFailed)
func (_Contract *ContractFilterer) FilterTaskFailed(opts *bind.FilterOpts, assigneeWhoFailed []common.Address) (*ContractTaskFailedIterator, error) {

	var assigneeWhoFailedRule []interface{}
	for _, assigneeWhoFailedItem := range assigneeWhoFailed {
		assigneeWhoFailedRule = append(assigneeWhoFailedRule, assigneeWhoFailedItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TaskFailed", assigneeWhoFailedRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskFailedIterator{contract: _Contract.contract, event: "TaskFailed", logs: logs, sub: sub}, nil
}

// WatchTaskFailed is a free log subscription operation binding the contract event 0x9ac4bafed1f4d0b87ae4f9fed31daee22a04893aa69e866007d5bc926f0ff17d.
//
// Solidity: event TaskFailed(string uuid, address indexed assigneeWhoFailed)
func (_Contract *ContractFilterer) WatchTaskFailed(opts *bind.WatchOpts, sink chan<- *ContractTaskFailed, assigneeWhoFailed []common.Address) (event.Subscription, error) {

	var assigneeWhoFailedRule []interface{}
	for _, assigneeWhoFailedItem := range assigneeWhoFailed {
		assigneeWhoFailedRule = append(assigneeWhoFailedRule, assigneeWhoFailedItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TaskFailed", assigneeWhoFailedRule)
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

// ParseTaskFailed is a log parse operation binding the contract event 0x9ac4bafed1f4d0b87ae4f9fed31daee22a04893aa69e866007d5bc926f0ff17d.
//
// Solidity: event TaskFailed(string uuid, address indexed assigneeWhoFailed)
func (_Contract *ContractFilterer) ParseTaskFailed(log types.Log) (*ContractTaskFailed, error) {
	event := new(ContractTaskFailed)
	if err := _Contract.contract.UnpackLog(event, "TaskFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
