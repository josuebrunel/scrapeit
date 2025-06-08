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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"}],\"name\":\"TaskClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"user_id\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"name\":\"TaskProcessed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"}],\"name\":\"complete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_user_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uuid\",\"type\":\"string\"}],\"name\":\"fail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"user_id\",\"type\":\"string\"},{\"internalType\":\"enumScraper.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506127b68061005b5f395ff3fe60806040526004361061007a575f3560e01c806378122f3a1161004d57806378122f3a146101315780638da5cb5b14610159578063eca4cbe614610183578063f3fe12c9146101ab5761007a565b806316c38b3c1461007e57806358709cf2146100a65780635c975abb146100eb5780635d28560a14610115575b5f5ffd5b348015610089575f5ffd5b506100a4600480360381019061009f91906119c2565b6101d3565b005b3480156100b1575f5ffd5b506100cc60048036038101906100c79190611b29565b61027d565b6040516100e29a99989796959493929190611c9a565b60405180910390f35b3480156100f6575f5ffd5b506100ff610548565b60405161010c9190611d5f565b60405180910390f35b61012f600480360381019061012a9190611d78565b61055a565b005b34801561013c575f5ffd5b5061015760048036038101906101529190611b29565b610989565b005b348015610164575f5ffd5b5061016d610fbe565b60405161017a9190611e1c565b60405180910390f35b34801561018e575f5ffd5b506101a960048036038101906101a49190611e35565b610fe2565b005b3480156101b6575f5ffd5b506101d160048036038101906101cc9190611b29565b6116ba565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610261576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025890611ef5565b60405180910390fd5b805f60146101000a81548160ff02191690831515021790555050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546102b490611f40565b80601f01602080910402602001604051908101604052809291908181526020018280546102e090611f40565b801561032b5780601f106103025761010080835404028352916020019161032b565b820191905f5260205f20905b81548152906001019060200180831161030e57829003601f168201915b50505050509080600101805461034090611f40565b80601f016020809104026020016040519081016040528092919081815260200182805461036c90611f40565b80156103b75780601f1061038e576101008083540402835291602001916103b7565b820191905f5260205f20905b81548152906001019060200180831161039a57829003601f168201915b505050505090806002015f9054906101000a900460ff16908060020160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600401805461042990611f40565b80601f016020809104026020016040519081016040528092919081815260200182805461045590611f40565b80156104a05780601f10610477576101008083540402835291602001916104a0565b820191905f5260205f20905b81548152906001019060200180831161048357829003601f168201915b5050505050908060050180546104b590611f40565b80601f01602080910402602001604051908101604052809291908181526020018280546104e190611f40565b801561052c5780601f106105035761010080835404028352916020019161052c565b820191905f5260205f20905b81548152906001019060200180831161050f57829003601f168201915b505050505090806006015490806007015490806008015490508a565b5f60149054906101000a900460ff1681565b5f60149054906101000a900460ff16156105a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a090611fba565b60405180910390fd5b5f8351116105ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e390612022565b60405180910390fd5b5f341161062e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106259061208a565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660018460405161065591906120e2565b90815260200160405180910390206003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d390612168565b60405180910390fd5b5f60038111156106ef576106ee611bd0565b5b6001846040516106ff91906120e2565b90815260200160405180910390206002015f9054906101000a900460ff16600381111561072f5761072e611bd0565b5b14610738575f5ffd5b5f6040518061014001604052808581526020018481526020015f600381111561076457610763611bd0565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160405180602001604052805f8152508152602001348152602001428152602001428152509050806001825f01516040516107e191906120e2565b90815260200160405180910390205f820151815f0190816108029190612326565b5060208201518160010190816108189190612326565b506040820151816002015f6101000a81548160ff0219169083600381111561084357610842611bd0565b5b021790555060608201518160020160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040190816108ea9190612326565b5060c08201518160050190816109009190612326565b5060e08201518160060155610100820151816007015561012082015181600801559050503373ffffffffffffffffffffffffffffffffffffffff167f62b751e0cbad6fdb8b8e78702ad231910ed198c3d2e6c6b2a2a93f1a02dc78f2825f015183602001518460a001513460405161097b94939291906123f5565b60405180910390a250505050565b805f60018260405161099b91906120e2565b9081526020016040518091039020604051806101400160405290815f820180546109c490611f40565b80601f01602080910402602001604051908101604052809291908181526020018280546109f090611f40565b8015610a3b5780601f10610a1257610100808354040283529160200191610a3b565b820191905f5260205f20905b815481529060010190602001808311610a1e57829003601f168201915b50505050508152602001600182018054610a5490611f40565b80601f0160208091040260200160405190810160405280929190818152602001828054610a8090611f40565b8015610acb5780601f10610aa257610100808354040283529160200191610acb565b820191905f5260205f20905b815481529060010190602001808311610aae57829003601f168201915b50505050508152602001600282015f9054906101000a900460ff166003811115610af857610af7611bd0565b5b6003811115610b0a57610b09611bd0565b5b81526020016002820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054610bc990611f40565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf590611f40565b8015610c405780601f10610c1757610100808354040283529160200191610c40565b820191905f5260205f20905b815481529060010190602001808311610c2357829003601f168201915b50505050508152602001600582018054610c5990611f40565b80601f0160208091040260200160405190810160405280929190818152602001828054610c8590611f40565b8015610cd05780601f10610ca757610100808354040283529160200191610cd0565b820191905f5260205f20905b815481529060010190602001808311610cb357829003601f168201915b50505050508152602001600682015481526020016007820154815260200160088201548152505090503373ffffffffffffffffffffffffffffffffffffffff16600183604051610d2091906120e2565b908152602001604051809103902060020160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610da8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9f90612497565b60405180910390fd5b60016003811115610dbc57610dbb611bd0565b5b600183604051610dcc91906120e2565b90815260200160405180910390206002015f9054906101000a900460ff166003811115610dfc57610dfb611bd0565b5b14610e3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e33906124ff565b60405180910390fd5b5f60149054906101000a900460ff1615610e8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8290611fba565b60405180910390fd5b5f835111610ece576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec590612022565b60405180910390fd5b5f600184604051610edf91906120e2565b908152602001604051809103902090506003816002015f6101000a81548160ff02191690836003811115610f1657610f15611bd0565b5b02179055505f8160020160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504281600801819055503373ffffffffffffffffffffffffffffffffffffffff167ff2d02e58fd2bc0370968819f376a61ab21c0ff70786f41b9247053e81dd2f977856003604051610fb0929190612540565b60405180910390a250505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b815f600182604051610ff491906120e2565b9081526020016040518091039020604051806101400160405290815f8201805461101d90611f40565b80601f016020809104026020016040519081016040528092919081815260200182805461104990611f40565b80156110945780601f1061106b57610100808354040283529160200191611094565b820191905f5260205f20905b81548152906001019060200180831161107757829003601f168201915b505050505081526020016001820180546110ad90611f40565b80601f01602080910402602001604051908101604052809291908181526020018280546110d990611f40565b80156111245780601f106110fb57610100808354040283529160200191611124565b820191905f5260205f20905b81548152906001019060200180831161110757829003601f168201915b50505050508152602001600282015f9054906101000a900460ff16600381111561115157611150611bd0565b5b600381111561116357611162611bd0565b5b81526020016002820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201805461122290611f40565b80601f016020809104026020016040519081016040528092919081815260200182805461124e90611f40565b80156112995780601f1061127057610100808354040283529160200191611299565b820191905f5260205f20905b81548152906001019060200180831161127c57829003601f168201915b505050505081526020016005820180546112b290611f40565b80601f01602080910402602001604051908101604052809291908181526020018280546112de90611f40565b80156113295780601f1061130057610100808354040283529160200191611329565b820191905f5260205f20905b81548152906001019060200180831161130c57829003601f168201915b50505050508152602001600682015481526020016007820154815260200160088201548152505090503373ffffffffffffffffffffffffffffffffffffffff1660018360405161137991906120e2565b908152602001604051809103902060020160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611401576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113f890612497565b60405180910390fd5b6001600381111561141557611414611bd0565b5b60018360405161142591906120e2565b90815260200160405180910390206002015f9054906101000a900460ff16600381111561145557611454611bd0565b5b14611495576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161148c906124ff565b60405180910390fd5b5f60149054906101000a900460ff16156114e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114db90611fba565b60405180910390fd5b5f845111611527576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151e90612022565b60405180910390fd5b5f60018560405161153891906120e2565b908152602001604051809103902090506002816002015f6101000a81548160ff0219169083600381111561156f5761156e611bd0565b5b0217905550838160050190816115859190612326565b504281600801819055503373ffffffffffffffffffffffffffffffffffffffff167ff2d02e58fd2bc0370968819f376a61ab21c0ff70786f41b9247053e81dd2f977866002876040516115da93929190612581565b60405180910390a25f8160020160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826006015460405161162f906125ee565b5f6040518083038185875af1925050503d805f8114611669576040519150601f19603f3d011682016040523d82523d5f602084013e61166e565b606091505b50509050806116b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116a99061264c565b60405180910390fd5b505050505050565b5f60149054906101000a900460ff1615611709576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161170090611fba565b60405180910390fd5b5f81511161174c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174390612022565b60405180910390fd5b5f60018260405161175d91906120e2565b908152602001604051809103902090505f60038111156117805761177f611bd0565b5b816002015f9054906101000a900460ff1660038111156117a3576117a2611bd0565b5b14806117e257506003808111156117bd576117bc611bd0565b5b816002015f9054906101000a900460ff1660038111156117e0576117df611bd0565b5b145b611821576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611818906126b4565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16816003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036118b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118a990612742565b60405180910390fd5b6001816002015f6101000a81548160ff021916908360038111156118d9576118d8611bd0565b5b0217905550338160020160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504281600801819055503373ffffffffffffffffffffffffffffffffffffffff167f54e973d1f7ca7aab25c55b07f06ce112eaa3d4c1563914ac7047182ccaaef4f7836040516119709190612760565b60405180910390a25050565b5f604051905090565b5f5ffd5b5f5ffd5b5f8115159050919050565b6119a18161198d565b81146119ab575f5ffd5b50565b5f813590506119bc81611998565b92915050565b5f602082840312156119d7576119d6611985565b5b5f6119e4848285016119ae565b91505092915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611a3b826119f5565b810181811067ffffffffffffffff82111715611a5a57611a59611a05565b5b80604052505050565b5f611a6c61197c565b9050611a788282611a32565b919050565b5f67ffffffffffffffff821115611a9757611a96611a05565b5b611aa0826119f5565b9050602081019050919050565b828183375f83830152505050565b5f611acd611ac884611a7d565b611a63565b905082815260208101848484011115611ae957611ae86119f1565b5b611af4848285611aad565b509392505050565b5f82601f830112611b1057611b0f6119ed565b5b8135611b20848260208601611abb565b91505092915050565b5f60208284031215611b3e57611b3d611985565b5b5f82013567ffffffffffffffff811115611b5b57611b5a611989565b5b611b6784828501611afc565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f611ba282611b70565b611bac8185611b7a565b9350611bbc818560208601611b8a565b611bc5816119f5565b840191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60048110611c0e57611c0d611bd0565b5b50565b5f819050611c1e82611bfd565b919050565b5f611c2d82611c11565b9050919050565b611c3d81611c23565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611c6c82611c43565b9050919050565b611c7c81611c62565b82525050565b5f819050919050565b611c9481611c82565b82525050565b5f610140820190508181035f830152611cb3818d611b98565b90508181036020830152611cc7818c611b98565b9050611cd6604083018b611c34565b611ce3606083018a611c73565b611cf06080830189611c73565b81810360a0830152611d028188611b98565b905081810360c0830152611d168187611b98565b9050611d2560e0830186611c8b565b611d33610100830185611c8b565b611d41610120830184611c8b565b9b9a5050505050505050505050565b611d598161198d565b82525050565b5f602082019050611d725f830184611d50565b92915050565b5f5f5f60608486031215611d8f57611d8e611985565b5b5f84013567ffffffffffffffff811115611dac57611dab611989565b5b611db886828701611afc565b935050602084013567ffffffffffffffff811115611dd957611dd8611989565b5b611de586828701611afc565b925050604084013567ffffffffffffffff811115611e0657611e05611989565b5b611e1286828701611afc565b9150509250925092565b5f602082019050611e2f5f830184611c73565b92915050565b5f5f60408385031215611e4b57611e4a611985565b5b5f83013567ffffffffffffffff811115611e6857611e67611989565b5b611e7485828601611afc565b925050602083013567ffffffffffffffff811115611e9557611e94611989565b5b611ea185828601611afc565b9150509250929050565b7f43616c6c6572206973206e6f74206f776e6572000000000000000000000000005f82015250565b5f611edf601383611b7a565b9150611eea82611eab565b602082019050919050565b5f6020820190508181035f830152611f0c81611ed3565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611f5757607f821691505b602082108103611f6a57611f69611f13565b5b50919050565b7f436f6e74726163742069732070617573656400000000000000000000000000005f82015250565b5f611fa4601283611b7a565b9150611faf82611f70565b602082019050919050565b5f6020820190508181035f830152611fd181611f98565b9050919050565b7f555549442063616e6e6f7420626520656d7074790000000000000000000000005f82015250565b5f61200c601483611b7a565b915061201782611fd8565b602082019050919050565b5f6020820190508181035f83015261203981612000565b9050919050565b7f496e76616c696420616d6f756e740000000000000000000000000000000000005f82015250565b5f612074600e83611b7a565b915061207f82612040565b602082019050919050565b5f6020820190508181035f8301526120a181612068565b9050919050565b5f81905092915050565b5f6120bc82611b70565b6120c681856120a8565b93506120d6818560208601611b8a565b80840191505092915050565b5f6120ed82846120b2565b915081905092915050565b7f5461736b20776974682074686973205555494420616c726561647920657869735f8201527f7473000000000000000000000000000000000000000000000000000000000000602082015250565b5f612152602283611b7a565b915061215d826120f8565b604082019050919050565b5f6020820190508181035f83015261217f81612146565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026121e27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826121a7565b6121ec86836121a7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61222761222261221d84611c82565b612204565b611c82565b9050919050565b5f819050919050565b6122408361220d565b61225461224c8261222e565b8484546121b3565b825550505050565b5f5f905090565b61226b61225c565b612276818484612237565b505050565b5b818110156122995761228e5f82612263565b60018101905061227c565b5050565b601f8211156122de576122af81612186565b6122b884612198565b810160208510156122c7578190505b6122db6122d385612198565b83018261227b565b50505b505050565b5f82821c905092915050565b5f6122fe5f19846008026122e3565b1980831691505092915050565b5f61231683836122ef565b9150826002028217905092915050565b61232f82611b70565b67ffffffffffffffff81111561234857612347611a05565b5b6123528254611f40565b61235d82828561229d565b5f60209050601f83116001811461238e575f841561237c578287015190505b612386858261230b565b8655506123ed565b601f19841661239c86612186565b5f5b828110156123c35784890151825560018201915060208501945060208101905061239e565b868310156123e057848901516123dc601f8916826122ef565b8355505b6001600288020188555050505b505050505050565b5f6080820190508181035f83015261240d8187611b98565b905081810360208301526124218186611b98565b905081810360408301526124358185611b98565b90506124446060830184611c8b565b95945050505050565b7f43616c6c6572206973206e6f74207461736b2061737369676e656500000000005f82015250565b5f612481601b83611b7a565b915061248c8261244d565b602082019050919050565b5f6020820190508181035f8301526124ae81612475565b9050919050565b7f5461736b206e6f7420696e20436c61696d6564207374617465000000000000005f82015250565b5f6124e9601983611b7a565b91506124f4826124b5565b602082019050919050565b5f6020820190508181035f830152612516816124dd565b9050919050565b50565b5f61252b5f83611b7a565b91506125368261251d565b5f82019050919050565b5f6060820190508181035f8301526125588185611b98565b90506125676020830184611c34565b818103604083015261257881612520565b90509392505050565b5f6060820190508181035f8301526125998186611b98565b90506125a86020830185611c34565b81810360408301526125ba8184611b98565b9050949350505050565b5f81905092915050565b5f6125d95f836125c4565b91506125e48261251d565b5f82019050919050565b5f6125f8826125ce565b9150819050919050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f612636600f83611b7a565b915061264182612602565b602082019050919050565b5f6020820190508181035f8301526126638161262a565b9050919050565b7f4e6f742070656e64696e67206f72206661696c656400000000000000000000005f82015250565b5f61269e601583611b7a565b91506126a98261266a565b602082019050919050565b5f6020820190508181035f8301526126cb81612692565b9050919050565b7f5375626d69747465722063616e6e6f7420636c61696d207468656972206f776e5f8201527f207461736b000000000000000000000000000000000000000000000000000000602082015250565b5f61272c602583611b7a565b9150612737826126d2565b604082019050919050565b5f6020820190508181035f83015261275981612720565b9050919050565b5f6020820190508181035f8301526127788184611b98565b90509291505056fea2646970667358221220f226ba86097c0395ee09b0d2c18603d38419852c91bc53859f513f94bbcfc7a764736f6c634300081e0033",
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
