// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifierstatetransition

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

// VerifierstatetransitionMetaData contains all meta data concerning the Verifierstatetransition contract.
var VerifierstatetransitionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_pubSignals\",\"type\":\"uint256[4]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610724806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80635fe8c13b14610030575b600080fd5b61004a6004803603810190610045919061064f565b610060565b60405161005791906106d3565b60405180910390f35b600061057c565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478110610098576000805260206000f35b50565b600060405183815284602082015285604082015260408160608360076107d05a03fa9150816100ce576000805260206000f35b825160408201526020830151606082015260408360808360066107d05a03fa9150816100fe576000805260206000f35b505050505050565b600060808601600087017f23782e74ab120e17332feb501bcb3340e46aa66d97d227ce10d3b0d6191f155581527f2549073e43bec0249113913f540b1c803eb0fceb45181b4370b694352b8626d560208201526101a960008801357f2df8f4ffecaca21322387e2814d5cbd3e9ea882b9870155f6a09b4473d4c80797f233b3e0bdba80ff1393fe3ee9d1483adf633903b226771e31fb3305e2774e0fa8461009b565b6101f960208801357f1e069468e84e4443468d989b6fc64820336b165e815f0cce7c344ab64c87acd67f23183180476734a897d90341b322f6226879f47d5705624f625a9f81b12be97c8461009b565b61024860408801357f0b766e26319340b06fb428ab1241f5d2c84220fc609dbbfacda8e85b4c62e8ef7ef6ce3d6de122311ef96a588c41d42a0e8aaabd531dd2bfbbe23d00b3ece1b98461009b565b61029860608801357f13ae53c8f7d53cda6035e94c207c087599487be203a9fded64d9b803843420547f037ec8ba945630f6d6b5387f0a4a130ddd6f216ea478bc08079a6a0d13417eab8461009b565b833582527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760208501357f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703066020830152843560408301526020850135606083015260408501356080830152606085013560a08301527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e260c08301527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d192660e08301527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101008301527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101208301527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101408301527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec8610160830152600088015161018083015260206000018801516101a08301527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26101c08301527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6101e08301527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102008301527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610220830152853561024083015260208601356102608301527f09633cb521508aaac061ec58387116514b5ca8a5ef113b973dc5947cfd07c22d6102808301527f1a008cf3fb47fe77514ec4146a500e2ea851479d91c6a8286ffae46e737793856102a08301527f2daeaf590139f49f18bea7f56c980627d13b44525c4af2fe0d5e66adfdd09cb26102c08301527f1517928b5d86d1c6fa6d777090519971ea5a5b58173252da0677879848a30c316102e08301526020826103008460086107d05a03fa82518116935050505095945050505050565b60405161038081016040526105946000840135610067565b6105a16020840135610067565b6105ae6040840135610067565b6105bb6060840135610067565b6105c86080840135610067565b6105d5818486888a610106565b8060005260206000f35b600080fd5b600080fd5b600081905082602060020282011115610605576106046105e4565b5b92915050565b600081905082604060020282011115610627576106266105e4565b5b92915050565b600081905082602060040282011115610649576106486105e4565b5b92915050565b600080600080610180858703121561066a576106696105df565b5b6000610678878288016105e9565b94505060406106898782880161060b565b93505060c061069a878288016105e9565b9250506101006106ac8782880161062d565b91505092959194509250565b60008115159050919050565b6106cd816106b8565b82525050565b60006020820190506106e860008301846106c4565b9291505056fea264697066735822122059c28832b616e4f6a684f7d68815ab5752c88e6234ad98d01bb1f96fa3a8b8d764736f6c63430008100033",
}

// VerifierstatetransitionABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierstatetransitionMetaData.ABI instead.
var VerifierstatetransitionABI = VerifierstatetransitionMetaData.ABI

// VerifierstatetransitionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifierstatetransitionMetaData.Bin instead.
var VerifierstatetransitionBin = VerifierstatetransitionMetaData.Bin

// DeployVerifierstatetransition deploys a new Ethereum contract, binding an instance of Verifierstatetransition to it.
func DeployVerifierstatetransition(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifierstatetransition, error) {
	parsed, err := VerifierstatetransitionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierstatetransitionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifierstatetransition{VerifierstatetransitionCaller: VerifierstatetransitionCaller{contract: contract}, VerifierstatetransitionTransactor: VerifierstatetransitionTransactor{contract: contract}, VerifierstatetransitionFilterer: VerifierstatetransitionFilterer{contract: contract}}, nil
}

