// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dutchAuction

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

// DutchAuctionMetaData contains all meta data concerning the DutchAuction contract.
var DutchAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_discountRate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"AuctionEnded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discountRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051610e8f380380610e8f83398181016040528101906100329190610205565b8473ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050835f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260a081815250508160c081815250508060e0818152505042610100818152505062093a80426100d591906102a9565b60018190555062093a8060e0516100ec91906102dc565b60c0511161012f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101269061039d565b60405180910390fd5b50505050506103bb565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101668261013d565b9050919050565b6101768161015c565b8114610180575f5ffd5b50565b5f815190506101918161016d565b92915050565b5f6101a18261013d565b9050919050565b6101b181610197565b81146101bb575f5ffd5b50565b5f815190506101cc816101a8565b92915050565b5f819050919050565b6101e4816101d2565b81146101ee575f5ffd5b50565b5f815190506101ff816101db565b92915050565b5f5f5f5f5f60a0868803121561021e5761021d610139565b5b5f61022b88828901610183565b955050602061023c888289016101be565b945050604061024d888289016101f1565b935050606061025e888289016101f1565b925050608061026f888289016101f1565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6102b3826101d2565b91506102be836101d2565b92508282019050808211156102d6576102d561027c565b5b92915050565b5f6102e6826101d2565b91506102f1836101d2565b92508282026102ff816101d2565b915082820484148315176103165761031561027c565b5b5092915050565b5f82825260208201905092915050565b7f5374617274696e672070726963652073686f756c6420626520677265617465725f8201527f207468616e203000000000000000000000000000000000000000000000000000602082015250565b5f61038760278361031d565b91506103928261032d565b604082019050919050565b5f6020820190508181035f8301526103b48161037b565b9050919050565b60805160a05160c05160e05161010051610a7461041b5f395f818161031b01526103bd01525f818161039c015261069c01525f81816103f3015261067801525f81816102c6015261052c01525f81816102f701526104cd0152610a745ff3fe6080604052600436106100a6575f3560e01c806378e979251161006357806378e97925146101a657806398d5fdca146101d0578063a6f2ae3a146101fa578063d6fbf20214610204578063e6c0e6d51461022e578063e852e74114610258576100a6565b806308551a53146100aa57806312065fe0146100d457806317d70f7c146100fe5780631be05289146101285780633197cbb6146101525780635bf8633a1461017c575b5f5ffd5b3480156100b5575f5ffd5b506100be610282565b6040516100cb919061070f565b60405180910390f35b3480156100df575f5ffd5b506100e86102a6565b6040516100f59190610740565b60405180910390f35b348015610109575f5ffd5b506101126102c4565b60405161011f9190610740565b60405180910390f35b348015610133575f5ffd5b5061013c6102e8565b6040516101499190610740565b60405180910390f35b34801561015d575f5ffd5b506101666102ef565b6040516101739190610740565b60405180910390f35b348015610187575f5ffd5b506101906102f5565b60405161019d9190610779565b60405180910390f35b3480156101b1575f5ffd5b506101ba610319565b6040516101c79190610740565b60405180910390f35b3480156101db575f5ffd5b506101e461033d565b6040516101f19190610740565b60405180910390f35b610202610421565b005b34801561020f575f5ffd5b50610218610676565b6040516102259190610740565b60405180910390f35b348015610239575f5ffd5b5061024261069a565b60405161024f9190610740565b60405180910390f35b348015610263575f5ffd5b5061026c6106be565b60405161027991906107ac565b60405180910390f35b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f3373ffffffffffffffffffffffffffffffffffffffff1631905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b62093a8081565b60015481565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f6001544210158061035a57505f60149054906101000a900460ff165b1561039a576040517f4d9638fc0000000000000000000000000000000000000000000000000000000081526004016103919061081f565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000426103e7919061086a565b6103f1919061089d565b7f000000000000000000000000000000000000000000000000000000000000000061041c919061086a565b905090565b6001544210158061043d57505f60149054906101000a900460ff165b1561047d576040517f4d9638fc0000000000000000000000000000000000000000000000000000000081526004016104749061081f565b60405180910390fd5b5f61048661033d565b9050803410156104cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c290610928565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd825f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16337f00000000000000000000000000000000000000000000000000000000000000006040518563ffffffff1660e01b8152600401610569939291906109a1565b5f604051808303818588803b158015610580575f5ffd5b505af1158015610592573d5f5f3e3d5ffd5b50505050505f81346105a4919061086a565b90504260018190555060015f60146101000a81548160ff0219169083151502179055505f811115610672575f3373ffffffffffffffffffffffffffffffffffffffff16826040516105f490610a03565b5f6040518083038185875af1925050503d805f811461062e576040519150601f19603f3d011682016040523d82523d5f602084013e610633565b606091505b505090507fbe74181457391058f21eeb20d41d41f37464c7b5b40f03b4e402e4ea16a1f48e8183604051610668929190610a17565b60405180910390a1505b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f60149054906101000a900460ff1681565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106f9826106d0565b9050919050565b610709816106ef565b82525050565b5f6020820190506107225f830184610700565b92915050565b5f819050919050565b61073a81610728565b82525050565b5f6020820190506107535f830184610731565b92915050565b5f610763826106d0565b9050919050565b61077381610759565b82525050565b5f60208201905061078c5f83018461076a565b92915050565b5f8115159050919050565b6107a681610792565b82525050565b5f6020820190506107bf5f83018461079d565b92915050565b5f82825260208201905092915050565b7f41756374696f6e20656e646564000000000000000000000000000000000000005f82015250565b5f610809600d836107c5565b9150610814826107d5565b602082019050919050565b5f6020820190508181035f830152610836816107fd565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61087482610728565b915061087f83610728565b92508282039050818111156108975761089661083d565b5b92915050565b5f6108a782610728565b91506108b283610728565b92508282026108c081610728565b915082820484148315176108d7576108d661083d565b5b5092915050565b7f4e6f7420656e6f756768206574686572000000000000000000000000000000005f82015250565b5f6109126010836107c5565b915061091d826108de565b602082019050919050565b5f6020820190508181035f83015261093f81610906565b9050919050565b5f819050919050565b5f61096961096461095f846106d0565b610946565b6106d0565b9050919050565b5f61097a8261094f565b9050919050565b5f61098b82610970565b9050919050565b61099b81610981565b82525050565b5f6060820190506109b45f830186610992565b6109c1602083018561076a565b6109ce6040830184610731565b949350505050565b5f81905092915050565b50565b5f6109ee5f836109d6565b91506109f9826109e0565b5f82019050919050565b5f610a0d826109e3565b9150819050919050565b5f604082019050610a2a5f83018561079d565b610a376020830184610731565b939250505056fea2646970667358221220af00c5966cd77d3f37166227ac431ba0d40b2af3ae6867d84132b5072264f6ce64736f6c634300081b0033",
}

// DutchAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use DutchAuctionMetaData.ABI instead.
var DutchAuctionABI = DutchAuctionMetaData.ABI

// DutchAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DutchAuctionMetaData.Bin instead.
var DutchAuctionBin = DutchAuctionMetaData.Bin

// DeployDutchAuction deploys a new Ethereum contract, binding an instance of DutchAuction to it.
func DeployDutchAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddress common.Address, _seller common.Address, _tokenId *big.Int, _startingPrice *big.Int, _discountRate *big.Int) (common.Address, *types.Transaction, *DutchAuction, error) {
	parsed, err := DutchAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DutchAuctionBin), backend, _nftAddress, _seller, _tokenId, _startingPrice, _discountRate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DutchAuction{DutchAuctionCaller: DutchAuctionCaller{contract: contract}, DutchAuctionTransactor: DutchAuctionTransactor{contract: contract}, DutchAuctionFilterer: DutchAuctionFilterer{contract: contract}}, nil
}

// DutchAuction is an auto generated Go binding around an Ethereum contract.
type DutchAuction struct {
	DutchAuctionCaller     // Read-only binding to the contract
	DutchAuctionTransactor // Write-only binding to the contract
	DutchAuctionFilterer   // Log filterer for contract events
}

// DutchAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DutchAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DutchAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DutchAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DutchAuctionSession struct {
	Contract     *DutchAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DutchAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DutchAuctionCallerSession struct {
	Contract *DutchAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DutchAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DutchAuctionTransactorSession struct {
	Contract     *DutchAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DutchAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DutchAuctionRaw struct {
	Contract *DutchAuction // Generic contract binding to access the raw methods on
}

// DutchAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DutchAuctionCallerRaw struct {
	Contract *DutchAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// DutchAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DutchAuctionTransactorRaw struct {
	Contract *DutchAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDutchAuction creates a new instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuction(address common.Address, backend bind.ContractBackend) (*DutchAuction, error) {
	contract, err := bindDutchAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DutchAuction{DutchAuctionCaller: DutchAuctionCaller{contract: contract}, DutchAuctionTransactor: DutchAuctionTransactor{contract: contract}, DutchAuctionFilterer: DutchAuctionFilterer{contract: contract}}, nil
}

// NewDutchAuctionCaller creates a new read-only instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionCaller(address common.Address, caller bind.ContractCaller) (*DutchAuctionCaller, error) {
	contract, err := bindDutchAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionCaller{contract: contract}, nil
}

// NewDutchAuctionTransactor creates a new write-only instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*DutchAuctionTransactor, error) {
	contract, err := bindDutchAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionTransactor{contract: contract}, nil
}

