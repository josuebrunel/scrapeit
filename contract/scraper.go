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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"}],\"name\":\"TaskClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assigneeWhoFailed\",\"type\":\"address\"}],\"name\":\"TaskFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"}],\"name\":\"complete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"fail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506124a18061005b5f395ff3fe60806040526004361061007a575f3560e01c806378122f3a1161004d57806378122f3a146101305780638da5cb5b14610158578063eca4cbe614610182578063f3fe12c9146101aa5761007a565b806316c38b3c1461007e578063198e2b8a146100a657806358709cf2146100c25780635c975abb14610106575b5f5ffd5b348015610089575f5ffd5b506100a4600480360381019061009f91906117eb565b6101d2565b005b6100c060048036038101906100bb9190611952565b61027c565b005b3480156100cd575f5ffd5b506100e860048036038101906100e391906119c8565b610687565b6040516100fd99989796959493929190611b39565b60405180910390f35b348015610111575f5ffd5b5061011a6108c6565b6040516101279190611be8565b60405180910390f35b34801561013b575f5ffd5b50610156600480360381019061015191906119c8565b6108d8565b005b348015610163575f5ffd5b5061016c610e7a565b6040516101799190611c01565b60405180910390f35b34801561018d575f5ffd5b506101a860048036038101906101a39190611952565b610e9e565b005b3480156101b5575f5ffd5b506101d060048036038101906101cb91906119c8565b6114e3565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610260576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025790611c64565b60405180910390fd5b805f60146101000a81548160ff02191690831515021790555050565b5f60149054906101000a900460ff16156102cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c290611ccc565b60405180910390fd5b5f82511161030e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030590611d34565b60405180910390fd5b5f3411610350576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034790611d9c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff166001836040516103779190611df4565b90815260200160405180910390206002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146103fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f590611e7a565b60405180910390fd5b5f600381111561041157610410611a6f565b5b6001836040516104219190611df4565b90815260200160405180910390206001015f9054906101000a900460ff16600381111561045157610450611a6f565b5b1461045a575f5ffd5b5f6040518061012001604052808481526020015f60038111156104805761047f611a6f565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160405180602001604052805f8152508152602001348152602001428152602001428152509050806001825f01516040516104fd9190611df4565b90815260200160405180910390205f820151815f01908161051e9190612095565b506020820151816001015f6101000a81548160ff0219169083600381111561054957610548611a6f565b5b021790555060408201518160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030190816105f09190612095565b5060a08201518160040190816106069190612095565b5060c0820151816005015560e0820151816006015561010082015181600701559050503373ffffffffffffffffffffffffffffffffffffffff167f9bc0fc01301fa260e9f61a4f2be24187c8a9f7cebcaf3a83251af30a355f62d2825f015183608001513460405161067a93929190612164565b60405180910390a2505050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546106be90611ec5565b80601f01602080910402602001604051908101604052809291908181526020018280546106ea90611ec5565b80156107355780601f1061070c57610100808354040283529160200191610735565b820191905f5260205f20905b81548152906001019060200180831161071857829003601f168201915b505050505090806001015f9054906101000a900460ff16908060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030180546107a790611ec5565b80601f01602080910402602001604051908101604052809291908181526020018280546107d390611ec5565b801561081e5780601f106107f55761010080835404028352916020019161081e565b820191905f5260205f20905b81548152906001019060200180831161080157829003601f168201915b50505050509080600401805461083390611ec5565b80601f016020809104026020016040519081016040528092919081815260200182805461085f90611ec5565b80156108aa5780601f10610881576101008083540402835291602001916108aa565b820191905f5260205f20905b81548152906001019060200180831161088d57829003601f168201915b5050505050908060050154908060060154908060070154905089565b5f60149054906101000a900460ff1681565b805f6001826040516108ea9190611df4565b9081526020016040518091039020604051806101200160405290815f8201805461091390611ec5565b80601f016020809104026020016040519081016040528092919081815260200182805461093f90611ec5565b801561098a5780601f106109615761010080835404028352916020019161098a565b820191905f5260205f20905b81548152906001019060200180831161096d57829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1660038111156109b7576109b6611a6f565b5b60038111156109c9576109c8611a6f565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054610a8890611ec5565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab490611ec5565b8015610aff5780601f10610ad657610100808354040283529160200191610aff565b820191905f5260205f20905b815481529060010190602001808311610ae257829003601f168201915b50505050508152602001600482018054610b1890611ec5565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4490611ec5565b8015610b8f5780601f10610b6657610100808354040283529160200191610b8f565b820191905f5260205f20905b815481529060010190602001808311610b7257829003601f168201915b50505050508152602001600582015481526020016006820154815260200160078201548152505090503373ffffffffffffffffffffffffffffffffffffffff16600183604051610bdf9190611df4565b908152602001604051809103902060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610c67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5e906121f1565b60405180910390fd5b60016003811115610c7b57610c7a611a6f565b5b600183604051610c8b9190611df4565b90815260200160405180910390206001015f9054906101000a900460ff166003811115610cbb57610cba611a6f565b5b14610cfb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf290612259565b60405180910390fd5b5f60149054906101000a900460ff1615610d4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4190611ccc565b60405180910390fd5b5f835111610d8d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8490611d34565b60405180910390fd5b5f600184604051610d9e9190611df4565b908152602001604051809103902090506003816001015f6101000a81548160ff02191690836003811115610dd557610dd4611a6f565b5b02179055505f8160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504281600701819055503373ffffffffffffffffffffffffffffffffffffffff167f9ac4bafed1f4d0b87ae4f9fed31daee22a04893aa69e866007d5bc926f0ff17d85604051610e6c9190612277565b60405180910390a250505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b815f600182604051610eb09190611df4565b9081526020016040518091039020604051806101200160405290815f82018054610ed990611ec5565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0590611ec5565b8015610f505780601f10610f2757610100808354040283529160200191610f50565b820191905f5260205f20905b815481529060010190602001808311610f3357829003601f168201915b50505050508152602001600182015f9054906101000a900460ff166003811115610f7d57610f7c611a6f565b5b6003811115610f8f57610f8e611a6f565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805461104e90611ec5565b80601f016020809104026020016040519081016040528092919081815260200182805461107a90611ec5565b80156110c55780601f1061109c576101008083540402835291602001916110c5565b820191905f5260205f20905b8154815290600101906020018083116110a857829003601f168201915b505050505081526020016004820180546110de90611ec5565b80601f016020809104026020016040519081016040528092919081815260200182805461110a90611ec5565b80156111555780601f1061112c57610100808354040283529160200191611155565b820191905f5260205f20905b81548152906001019060200180831161113857829003601f168201915b50505050508152602001600582015481526020016006820154815260200160078201548152505090503373ffffffffffffffffffffffffffffffffffffffff166001836040516111a59190611df4565b908152602001604051809103902060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461122d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611224906121f1565b60405180910390fd5b6001600381111561124157611240611a6f565b5b6001836040516112519190611df4565b90815260200160405180910390206001015f9054906101000a900460ff16600381111561128157611280611a6f565b5b146112c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112b890612259565b60405180910390fd5b5f60149054906101000a900460ff1615611310576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130790611ccc565b60405180910390fd5b5f845111611353576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161134a90611d34565b60405180910390fd5b5f6001856040516113649190611df4565b908152602001604051809103902090506002816001015f6101000a81548160ff0219169083600381111561139b5761139a611a6f565b5b0217905550838160040190816113b19190612095565b504281600701819055503373ffffffffffffffffffffffffffffffffffffffff167f1e6cd60eb3a95fa354fda400577a90e42e85f27a8031114340936da2b89c532f8686604051611403929190612297565b60405180910390a25f8160010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168260050154604051611458906122f9565b5f6040518083038185875af1925050503d805f8114611492576040519150601f19603f3d011682016040523d82523d5f602084013e611497565b606091505b50509050806114db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114d290612357565b60405180910390fd5b505050505050565b5f60149054906101000a900460ff1615611532576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161152990611ccc565b60405180910390fd5b5f815111611575576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161156c90611d34565b60405180910390fd5b5f6001826040516115869190611df4565b908152602001604051809103902090505f60038111156115a9576115a8611a6f565b5b816001015f9054906101000a900460ff1660038111156115cc576115cb611a6f565b5b148061160b57506003808111156115e6576115e5611a6f565b5b816001015f9054906101000a900460ff16600381111561160957611608611a6f565b5b145b61164a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611641906123bf565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16816002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036116db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116d29061244d565b60405180910390fd5b6001816001015f6101000a81548160ff0219169083600381111561170257611701611a6f565b5b0217905550338160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504281600701819055503373ffffffffffffffffffffffffffffffffffffffff167f54e973d1f7ca7aab25c55b07f06ce112eaa3d4c1563914ac7047182ccaaef4f7836040516117999190612277565b60405180910390a25050565b5f604051905090565b5f5ffd5b5f5ffd5b5f8115159050919050565b6117ca816117b6565b81146117d4575f5ffd5b50565b5f813590506117e5816117c1565b92915050565b5f60208284031215611800576117ff6117ae565b5b5f61180d848285016117d7565b91505092915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6118648261181e565b810181811067ffffffffffffffff821117156118835761188261182e565b5b80604052505050565b5f6118956117a5565b90506118a1828261185b565b919050565b5f67ffffffffffffffff8211156118c0576118bf61182e565b5b6118c98261181e565b9050602081019050919050565b828183375f83830152505050565b5f6118f66118f1846118a6565b61188c565b9050828152602081018484840111156119125761191161181a565b5b61191d8482856118d6565b509392505050565b5f82601f83011261193957611938611816565b5b81356119498482602086016118e4565b91505092915050565b5f5f60408385031215611968576119676117ae565b5b5f83013567ffffffffffffffff811115611985576119846117b2565b5b61199185828601611925565b925050602083013567ffffffffffffffff8111156119b2576119b16117b2565b5b6119be85828601611925565b9150509250929050565b5f602082840312156119dd576119dc6117ae565b5b5f82013567ffffffffffffffff8111156119fa576119f96117b2565b5b611a0684828501611925565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f611a4182611a0f565b611a4b8185611a19565b9350611a5b818560208601611a29565b611a648161181e565b840191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60048110611aad57611aac611a6f565b5b50565b5f819050611abd82611a9c565b919050565b5f611acc82611ab0565b9050919050565b611adc81611ac2565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611b0b82611ae2565b9050919050565b611b1b81611b01565b82525050565b5f819050919050565b611b3381611b21565b82525050565b5f610120820190508181035f830152611b52818c611a37565b9050611b61602083018b611ad3565b611b6e604083018a611b12565b611b7b6060830189611b12565b8181036080830152611b8d8188611a37565b905081810360a0830152611ba18187611a37565b9050611bb060c0830186611b2a565b611bbd60e0830185611b2a565b611bcb610100830184611b2a565b9a9950505050505050505050565b611be2816117b6565b82525050565b5f602082019050611bfb5f830184611bd9565b92915050565b5f602082019050611c145f830184611b12565b92915050565b7f43616c6c6572206973206e6f74206f776e6572000000000000000000000000005f82015250565b5f611c4e601383611a19565b9150611c5982611c1a565b602082019050919050565b5f6020820190508181035f830152611c7b81611c42565b9050919050565b7f436f6e74726163742069732070617573656400000000000000000000000000005f82015250565b5f611cb6601283611a19565b9150611cc182611c82565b602082019050919050565b5f6020820190508181035f830152611ce381611caa565b9050919050565b7f555549442063616e6e6f7420626520656d7074790000000000000000000000005f82015250565b5f611d1e601483611a19565b9150611d2982611cea565b602082019050919050565b5f6020820190508181035f830152611d4b81611d12565b9050919050565b7f496e76616c696420616d6f756e740000000000000000000000000000000000005f82015250565b5f611d86600e83611a19565b9150611d9182611d52565b602082019050919050565b5f6020820190508181035f830152611db381611d7a565b9050919050565b5f81905092915050565b5f611dce82611a0f565b611dd88185611dba565b9350611de8818560208601611a29565b80840191505092915050565b5f611dff8284611dc4565b915081905092915050565b7f5461736b20776974682074686973205555494420616c726561647920657869735f8201527f7473000000000000000000000000000000000000000000000000000000000000602082015250565b5f611e64602283611a19565b9150611e6f82611e0a565b604082019050919050565b5f6020820190508181035f830152611e9181611e58565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611edc57607f821691505b602082108103611eef57611eee611e98565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611f517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611f16565b611f5b8683611f16565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611f96611f91611f8c84611b21565b611f73565b611b21565b9050919050565b5f819050919050565b611faf83611f7c565b611fc3611fbb82611f9d565b848454611f22565b825550505050565b5f5f905090565b611fda611fcb565b611fe5818484611fa6565b505050565b5b8181101561200857611ffd5f82611fd2565b600181019050611feb565b5050565b601f82111561204d5761201e81611ef5565b61202784611f07565b81016020851015612036578190505b61204a61204285611f07565b830182611fea565b50505b505050565b5f82821c905092915050565b5f61206d5f1984600802612052565b1980831691505092915050565b5f612085838361205e565b9150826002028217905092915050565b61209e82611a0f565b67ffffffffffffffff8111156120b7576120b661182e565b5b6120c18254611ec5565b6120cc82828561200c565b5f60209050601f8311600181146120fd575f84156120eb578287015190505b6120f5858261207a565b86555061215c565b601f19841661210b86611ef5565b5f5b828110156121325784890151825560018201915060208501945060208101905061210d565b8683101561214f578489015161214b601f89168261205e565b8355505b6001600288020188555050505b505050505050565b5f6060820190508181035f83015261217c8186611a37565b905081810360208301526121908185611a37565b905061219f6040830184611b2a565b949350505050565b7f43616c6c6572206973206e6f74207461736b2061737369676e656500000000005f82015250565b5f6121db601b83611a19565b91506121e6826121a7565b602082019050919050565b5f6020820190508181035f830152612208816121cf565b9050919050565b7f5461736b206e6f7420696e20436c61696d6564207374617465000000000000005f82015250565b5f612243601983611a19565b915061224e8261220f565b602082019050919050565b5f6020820190508181035f83015261227081612237565b9050919050565b5f6020820190508181035f83015261228f8184611a37565b905092915050565b5f6040820190508181035f8301526122af8185611a37565b905081810360208301526122c38184611a37565b90509392505050565b5f81905092915050565b50565b5f6122e45f836122cc565b91506122ef826122d6565b5f82019050919050565b5f612303826122d9565b9150819050919050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f612341600f83611a19565b915061234c8261230d565b602082019050919050565b5f6020820190508181035f83015261236e81612335565b9050919050565b7f4e6f742070656e64696e67206f72206661696c656400000000000000000000005f82015250565b5f6123a9601583611a19565b91506123b482612375565b602082019050919050565b5f6020820190508181035f8301526123d68161239d565b9050919050565b7f5375626d69747465722063616e6e6f7420636c61696d207468656972206f776e5f8201527f207461736b000000000000000000000000000000000000000000000000000000602082015250565b5f612437602583611a19565b9150612442826123dd565b604082019050919050565b5f6020820190508181035f8301526124648161242b565b905091905056fea2646970667358221220ee65c2e354600b1c0fe90c6f931a2be5d94854739e9b31696caf18ec459907c764736f6c634300081e0033",
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
// Solidity: function tasks(string ) view returns(string uuid, uint8 status, address assignee, address submitter, string input, string output, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	Uuid      string
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
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Assignee = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Submitter = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Input = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Output = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string uuid, uint8 status, address assignee, address submitter, string input, string output, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractSession) Tasks(arg0 string) (struct {
	Uuid      string
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
// Solidity: function tasks(string ) view returns(string uuid, uint8 status, address assignee, address submitter, string input, string output, uint256 amount, uint256 created_at, uint256 updated_at)
func (_Contract *ContractCallerSession) Tasks(arg0 string) (struct {
	Uuid      string
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
