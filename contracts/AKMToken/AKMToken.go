// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AKMToken

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

// AKMTokenMetaData contains all meta data concerning the AKMToken contract.
var AKMTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611c4d380380611c4d83398181016040528101906100319190610543565b806040518060400160405280600881526020017f414b4d546f6b656e0000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f414b4d000000000000000000000000000000000000000000000000000000000081525081600390816100ad91906107ab565b5080600490816100bd91906107ab565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610130575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016101279190610889565b60405180910390fd5b61013f8161017c60201b60201c565b506101763361015261023f60201b60201c565b600a61015e9190610a0a565b6103e861016b9190610a54565b61024760201b60201c565b50610b25565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f6012905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036102b7575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016102ae9190610889565b60405180910390fd5b6102c85f83836102cc60201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361031c578060025f8282546103109190610a95565b925050819055506103ea565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050818110156103a5578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161039c93929190610ad7565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610431578060025f828254039250508190555061047b565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516104d89190610b0c565b60405180910390a3505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610512826104e9565b9050919050565b61052281610508565b811461052c575f5ffd5b50565b5f8151905061053d81610519565b92915050565b5f60208284031215610558576105576104e5565b5b5f6105658482850161052f565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806105e957607f821691505b6020821081036105fc576105fb6105a5565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261065e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610623565b6106688683610623565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6106ac6106a76106a284610680565b610689565b610680565b9050919050565b5f819050919050565b6106c583610692565b6106d96106d1826106b3565b84845461062f565b825550505050565b5f5f905090565b6106f06106e1565b6106fb8184846106bc565b505050565b5b8181101561071e576107135f826106e8565b600181019050610701565b5050565b601f8211156107635761073481610602565b61073d84610614565b8101602085101561074c578190505b61076061075885610614565b830182610700565b50505b505050565b5f82821c905092915050565b5f6107835f1984600802610768565b1980831691505092915050565b5f61079b8383610774565b9150826002028217905092915050565b6107b48261056e565b67ffffffffffffffff8111156107cd576107cc610578565b5b6107d782546105d2565b6107e2828285610722565b5f60209050601f831160018114610813575f8415610801578287015190505b61080b8582610790565b865550610872565b601f19841661082186610602565b5f5b8281101561084857848901518255600182019150602085019450602081019050610823565b868310156108655784890151610861601f891682610774565b8355505b6001600288020188555050505b505050505050565b61088381610508565b82525050565b5f60208201905061089c5f83018461087a565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b600185111561092457808604811115610900576108ff6108a2565b5b600185161561090f5780820291505b808102905061091d856108cf565b94506108e4565b94509492505050565b5f8261093c57600190506109f7565b81610949575f90506109f7565b816001811461095f576002811461096957610998565b60019150506109f7565b60ff84111561097b5761097a6108a2565b5b8360020a915084821115610992576109916108a2565b5b506109f7565b5060208310610133831016604e8410600b84101617156109cd5782820a9050838111156109c8576109c76108a2565b5b6109f7565b6109da84848460016108db565b925090508184048111156109f1576109f06108a2565b5b81810290505b9392505050565b5f60ff82169050919050565b5f610a1482610680565b9150610a1f836109fe565b9250610a4c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461092d565b905092915050565b5f610a5e82610680565b9150610a6983610680565b9250828202610a7781610680565b91508282048414831517610a8e57610a8d6108a2565b5b5092915050565b5f610a9f82610680565b9150610aaa83610680565b9250828201905080821115610ac257610ac16108a2565b5b92915050565b610ad181610680565b82525050565b5f606082019050610aea5f83018661087a565b610af76020830185610ac8565b610b046040830184610ac8565b949350505050565b5f602082019050610b1f5f830184610ac8565b92915050565b61111b80610b325f395ff3fe608060405234801561000f575f5ffd5b50600436106100cd575f3560e01c806370a082311161008a57806395d89b411161006457806395d89b41146101ff578063a9059cbb1461021d578063dd62ed3e1461024d578063f2fde38b1461027d576100cd565b806370a08231146101a7578063715018a6146101d75780638da5cb5b146101e1576100cd565b806306fdde03146100d1578063095ea7b3146100ef57806318160ddd1461011f57806323b872dd1461013d578063313ce5671461016d57806340c10f191461018b575b5f5ffd5b6100d9610299565b6040516100e69190610d94565b60405180910390f35b61010960048036038101906101049190610e45565b610329565b6040516101169190610e9d565b60405180910390f35b61012761034b565b6040516101349190610ec5565b60405180910390f35b61015760048036038101906101529190610ede565b610354565b6040516101649190610e9d565b60405180910390f35b610175610382565b6040516101829190610f49565b60405180910390f35b6101a560048036038101906101a09190610e45565b61038a565b005b6101c160048036038101906101bc9190610f62565b6103a0565b6040516101ce9190610ec5565b60405180910390f35b6101df6103e5565b005b6101e96103f8565b6040516101f69190610f9c565b60405180910390f35b610207610420565b6040516102149190610d94565b60405180910390f35b61023760048036038101906102329190610e45565b6104b0565b6040516102449190610e9d565b60405180910390f35b61026760048036038101906102629190610fb5565b6104d2565b6040516102749190610ec5565b60405180910390f35b61029760048036038101906102929190610f62565b610554565b005b6060600380546102a890611020565b80601f01602080910402602001604051908101604052809291908181526020018280546102d490611020565b801561031f5780601f106102f65761010080835404028352916020019161031f565b820191905f5260205f20905b81548152906001019060200180831161030257829003601f168201915b5050505050905090565b5f5f6103336105d8565b90506103408185856105df565b600191505092915050565b5f600254905090565b5f5f61035e6105d8565b905061036b8582856105f1565b610376858585610683565b60019150509392505050565b5f6012905090565b610392610773565b61039c82826107fa565b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6103ed610773565b6103f65f610879565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461042f90611020565b80601f016020809104026020016040519081016040528092919081815260200182805461045b90611020565b80156104a65780601f1061047d576101008083540402835291602001916104a6565b820191905f5260205f20905b81548152906001019060200180831161048957829003601f168201915b5050505050905090565b5f5f6104ba6105d8565b90506104c7818585610683565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b61055c610773565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105cc575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016105c39190610f9c565b60405180910390fd5b6105d581610879565b50565b5f33905090565b6105ec838383600161093c565b505050565b5f6105fc84846104d2565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461067d578181101561066e578281836040517ffb8f41b200000000000000000000000000000000000000000000000000000000815260040161066593929190611050565b60405180910390fd5b61067c84848484035f61093c565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036106f3575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016106ea9190610f9c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610763575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161075a9190610f9c565b60405180910390fd5b61076e838383610b0b565b505050565b61077b6105d8565b73ffffffffffffffffffffffffffffffffffffffff166107996103f8565b73ffffffffffffffffffffffffffffffffffffffff16146107f8576107bc6105d8565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016107ef9190610f9c565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361086a575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016108619190610f9c565b60405180910390fd5b6108755f8383610b0b565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036109ac575f6040517fe602df050000000000000000000000000000000000000000000000000000000081526004016109a39190610f9c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a1c575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610a139190610f9c565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610b05578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610afc9190610ec5565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b5b578060025f828254610b4f91906110b2565b92505081905550610c29565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610be4578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610bdb93929190611050565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c70578060025f8282540392505081905550610cba565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d179190610ec5565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610d6682610d24565b610d708185610d2e565b9350610d80818560208601610d3e565b610d8981610d4c565b840191505092915050565b5f6020820190508181035f830152610dac8184610d5c565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610de182610db8565b9050919050565b610df181610dd7565b8114610dfb575f5ffd5b50565b5f81359050610e0c81610de8565b92915050565b5f819050919050565b610e2481610e12565b8114610e2e575f5ffd5b50565b5f81359050610e3f81610e1b565b92915050565b5f5f60408385031215610e5b57610e5a610db4565b5b5f610e6885828601610dfe565b9250506020610e7985828601610e31565b9150509250929050565b5f8115159050919050565b610e9781610e83565b82525050565b5f602082019050610eb05f830184610e8e565b92915050565b610ebf81610e12565b82525050565b5f602082019050610ed85f830184610eb6565b92915050565b5f5f5f60608486031215610ef557610ef4610db4565b5b5f610f0286828701610dfe565b9350506020610f1386828701610dfe565b9250506040610f2486828701610e31565b9150509250925092565b5f60ff82169050919050565b610f4381610f2e565b82525050565b5f602082019050610f5c5f830184610f3a565b92915050565b5f60208284031215610f7757610f76610db4565b5b5f610f8484828501610dfe565b91505092915050565b610f9681610dd7565b82525050565b5f602082019050610faf5f830184610f8d565b92915050565b5f5f60408385031215610fcb57610fca610db4565b5b5f610fd885828601610dfe565b9250506020610fe985828601610dfe565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061103757607f821691505b60208210810361104a57611049610ff3565b5b50919050565b5f6060820190506110635f830186610f8d565b6110706020830185610eb6565b61107d6040830184610eb6565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6110bc82610e12565b91506110c783610e12565b92508282019050808211156110df576110de611085565b5b9291505056fea26469706673582212200b80dfda725dcaa4f40432fc74089f837126d5e8950a455ac6211a96a91c741664736f6c634300081b0033",
}

// AKMTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use AKMTokenMetaData.ABI instead.
var AKMTokenABI = AKMTokenMetaData.ABI

// AKMTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AKMTokenMetaData.Bin instead.
var AKMTokenBin = AKMTokenMetaData.Bin

// DeployAKMToken deploys a new Ethereum contract, binding an instance of AKMToken to it.
func DeployAKMToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address) (common.Address, *types.Transaction, *AKMToken, error) {
	parsed, err := AKMTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AKMTokenBin), backend, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AKMToken{AKMTokenCaller: AKMTokenCaller{contract: contract}, AKMTokenTransactor: AKMTokenTransactor{contract: contract}, AKMTokenFilterer: AKMTokenFilterer{contract: contract}}, nil
}

// AKMToken is an auto generated Go binding around an Ethereum contract.
type AKMToken struct {
	AKMTokenCaller     // Read-only binding to the contract
	AKMTokenTransactor // Write-only binding to the contract
	AKMTokenFilterer   // Log filterer for contract events
}

// AKMTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AKMTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AKMTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AKMTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AKMTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AKMTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AKMTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AKMTokenSession struct {
	Contract     *AKMToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AKMTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AKMTokenCallerSession struct {
	Contract *AKMTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AKMTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AKMTokenTransactorSession struct {
	Contract     *AKMTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AKMTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AKMTokenRaw struct {
	Contract *AKMToken // Generic contract binding to access the raw methods on
}

// AKMTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AKMTokenCallerRaw struct {
	Contract *AKMTokenCaller // Generic read-only contract binding to access the raw methods on
}

// AKMTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AKMTokenTransactorRaw struct {
	Contract *AKMTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAKMToken creates a new instance of AKMToken, bound to a specific deployed contract.
func NewAKMToken(address common.Address, backend bind.ContractBackend) (*AKMToken, error) {
	contract, err := bindAKMToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AKMToken{AKMTokenCaller: AKMTokenCaller{contract: contract}, AKMTokenTransactor: AKMTokenTransactor{contract: contract}, AKMTokenFilterer: AKMTokenFilterer{contract: contract}}, nil
}

// NewAKMTokenCaller creates a new read-only instance of AKMToken, bound to a specific deployed contract.
func NewAKMTokenCaller(address common.Address, caller bind.ContractCaller) (*AKMTokenCaller, error) {
	contract, err := bindAKMToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AKMTokenCaller{contract: contract}, nil
}

// NewAKMTokenTransactor creates a new write-only instance of AKMToken, bound to a specific deployed contract.
func NewAKMTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*AKMTokenTransactor, error) {
	contract, err := bindAKMToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AKMTokenTransactor{contract: contract}, nil
}