// NewDutchAuctionFilterer creates a new log filterer instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*DutchAuctionFilterer, error) {
	contract, err := bindDutchAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionFilterer{contract: contract}, nil
}

// bindDutchAuction binds a generic wrapper to an already deployed contract.
func bindDutchAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DutchAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DutchAuction *DutchAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DutchAuction.Contract.DutchAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DutchAuction *DutchAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchAuction.Contract.DutchAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DutchAuction *DutchAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DutchAuction.Contract.DutchAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DutchAuction *DutchAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DutchAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DutchAuction *DutchAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DutchAuction *DutchAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DutchAuction.Contract.contract.Transact(opts, method, params...)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) DURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_DutchAuction *DutchAuctionSession) DURATION() (*big.Int, error) {
	return _DutchAuction.Contract.DURATION(&_DutchAuction.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) DURATION() (*big.Int, error) {
	return _DutchAuction.Contract.DURATION(&_DutchAuction.CallOpts)
}

// DiscountRate is a free data retrieval call binding the contract method 0xe6c0e6d5.
//
// Solidity: function discountRate() view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) DiscountRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "discountRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiscountRate is a free data retrieval call binding the contract method 0xe6c0e6d5.
//
// Solidity: function discountRate() view returns(uint256)
func (_DutchAuction *DutchAuctionSession) DiscountRate() (*big.Int, error) {
	return _DutchAuction.Contract.DiscountRate(&_DutchAuction.CallOpts)
}

// DiscountRate is a free data retrieval call binding the contract method 0xe6c0e6d5.
//
// Solidity: function discountRate() view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) DiscountRate() (*big.Int, error) {
	return _DutchAuction.Contract.DiscountRate(&_DutchAuction.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_DutchAuction *DutchAuctionSession) EndTime() (*big.Int, error) {
	return _DutchAuction.Contract.EndTime(&_DutchAuction.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) EndTime() (*big.Int, error) {
	return _DutchAuction.Contract.EndTime(&_DutchAuction.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_DutchAuction *DutchAuctionSession) GetBalance() (*big.Int, error) {
	return _DutchAuction.Contract.GetBalance(&_DutchAuction.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) GetBalance() (*big.Int, error) {
	return _DutchAuction.Contract.GetBalance(&_DutchAuction.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) GetPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "getPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_DutchAuction *DutchAuctionSession) GetPrice() (*big.Int, error) {
	return _DutchAuction.Contract.GetPrice(&_DutchAuction.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) GetPrice() (*big.Int, error) {
	return _DutchAuction.Contract.GetPrice(&_DutchAuction.CallOpts)
}

// IsSold is a free data retrieval call binding the contract method 0xe852e741.
//
// Solidity: function isSold() view returns(bool)
func (_DutchAuction *DutchAuctionCaller) IsSold(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "isSold")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSold is a free data retrieval call binding the contract method 0xe852e741.
//
// Solidity: function isSold() view returns(bool)
func (_DutchAuction *DutchAuctionSession) IsSold() (bool, error) {
	return _DutchAuction.Contract.IsSold(&_DutchAuction.CallOpts)
}

// IsSold is a free data retrieval call binding the contract method 0xe852e741.
//
// Solidity: function isSold() view returns(bool)
func (_DutchAuction *DutchAuctionCallerSession) IsSold() (bool, error) {
	return _DutchAuction.Contract.IsSold(&_DutchAuction.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_DutchAuction *DutchAuctionCaller) NftAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "nftAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_DutchAuction *DutchAuctionSession) NftAddress() (common.Address, error) {
	return _DutchAuction.Contract.NftAddress(&_DutchAuction.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_DutchAuction *DutchAuctionCallerSession) NftAddress() (common.Address, error) {
	return _DutchAuction.Contract.NftAddress(&_DutchAuction.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_DutchAuction *DutchAuctionCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_DutchAuction *DutchAuctionSession) Seller() (common.Address, error) {
	return _DutchAuction.Contract.Seller(&_DutchAuction.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_DutchAuction *DutchAuctionCallerSession) Seller() (common.Address, error) {
	return _DutchAuction.Contract.Seller(&_DutchAuction.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DutchAuction *DutchAuctionSession) StartTime() (*big.Int, error) {
	return _DutchAuction.Contract.StartTime(&_DutchAuction.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) StartTime() (*big.Int, error) {
	return _DutchAuction.Contract.StartTime(&_DutchAuction.CallOpts)
}

// StartingPrice is a free data retrieval call binding the contract method 0xd6fbf202.
//
// Solidity: function startingPrice() view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) StartingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "startingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingPrice is a free data retrieval call binding the contract method 0xd6fbf202.
//
// Solidity: function startingPrice() view returns(uint256)
func (_DutchAuction *DutchAuctionSession) StartingPrice() (*big.Int, error) {
	return _DutchAuction.Contract.StartingPrice(&_DutchAuction.CallOpts)
}

// StartingPrice is a free data retrieval call binding the contract method 0xd6fbf202.
//
// Solidity: function startingPrice() view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) StartingPrice() (*big.Int, error) {
	return _DutchAuction.Contract.StartingPrice(&_DutchAuction.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_DutchAuction *DutchAuctionSession) TokenId() (*big.Int, error) {
	return _DutchAuction.Contract.TokenId(&_DutchAuction.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) TokenId() (*big.Int, error) {
	return _DutchAuction.Contract.TokenId(&_DutchAuction.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_DutchAuction *DutchAuctionTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_DutchAuction *DutchAuctionSession) Buy() (*types.Transaction, error) {
	return _DutchAuction.Contract.Buy(&_DutchAuction.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_DutchAuction *DutchAuctionTransactorSession) Buy() (*types.Transaction, error) {
	return _DutchAuction.Contract.Buy(&_DutchAuction.TransactOpts)
}

// DutchAuctionRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the DutchAuction contract.
type DutchAuctionRefundIterator struct {
	Event *DutchAuctionRefund // Event containing the contract specifics and raw log

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
func (it *DutchAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchAuctionRefund)
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
		it.Event = new(DutchAuctionRefund)
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
func (it *DutchAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchAuctionRefund represents a Refund event raised by the DutchAuction contract.
type DutchAuctionRefund struct {
	Success bool
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xbe74181457391058f21eeb20d41d41f37464c7b5b40f03b4e402e4ea16a1f48e.
//
// Solidity: event Refund(bool success, uint256 amount)
func (_DutchAuction *DutchAuctionFilterer) FilterRefund(opts *bind.FilterOpts) (*DutchAuctionRefundIterator, error) {

	logs, sub, err := _DutchAuction.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &DutchAuctionRefundIterator{contract: _DutchAuction.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xbe74181457391058f21eeb20d41d41f37464c7b5b40f03b4e402e4ea16a1f48e.
//
// Solidity: event Refund(bool success, uint256 amount)
func (_DutchAuction *DutchAuctionFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *DutchAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _DutchAuction.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchAuctionRefund)
				if err := _DutchAuction.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0xbe74181457391058f21eeb20d41d41f37464c7b5b40f03b4e402e4ea16a1f48e.
//
// Solidity: event Refund(bool success, uint256 amount)
func (_DutchAuction *DutchAuctionFilterer) ParseRefund(log types.Log) (*DutchAuctionRefund, error) {
	event := new(DutchAuctionRefund)
	if err := _DutchAuction.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
