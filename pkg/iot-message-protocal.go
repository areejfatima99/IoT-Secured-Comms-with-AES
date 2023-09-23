// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pkg

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

// IoTMessageProtocolMessage is an auto generated low-level Go binding around an user-defined struct.
type IoTMessageProtocolMessage struct {
	Sender           common.Address
	Recipient        common.Address
	EncryptedMessage []byte
}

// PkgMetaData contains all meta data concerning the Pkg contract.
var PkgMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encryptedMessage\",\"type\":\"bytes\"}],\"internalType\":\"structIoTMessageProtocol.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messages\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encryptedMessage\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_encryptedMessage\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506106b58061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80630d80fefd146100435780635ff6cbf31461006e578063bb5ddb0f14610083575b5f80fd5b61005661005136600461034d565b610098565b604051610065939291906103a7565b60405180910390f35b61007661015f565b60405161006591906103db565b610096610091366004610478565b61026e565b005b5f81815481106100a6575f80fd5b5f9182526020909120600390910201805460018201546002830180546001600160a01b0393841695509290911692916100de90610541565b80601f016020809104026020016040519081016040528092919081815260200182805461010a90610541565b80156101555780601f1061012c57610100808354040283529160200191610155565b820191905f5260205f20905b81548152906001019060200180831161013857829003601f168201915b5050505050905083565b60605f805480602002602001604051908101604052809291908181526020015f905b82821015610265575f848152602090819020604080516060810182526003860290920180546001600160a01b03908116845260018201541693830193909352600283018054929392918401916101d690610541565b80601f016020809104026020016040519081016040528092919081815260200182805461020290610541565b801561024d5780601f106102245761010080835404028352916020019161024d565b820191905f5260205f20905b81548152906001019060200180831161023057829003601f168201915b50505050508152505081526020019060010190610181565b50505050905090565b604080516060810182523381526001600160a01b03848116602083019081529282018481525f8054600181018255908052835160039091027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563810180549285166001600160a01b031993841617815595517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5648201805491909516921691909117909255519192839290917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e565019061034590826105c3565b505050505050565b5f6020828403121561035d575f80fd5b5035919050565b5f81518084525f5b818110156103885760208185018101518683018201520161036c565b505f602082860101526020601f19601f83011685010191505092915050565b6001600160a01b038481168252831660208201526060604082018190525f906103d290830184610364565b95945050505050565b5f6020808301818452808551808352604092508286019150828160051b8701018488015f5b8381101561045657888303603f19018552815180516001600160a01b039081168552888201511688850152860151606087850181905261044281860183610364565b968901969450505090860190600101610400565b509098975050505050505050565b634e487b7160e01b5f52604160045260245ffd5b5f8060408385031215610489575f80fd5b82356001600160a01b038116811461049f575f80fd5b9150602083013567ffffffffffffffff808211156104bb575f80fd5b818501915085601f8301126104ce575f80fd5b8135818111156104e0576104e0610464565b604051601f8201601f19908116603f0116810190838211818310171561050857610508610464565b81604052828152886020848701011115610520575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b600181811c9082168061055557607f821691505b60208210810361057357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156105be575f81815260208120601f850160051c8101602086101561059f5750805b601f850160051c820191505b81811015610345578281556001016105ab565b505050565b815167ffffffffffffffff8111156105dd576105dd610464565b6105f1816105eb8454610541565b84610579565b602080601f831160018114610624575f841561060d5750858301515b5f19600386901b1c1916600185901b178555610345565b5f85815260208120601f198616915b8281101561065257888601518255948401946001909101908401610633565b508582101561066f57878501515f19600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220220a5aa1057639514cbdbd060db9458d5d865ac039518a04b55bdb83b117254464736f6c63430008150033",
}

// PkgABI is the input ABI used to generate the binding from.
// Deprecated: Use PkgMetaData.ABI instead.
var PkgABI = PkgMetaData.ABI

// PkgBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PkgMetaData.Bin instead.
var PkgBin = PkgMetaData.Bin

// DeployPkg deploys a new Ethereum contract, binding an instance of Pkg to it.
func DeployPkg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pkg, error) {
	parsed, err := PkgMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PkgBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pkg{PkgCaller: PkgCaller{contract: contract}, PkgTransactor: PkgTransactor{contract: contract}, PkgFilterer: PkgFilterer{contract: contract}}, nil
}

// Pkg is an auto generated Go binding around an Ethereum contract.
type Pkg struct {
	PkgCaller     // Read-only binding to the contract
	PkgTransactor // Write-only binding to the contract
	PkgFilterer   // Log filterer for contract events
}