// Verifierstatetransition is an auto generated Go binding around an Ethereum contract.
type Verifierstatetransition struct {
	VerifierstatetransitionCaller     // Read-only binding to the contract
	VerifierstatetransitionTransactor // Write-only binding to the contract
	VerifierstatetransitionFilterer   // Log filterer for contract events
}

// VerifierstatetransitionCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierstatetransitionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierstatetransitionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierstatetransitionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierstatetransitionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierstatetransitionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierstatetransitionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierstatetransitionSession struct {
	Contract     *Verifierstatetransition // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VerifierstatetransitionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierstatetransitionCallerSession struct {
	Contract *VerifierstatetransitionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// VerifierstatetransitionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierstatetransitionTransactorSession struct {
	Contract     *VerifierstatetransitionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// VerifierstatetransitionRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierstatetransitionRaw struct {
	Contract *Verifierstatetransition // Generic contract binding to access the raw methods on
}

// VerifierstatetransitionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierstatetransitionCallerRaw struct {
	Contract *VerifierstatetransitionCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierstatetransitionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierstatetransitionTransactorRaw struct {
	Contract *VerifierstatetransitionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierstatetransition creates a new instance of Verifierstatetransition, bound to a specific deployed contract.
func NewVerifierstatetransition(address common.Address, backend bind.ContractBackend) (*Verifierstatetransition, error) {
	contract, err := bindVerifierstatetransition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifierstatetransition{VerifierstatetransitionCaller: VerifierstatetransitionCaller{contract: contract}, VerifierstatetransitionTransactor: VerifierstatetransitionTransactor{contract: contract}, VerifierstatetransitionFilterer: VerifierstatetransitionFilterer{contract: contract}}, nil
}

// NewVerifierstatetransitionCaller creates a new read-only instance of Verifierstatetransition, bound to a specific deployed contract.
func NewVerifierstatetransitionCaller(address common.Address, caller bind.ContractCaller) (*VerifierstatetransitionCaller, error) {
	contract, err := bindVerifierstatetransition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierstatetransitionCaller{contract: contract}, nil
}

// NewVerifierstatetransitionTransactor creates a new write-only instance of Verifierstatetransition, bound to a specific deployed contract.
func NewVerifierstatetransitionTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierstatetransitionTransactor, error) {
	contract, err := bindVerifierstatetransition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierstatetransitionTransactor{contract: contract}, nil
}

// NewVerifierstatetransitionFilterer creates a new log filterer instance of Verifierstatetransition, bound to a specific deployed contract.
func NewVerifierstatetransitionFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierstatetransitionFilterer, error) {
	contract, err := bindVerifierstatetransition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierstatetransitionFilterer{contract: contract}, nil
}

// bindVerifierstatetransition binds a generic wrapper to an already deployed contract.
func bindVerifierstatetransition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierstatetransitionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifierstatetransition *VerifierstatetransitionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifierstatetransition.Contract.VerifierstatetransitionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifierstatetransition *VerifierstatetransitionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifierstatetransition.Contract.VerifierstatetransitionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifierstatetransition *VerifierstatetransitionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifierstatetransition.Contract.VerifierstatetransitionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifierstatetransition *VerifierstatetransitionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifierstatetransition.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifierstatetransition *VerifierstatetransitionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifierstatetransition.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifierstatetransition *VerifierstatetransitionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifierstatetransition.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe8c13b.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[4] _pubSignals) view returns(bool)
func (_Verifierstatetransition *VerifierstatetransitionCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifierstatetransition.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe8c13b.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[4] _pubSignals) view returns(bool)
func (_Verifierstatetransition *VerifierstatetransitionSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [4]*big.Int) (bool, error) {
	return _Verifierstatetransition.Contract.VerifyProof(&_Verifierstatetransition.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe8c13b.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[4] _pubSignals) view returns(bool)
func (_Verifierstatetransition *VerifierstatetransitionCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [4]*big.Int) (bool, error) {
	return _Verifierstatetransition.Contract.VerifyProof(&_Verifierstatetransition.CallOpts, _pA, _pB, _pC, _pubSignals)
}
