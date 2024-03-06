// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifiermtpwrapper

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

// VerifiermtpwrapperMetaData contains all meta data concerning the Verifiermtpwrapper contract.
var VerifiermtpwrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"input\",\"type\":\"uint256[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[11]\",\"name\":\"_pubSignals\",\"type\":\"uint256[11]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610f84806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632612907c1461003b578063b9c6ea871461006b575b600080fd5b61005560048036038101906100509190610ab3565b61009b565b6040516100629190610b58565b60405180910390f35b61008560048036038101906100809190610b95565b6101cd565b6040516100929190610b58565b60405180910390f35b60006100a56109dd565b600b84849050146100eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100e290610c5b565b60405180910390fd5b60005b600b8110156101405784848281811061010a57610109610c7b565b5b905060200201358282600b811061012457610123610c7b565b5b602002018181525050808061013890610ce3565b9150506100ee565b503073ffffffffffffffffffffffffffffffffffffffff1663b9c6ea87888888856040518563ffffffff1660e01b81526004016101809493929190610eae565b602060405180830381865afa15801561019d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101c19190610f21565b91505095945050505050565b600061091b565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478110610205576000805260206000f35b50565b600060405183815284602082015285604082015260408160608360076107d05a03fa91508161023b576000805260206000f35b825160408201526020830151606082015260408360808360066107d05a03fa91508161026b576000805260206000f35b505050505050565b600060808601600087017f02e76356c6f5d7f49e6d4f6ec3a93cdc50cc3c271b41a11b21fc6c219031429e81527f21d57a61a975d9eb14e29a76de5ad78404f7db5867e0c49d51e939f1cbdd298d602082015261031660008801357f23c72bedb38cd965afd5a35233c42c2cd8fa23d2291ced5153eb0945cc42c2707f2b0b7e4eb338e79dc49653f0a32a4acb42d6d65d67d6a8796c8d292012f9914284610208565b61036660208801357f29cbfecbf222d323af8e1c9654292ece38b45352b18200308e6259773dd5457a7f0bc7aa6c1aef693444e35935cea1fa02ceef7c4eb684e0a71c380bd1cbea394684610208565b6103b660408801357f286a77cb393820c524cdd54518ea43d3f27d595bf6a910694a1170fbe35685427f182bf69bce645da388b8fb33616b6c43c48a50c499d07be536520191ef91dc0884610208565b61040660608801357f105535f6a709c5189e305e1a0c70efb5fae522a60afe6fed5c805dbbc8779e297f198e0c3ad41c05fced9aa2d6e27d9a2c9818b5b88f19c2e698a8648129b6af6a84610208565b61045560808801357f1243fb77067200f806f576091ea92a5281f060234e6fc57df1901b4e3450bb697e9237f01920dbe104631d35579740ad6195073ec934184477efaef11e8ee3aa84610208565b6104a560a08801357f2b72e8feb8d4ff6aa248af1456fc4899864b0489b8279a5bb94e5f639a8f1ac67f1b09b7fd32d49bbfe0e15e8258a70ae9ebb9316cef0f9b6feb828a5fc6a2452884610208565b6104f560c08801357f1823654e6782371e622b0629518c890a115d6d2cae9bb2c55180143aa0212bd77f2f894980bdbd71026b861c9c9feaf6cd797e28188a9b4154ba63ce83f36a29f284610208565b61054460e08801357e6a9965248ae4075ea54dc532370d14ff99d0c686fc5c23b82176baf60dbb7b7f013d2f682053227157adb588e8533b32e74b1e042d55f83933244f08650abdbd84610208565b6105956101008801357f1c09ec2c3abe8e173ecbe443fd8e7c2f5ec2024843440235cc1196ee8070f78c7f1c9cc75a83a50251238fc513eea4f2effb5e26d5b592b48250ed3e5c9759179084610208565b6105e66101208801357f0e9f0a02b712bb78f13b08c8a99ded7165ed62f95feb1c4c0573073b63aca9217f18917ec1f1868dadff1a3196e6804296cac2d322e29574a93c37f25e4003c1ca84610208565b6106376101408801357f0bc5e6897fac05c15be2a6dbe4ab2a5eb0fb51a5d67efb83a582dc5ae2d410e47f1e18dba9927d3c91dbcba83e9cb8f533b2c9c9282fa808f11d5f7bcf404b142f84610208565b833582527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760208501357f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703066020830152843560408301526020850135606083015260408501356080830152606085013560a08301527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e260c08301527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d192660e08301527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101008301527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101208301527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101408301527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec8610160830152600088015161018083015260206000018801516101a08301527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26101c08301527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6101e08301527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102008301527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610220830152853561024083015260208601356102608301527f1642e1e5564fdfd598d06d6c6c0bcb91104f0b5be1af80ee3a05cbb870931bd56102808301527f2eaa248f26eb631ada70bf8472546eb9deaa6d10d2bbdf452eeac795bcd4cba86102a08301527f0aba69025c669be9ea7cfc27e650a4f33580b7fbed2c8c48254895820422c08f6102c08301527f12ec59350afb7e582f8d6c48d20afe0343cf2396a8af0d260ed0b7e5c0baf1386102e08301526020826103008460086107d05a03fa82518116935050505095945050505050565b604051610380810160405261093360008401356101d4565b61094060208401356101d4565b61094d60408401356101d4565b61095a60608401356101d4565b61096760808401356101d4565b61097460a08401356101d4565b61098160c08401356101d4565b61098e60e08401356101d4565b61099c6101008401356101d4565b6109aa6101208401356101d4565b6109b86101408401356101d4565b6109c66101608401356101d4565b6109d3818486888a610273565b8060005260206000f35b604051806101600160405280600b90602082028036833780820191505090505090565b600080fd5b600080fd5b600080fd5b600081905082602060020282011115610a2b57610a2a610a0a565b5b92915050565b600081905082604060020282011115610a4d57610a4c610a0a565b5b92915050565b600080fd5b600080fd5b60008083601f840112610a7357610a72610a53565b5b8235905067ffffffffffffffff811115610a9057610a8f610a58565b5b602083019150836020820283011115610aac57610aab610a0a565b5b9250929050565b60008060008060006101208688031215610ad057610acf610a00565b5b6000610ade88828901610a0f565b9550506040610aef88828901610a31565b94505060c0610b0088828901610a0f565b93505061010086013567ffffffffffffffff811115610b2257610b21610a05565b5b610b2e88828901610a5d565b92509250509295509295909350565b60008115159050919050565b610b5281610b3d565b82525050565b6000602082019050610b6d6000830184610b49565b92915050565b6000819050826020600b0282011115610b8f57610b8e610a0a565b5b92915050565b6000806000806102608587031215610bb057610baf610a00565b5b6000610bbe87828801610a0f565b9450506040610bcf87828801610a31565b93505060c0610be087828801610a0f565b925050610100610bf287828801610b73565b91505092959194509250565b600082825260208201905092915050565b7f6578706563746564206172726179206c656e6774682069732031310000000000600082015250565b6000610c45601b83610bfe565b9150610c5082610c0f565b602082019050919050565b60006020820190508181036000830152610c7481610c38565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b6000610cee82610cd9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d2057610d1f610caa565b5b600182019050919050565b82818337505050565b610d4060408383610d2b565b5050565b600060029050919050565b600081905092915050565b6000819050919050565b610d7060408383610d2b565b5050565b6000610d808383610d64565b60408301905092915050565b600082905092915050565b6000604082019050919050565b610dad81610d44565b610db78184610d4f565b9250610dc282610d5a565b8060005b83811015610dfb57610dd88284610d8c565b610de28782610d74565b9650610ded83610d97565b925050600181019050610dc6565b505050505050565b6000600b9050919050565b600081905092915050565b6000819050919050565b610e2c81610cd9565b82525050565b6000610e3e8383610e23565b60208301905092915050565b6000602082019050919050565b610e6081610e03565b610e6a8184610e0e565b9250610e7582610e19565b8060005b83811015610ea6578151610e8d8782610e32565b9650610e9883610e4a565b925050600181019050610e79565b505050505050565b600061026082019050610ec46000830187610d34565b610ed16040830186610da4565b610ede60c0830185610d34565b610eec610100830184610e57565b95945050505050565b610efe81610b3d565b8114610f0957600080fd5b50565b600081519050610f1b81610ef5565b92915050565b600060208284031215610f3757610f36610a00565b5b6000610f4584828501610f0c565b9150509291505056fea2646970667358221220ae90c4e25b1c95e278e3472a4e6693bf4e4508de26bcf6702f2f1aaa34f6fcfd64736f6c63430008100033",
}

// VerifiermtpwrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifiermtpwrapperMetaData.ABI instead.
var VerifiermtpwrapperABI = VerifiermtpwrapperMetaData.ABI

// VerifiermtpwrapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifiermtpwrapperMetaData.Bin instead.
var VerifiermtpwrapperBin = VerifiermtpwrapperMetaData.Bin

// DeployVerifiermtpwrapper deploys a new Ethereum contract, binding an instance of Verifiermtpwrapper to it.
func DeployVerifiermtpwrapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifiermtpwrapper, error) {
	parsed, err := VerifiermtpwrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifiermtpwrapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifiermtpwrapper{VerifiermtpwrapperCaller: VerifiermtpwrapperCaller{contract: contract}, VerifiermtpwrapperTransactor: VerifiermtpwrapperTransactor{contract: contract}, VerifiermtpwrapperFilterer: VerifiermtpwrapperFilterer{contract: contract}}, nil
}

// Verifiermtpwrapper is an auto generated Go binding around an Ethereum contract.
type Verifiermtpwrapper struct {
	VerifiermtpwrapperCaller     // Read-only binding to the contract
	VerifiermtpwrapperTransactor // Write-only binding to the contract
	VerifiermtpwrapperFilterer   // Log filterer for contract events
}

// VerifiermtpwrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifiermtpwrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiermtpwrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifiermtpwrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiermtpwrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifiermtpwrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiermtpwrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifiermtpwrapperSession struct {
	Contract     *Verifiermtpwrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifiermtpwrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifiermtpwrapperCallerSession struct {
	Contract *VerifiermtpwrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VerifiermtpwrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifiermtpwrapperTransactorSession struct {
	Contract     *VerifiermtpwrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VerifiermtpwrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifiermtpwrapperRaw struct {
	Contract *Verifiermtpwrapper // Generic contract binding to access the raw methods on
}

// VerifiermtpwrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifiermtpwrapperCallerRaw struct {
	Contract *VerifiermtpwrapperCaller // Generic read-only contract binding to access the raw methods on
}

// VerifiermtpwrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifiermtpwrapperTransactorRaw struct {
	Contract *VerifiermtpwrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifiermtpwrapper creates a new instance of Verifiermtpwrapper, bound to a specific deployed contract.
func NewVerifiermtpwrapper(address common.Address, backend bind.ContractBackend) (*Verifiermtpwrapper, error) {
	contract, err := bindVerifiermtpwrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifiermtpwrapper{VerifiermtpwrapperCaller: VerifiermtpwrapperCaller{contract: contract}, VerifiermtpwrapperTransactor: VerifiermtpwrapperTransactor{contract: contract}, VerifiermtpwrapperFilterer: VerifiermtpwrapperFilterer{contract: contract}}, nil
}

// NewVerifiermtpwrapperCaller creates a new read-only instance of Verifiermtpwrapper, bound to a specific deployed contract.
func NewVerifiermtpwrapperCaller(address common.Address, caller bind.ContractCaller) (*VerifiermtpwrapperCaller, error) {
	contract, err := bindVerifiermtpwrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiermtpwrapperCaller{contract: contract}, nil
}

// NewVerifiermtpwrapperTransactor creates a new write-only instance of Verifiermtpwrapper, bound to a specific deployed contract.
func NewVerifiermtpwrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifiermtpwrapperTransactor, error) {
	contract, err := bindVerifiermtpwrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiermtpwrapperTransactor{contract: contract}, nil
}