// NewAKMTokenFilterer creates a new log filterer instance of AKMToken, bound to a specific deployed contract.
func NewAKMTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*AKMTokenFilterer, error) {
	contract, err := bindAKMToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AKMTokenFilterer{contract: contract}, nil
}

// bindAKMToken binds a generic wrapper to an already deployed contract.
func bindAKMToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AKMTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AKMToken *AKMTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AKMToken.Contract.AKMTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AKMToken *AKMTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AKMToken.Contract.AKMTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AKMToken *AKMTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AKMToken.Contract.AKMTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AKMToken *AKMTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AKMToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AKMToken *AKMTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AKMToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AKMToken *AKMTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AKMToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AKMToken *AKMTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AKMToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AKMToken *AKMTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AKMToken.Contract.Allowance(&_AKMToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AKMToken *AKMTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AKMToken.Contract.Allowance(&_AKMToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AKMToken *AKMTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AKMToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AKMToken *AKMTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AKMToken.Contract.BalanceOf(&_AKMToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AKMToken *AKMTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AKMToken.Contract.BalanceOf(&_AKMToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AKMToken *AKMTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AKMToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AKMToken *AKMTokenSession) Decimals() (uint8, error) {
	return _AKMToken.Contract.Decimals(&_AKMToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AKMToken *AKMTokenCallerSession) Decimals() (uint8, error) {
	return _AKMToken.Contract.Decimals(&_AKMToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AKMToken *AKMTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AKMToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AKMToken *AKMTokenSession) Name() (string, error) {
	return _AKMToken.Contract.Name(&_AKMToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AKMToken *AKMTokenCallerSession) Name() (string, error) {
	return _AKMToken.Contract.Name(&_AKMToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AKMToken *AKMTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AKMToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AKMToken *AKMTokenSession) Owner() (common.Address, error) {
	return _AKMToken.Contract.Owner(&_AKMToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AKMToken *AKMTokenCallerSession) Owner() (common.Address, error) {
	return _AKMToken.Contract.Owner(&_AKMToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AKMToken *AKMTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AKMToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AKMToken *AKMTokenSession) Symbol() (string, error) {
	return _AKMToken.Contract.Symbol(&_AKMToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AKMToken *AKMTokenCallerSession) Symbol() (string, error) {
	return _AKMToken.Contract.Symbol(&_AKMToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AKMToken *AKMTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AKMToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AKMToken *AKMTokenSession) TotalSupply() (*big.Int, error) {
	return _AKMToken.Contract.TotalSupply(&_AKMToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AKMToken *AKMTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _AKMToken.Contract.TotalSupply(&_AKMToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AKMToken *AKMTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AKMToken *AKMTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.Contract.Approve(&_AKMToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AKMToken *AKMTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.Contract.Approve(&_AKMToken.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_AKMToken *AKMTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AKMToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_AKMToken *AKMTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AKMToken.Contract.Mint(&_AKMToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_AKMToken *AKMTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AKMToken.Contract.Mint(&_AKMToken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AKMToken *AKMTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AKMToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AKMToken *AKMTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _AKMToken.Contract.RenounceOwnership(&_AKMToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AKMToken *AKMTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AKMToken.Contract.RenounceOwnership(&_AKMToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AKMToken *AKMTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AKMToken *AKMTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.Contract.Transfer(&_AKMToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AKMToken *AKMTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.Contract.Transfer(&_AKMToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AKMToken *AKMTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AKMToken *AKMTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.Contract.TransferFrom(&_AKMToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AKMToken *AKMTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AKMToken.Contract.TransferFrom(&_AKMToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AKMToken *AKMTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AKMToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AKMToken *AKMTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AKMToken.Contract.TransferOwnership(&_AKMToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AKMToken *AKMTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AKMToken.Contract.TransferOwnership(&_AKMToken.TransactOpts, newOwner)
}

// AKMTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AKMToken contract.
type AKMTokenApprovalIterator struct {
	Event *AKMTokenApproval // Event containing the contract specifics and raw log

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
func (it *AKMTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AKMTokenApproval)
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
		it.Event = new(AKMTokenApproval)
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
func (it *AKMTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AKMTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AKMTokenApproval represents a Approval event raised by the AKMToken contract.
type AKMTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AKMToken *AKMTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AKMTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AKMToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AKMTokenApprovalIterator{contract: _AKMToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AKMToken *AKMTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AKMTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AKMToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AKMTokenApproval)
				if err := _AKMToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AKMToken *AKMTokenFilterer) ParseApproval(log types.Log) (*AKMTokenApproval, error) {
	event := new(AKMTokenApproval)
	if err := _AKMToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AKMTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AKMToken contract.
type AKMTokenOwnershipTransferredIterator struct {
	Event *AKMTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AKMTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AKMTokenOwnershipTransferred)
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
		it.Event = new(AKMTokenOwnershipTransferred)
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
func (it *AKMTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AKMTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AKMTokenOwnershipTransferred represents a OwnershipTransferred event raised by the AKMToken contract.
type AKMTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AKMToken *AKMTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AKMTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AKMToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AKMTokenOwnershipTransferredIterator{contract: _AKMToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AKMToken *AKMTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AKMTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AKMToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AKMTokenOwnershipTransferred)
				if err := _AKMToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AKMToken *AKMTokenFilterer) ParseOwnershipTransferred(log types.Log) (*AKMTokenOwnershipTransferred, error) {
	event := new(AKMTokenOwnershipTransferred)
	if err := _AKMToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AKMTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AKMToken contract.
type AKMTokenTransferIterator struct {
	Event *AKMTokenTransfer // Event containing the contract specifics and raw log

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
func (it *AKMTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AKMTokenTransfer)
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
		it.Event = new(AKMTokenTransfer)
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
func (it *AKMTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AKMTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AKMTokenTransfer represents a Transfer event raised by the AKMToken contract.
type AKMTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AKMToken *AKMTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AKMTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AKMToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AKMTokenTransferIterator{contract: _AKMToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AKMToken *AKMTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AKMTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AKMToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AKMTokenTransfer)
				if err := _AKMToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AKMToken *AKMTokenFilterer) ParseTransfer(log types.Log) (*AKMTokenTransfer, error) {
	event := new(AKMTokenTransfer)
	if err := _AKMToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