// PkgCaller is an auto generated read-only Go binding around an Ethereum contract.
type PkgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PkgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PkgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PkgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PkgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PkgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PkgSession struct {
	Contract     *Pkg              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PkgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PkgCallerSession struct {
	Contract *PkgCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PkgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PkgTransactorSession struct {
	Contract     *PkgTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PkgRaw is an auto generated low-level Go binding around an Ethereum contract.
type PkgRaw struct {
	Contract *Pkg // Generic contract binding to access the raw methods on
}

// PkgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PkgCallerRaw struct {
	Contract *PkgCaller // Generic read-only contract binding to access the raw methods on
}

// PkgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PkgTransactorRaw struct {
	Contract *PkgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPkg creates a new instance of Pkg, bound to a specific deployed contract.
func NewPkg(address common.Address, backend bind.ContractBackend) (*Pkg, error) {
	contract, err := bindPkg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pkg{PkgCaller: PkgCaller{contract: contract}, PkgTransactor: PkgTransactor{contract: contract}, PkgFilterer: PkgFilterer{contract: contract}}, nil
}

// NewPkgCaller creates a new read-only instance of Pkg, bound to a specific deployed contract.
func NewPkgCaller(address common.Address, caller bind.ContractCaller) (*PkgCaller, error) {
	contract, err := bindPkg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PkgCaller{contract: contract}, nil
}

// NewPkgTransactor creates a new write-only instance of Pkg, bound to a specific deployed contract.
func NewPkgTransactor(address common.Address, transactor bind.ContractTransactor) (*PkgTransactor, error) {
	contract, err := bindPkg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PkgTransactor{contract: contract}, nil
}

// NewPkgFilterer creates a new log filterer instance of Pkg, bound to a specific deployed contract.
func NewPkgFilterer(address common.Address, filterer bind.ContractFilterer) (*PkgFilterer, error) {
	contract, err := bindPkg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PkgFilterer{contract: contract}, nil
}

// bindPkg binds a generic wrapper to an already deployed contract.
func bindPkg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PkgMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pkg *PkgRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pkg.Contract.PkgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pkg *PkgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pkg.Contract.PkgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pkg *PkgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pkg.Contract.PkgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pkg *PkgCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pkg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pkg *PkgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pkg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pkg *PkgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pkg.Contract.contract.Transact(opts, method, params...)
}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((address,address,bytes)[])
func (_Pkg *PkgCaller) GetMessages(opts *bind.CallOpts) ([]IoTMessageProtocolMessage, error) {
	var out []interface{}
	err := _Pkg.contract.Call(opts, &out, "getMessages")

	if err != nil {
		return *new([]IoTMessageProtocolMessage), err
	}

	out0 := *abi.ConvertType(out[0], new([]IoTMessageProtocolMessage)).(*[]IoTMessageProtocolMessage)

	return out0, err

}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((address,address,bytes)[])
func (_Pkg *PkgSession) GetMessages() ([]IoTMessageProtocolMessage, error) {
	return _Pkg.Contract.GetMessages(&_Pkg.CallOpts)
}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((address,address,bytes)[])
func (_Pkg *PkgCallerSession) GetMessages() ([]IoTMessageProtocolMessage, error) {
	return _Pkg.Contract.GetMessages(&_Pkg.CallOpts)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(address sender, address recipient, bytes encryptedMessage)
func (_Pkg *PkgCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sender           common.Address
	Recipient        common.Address
	EncryptedMessage []byte
}, error) {
	var out []interface{}
	err := _Pkg.contract.Call(opts, &out, "messages", arg0)

	outstruct := new(struct {
		Sender           common.Address
		Recipient        common.Address
		EncryptedMessage []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.EncryptedMessage = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(address sender, address recipient, bytes encryptedMessage)
func (_Pkg *PkgSession) Messages(arg0 *big.Int) (struct {
	Sender           common.Address
	Recipient        common.Address
	EncryptedMessage []byte
}, error) {
	return _Pkg.Contract.Messages(&_Pkg.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(address sender, address recipient, bytes encryptedMessage)
func (_Pkg *PkgCallerSession) Messages(arg0 *big.Int) (struct {
	Sender           common.Address
	Recipient        common.Address
	EncryptedMessage []byte
}, error) {
	return _Pkg.Contract.Messages(&_Pkg.CallOpts, arg0)
}

// SendMessage is a paid mutator transaction binding the contract method 0xbb5ddb0f.
//
// Solidity: function sendMessage(address _recipient, bytes _encryptedMessage) returns()
func (_Pkg *PkgTransactor) SendMessage(opts *bind.TransactOpts, _recipient common.Address, _encryptedMessage []byte) (*types.Transaction, error) {
	return _Pkg.contract.Transact(opts, "sendMessage", _recipient, _encryptedMessage)
}

// SendMessage is a paid mutator transaction binding the contract method 0xbb5ddb0f.
//
// Solidity: function sendMessage(address _recipient, bytes _encryptedMessage) returns()
func (_Pkg *PkgSession) SendMessage(_recipient common.Address, _encryptedMessage []byte) (*types.Transaction, error) {
	return _Pkg.Contract.SendMessage(&_Pkg.TransactOpts, _recipient, _encryptedMessage)
}

// SendMessage is a paid mutator transaction binding the contract method 0xbb5ddb0f.
//
// Solidity: function sendMessage(address _recipient, bytes _encryptedMessage) returns()
func (_Pkg *PkgTransactorSession) SendMessage(_recipient common.Address, _encryptedMessage []byte) (*types.Transaction, error) {
	return _Pkg.Contract.SendMessage(&_Pkg.TransactOpts, _recipient, _encryptedMessage)
}