// NewVerifiermtpwrapperFilterer creates a new log filterer instance of Verifiermtpwrapper, bound to a specific deployed contract.
func NewVerifiermtpwrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifiermtpwrapperFilterer, error) {
	contract, err := bindVerifiermtpwrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifiermtpwrapperFilterer{contract: contract}, nil
}

// bindVerifiermtpwrapper binds a generic wrapper to an already deployed contract.
func bindVerifiermtpwrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifiermtpwrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifiermtpwrapper *VerifiermtpwrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifiermtpwrapper.Contract.VerifiermtpwrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifiermtpwrapper *VerifiermtpwrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifiermtpwrapper.Contract.VerifiermtpwrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifiermtpwrapper *VerifiermtpwrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifiermtpwrapper.Contract.VerifiermtpwrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifiermtpwrapper *VerifiermtpwrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifiermtpwrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifiermtpwrapper *VerifiermtpwrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifiermtpwrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifiermtpwrapper *VerifiermtpwrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifiermtpwrapper.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifiermtpwrapper *VerifiermtpwrapperCaller) Verify(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifiermtpwrapper.contract.Call(opts, &out, "verify", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifiermtpwrapper *VerifiermtpwrapperSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	return _Verifiermtpwrapper.Contract.Verify(&_Verifiermtpwrapper.CallOpts, a, b, c, input)
}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifiermtpwrapper *VerifiermtpwrapperCallerSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	return _Verifiermtpwrapper.Contract.Verify(&_Verifiermtpwrapper.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[11] _pubSignals) view returns(bool)
func (_Verifiermtpwrapper *VerifiermtpwrapperCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [11]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifiermtpwrapper.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[11] _pubSignals) view returns(bool)
func (_Verifiermtpwrapper *VerifiermtpwrapperSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [11]*big.Int) (bool, error) {
	return _Verifiermtpwrapper.Contract.VerifyProof(&_Verifiermtpwrapper.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[11] _pubSignals) view returns(bool)
func (_Verifiermtpwrapper *VerifiermtpwrapperCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [11]*big.Int) (bool, error) {
	return _Verifiermtpwrapper.Contract.VerifyProof(&_Verifiermtpwrapper.CallOpts, _pA, _pB, _pC, _pubSignals)
}
