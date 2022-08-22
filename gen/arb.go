// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arb

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
)

// ArbMetaData contains all meta data concerning the Arb contract.
var ArbMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"dualDexTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"estimateDualDexTrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token3\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"estimateTriDexTrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getAmountOutMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61177e8061010d6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063715018a611610066578063715018a61461016b5780638da5cb5b14610175578063bcdb446b14610193578063f2fde38b1461019d578063f8b2cb4f146101b95761009e565b8063068e7ca1146100a357806316114acd146100d35780631aa51318146100ef57806326a215ec1461010b57806343cf6f241461013b575b600080fd5b6100bd60048036038101906100b89190610ee1565b6101e9565b6040516100ca9190610f6b565b60405180910390f35b6100ed60048036038101906100e89190610f86565b610218565b005b61010960048036038101906101049190610ee1565b610395565b005b61012560048036038101906101209190610fb3565b61067f565b6040516101329190610f6b565b60405180910390f35b61015560048036038101906101509190611055565b6106c1565b6040516101629190610f6b565b60405180910390f35b610173610868565b005b61017d6108f0565b60405161018a91906110cb565b60405180910390f35b61019b610919565b005b6101b760048036038101906101b29190610f86565b6109de565b005b6101d360048036038101906101ce9190610f86565b610ad5565b6040516101e09190610f6b565b60405180910390f35b6000806101f8878686866106c1565b90506000610208878688856106c1565b9050809250505095945050505050565b610220610b5d565b73ffffffffffffffffffffffffffffffffffffffff1661023e6108f0565b73ffffffffffffffffffffffffffffffffffffffff1614610294576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028b90611143565b60405180910390fd5b60008190508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb338373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102ef91906110cb565b602060405180830381865afa15801561030c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103309190611178565b6040518363ffffffff1660e01b815260040161034d9291906111a5565b6020604051808303816000875af115801561036c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103909190611206565b505050565b61039d610b5d565b73ffffffffffffffffffffffffffffffffffffffff166103bb6108f0565b73ffffffffffffffffffffffffffffffffffffffff1614610411576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040890611143565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161044c91906110cb565b602060405180830381865afa158015610469573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048d9190611178565b905060008373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016104ca91906110cb565b602060405180830381865afa1580156104e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050b9190611178565b905061051987868686610b65565b60008473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161055491906110cb565b602060405180830381865afa158015610571573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105959190611178565b9050600082826105a59190611262565b90506105b388878984610b65565b60008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016105ee91906110cb565b602060405180830381865afa15801561060b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062f9190611178565b9050848111610673576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066a906112e2565b60405180910390fd5b50505050505050505050565b60008061068e898787866106c1565b9050600061069e898787856106c1565b905060006106ae89878a856106c1565b9050809350505050979650505050505050565b60006060600267ffffffffffffffff8111156106e0576106df611302565b5b60405190808252806020026020018201604052801561070e5781602001602082028036833780820191505090505b509050848160008151811061072657610725611331565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050838160018151811061077557610774611331565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060008673ffffffffffffffffffffffffffffffffffffffff1663d06ca61f85846040518363ffffffff1660e01b81526004016107ec92919061141e565b600060405180830381865afa158015610809573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108329190611578565b905080600183516108439190611262565b8151811061085457610853611331565b5b602002602001015192505050949350505050565b610870610b5d565b73ffffffffffffffffffffffffffffffffffffffff1661088e6108f0565b73ffffffffffffffffffffffffffffffffffffffff16146108e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108db90611143565b60405180910390fd5b6108ee6000610d75565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610921610b5d565b73ffffffffffffffffffffffffffffffffffffffff1661093f6108f0565b73ffffffffffffffffffffffffffffffffffffffff1614610995576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098c90611143565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156109db573d6000803e3d6000fd5b50565b6109e6610b5d565b73ffffffffffffffffffffffffffffffffffffffff16610a046108f0565b73ffffffffffffffffffffffffffffffffffffffff1614610a5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5190611143565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ac9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac090611633565b60405180910390fd5b610ad281610d75565b50565b6000808273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610b1191906110cb565b602060405180830381865afa158015610b2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b529190611178565b905080915050919050565b600033905090565b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b385836040518363ffffffff1660e01b8152600401610ba09291906111a5565b6020604051808303816000875af1158015610bbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be39190611206565b506060600267ffffffffffffffff811115610c0157610c00611302565b5b604051908082528060200260200182016040528015610c2f5781602001602082028036833780820191505090505b5090508381600081518110610c4757610c46611331565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281600181518110610c9657610c95611331565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600061012c42610ce09190611653565b90508573ffffffffffffffffffffffffffffffffffffffff166338ed17398460018530866040518663ffffffff1660e01b8152600401610d249594939291906116ee565b6000604051808303816000875af1158015610d43573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610d6c9190611578565b50505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e7882610e4d565b9050919050565b610e8881610e6d565b8114610e9357600080fd5b50565b600081359050610ea581610e7f565b92915050565b6000819050919050565b610ebe81610eab565b8114610ec957600080fd5b50565b600081359050610edb81610eb5565b92915050565b600080600080600060a08688031215610efd57610efc610e43565b5b6000610f0b88828901610e96565b9550506020610f1c88828901610e96565b9450506040610f2d88828901610e96565b9350506060610f3e88828901610e96565b9250506080610f4f88828901610ecc565b9150509295509295909350565b610f6581610eab565b82525050565b6000602082019050610f806000830184610f5c565b92915050565b600060208284031215610f9c57610f9b610e43565b5b6000610faa84828501610e96565b91505092915050565b600080600080600080600060e0888a031215610fd257610fd1610e43565b5b6000610fe08a828b01610e96565b9750506020610ff18a828b01610e96565b96505060406110028a828b01610e96565b95505060606110138a828b01610e96565b94505060806110248a828b01610e96565b93505060a06110358a828b01610e96565b92505060c06110468a828b01610ecc565b91505092959891949750929550565b6000806000806080858703121561106f5761106e610e43565b5b600061107d87828801610e96565b945050602061108e87828801610e96565b935050604061109f87828801610e96565b92505060606110b087828801610ecc565b91505092959194509250565b6110c581610e6d565b82525050565b60006020820190506110e060008301846110bc565b92915050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061112d6020836110e6565b9150611138826110f7565b602082019050919050565b6000602082019050818103600083015261115c81611120565b9050919050565b60008151905061117281610eb5565b92915050565b60006020828403121561118e5761118d610e43565b5b600061119c84828501611163565b91505092915050565b60006040820190506111ba60008301856110bc565b6111c76020830184610f5c565b9392505050565b60008115159050919050565b6111e3816111ce565b81146111ee57600080fd5b50565b600081519050611200816111da565b92915050565b60006020828403121561121c5761121b610e43565b5b600061122a848285016111f1565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061126d82610eab565b915061127883610eab565b92508282101561128b5761128a611233565b5b828203905092915050565b7f54726164652052657665727465642c204e6f2050726f666974204d6164650000600082015250565b60006112cc601e836110e6565b91506112d782611296565b602082019050919050565b600060208201905081810360008301526112fb816112bf565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61139581610e6d565b82525050565b60006113a7838361138c565b60208301905092915050565b6000602082019050919050565b60006113cb82611360565b6113d5818561136b565b93506113e08361137c565b8060005b838110156114115781516113f8888261139b565b9750611403836113b3565b9250506001810190506113e4565b5085935050505092915050565b60006040820190506114336000830185610f5c565b818103602083015261144581846113c0565b90509392505050565b600080fd5b6000601f19601f8301169050919050565b61146d82611453565b810181811067ffffffffffffffff8211171561148c5761148b611302565b5b80604052505050565b600061149f610e39565b90506114ab8282611464565b919050565b600067ffffffffffffffff8211156114cb576114ca611302565b5b602082029050602081019050919050565b600080fd5b60006114f46114ef846114b0565b611495565b90508083825260208201905060208402830185811115611517576115166114dc565b5b835b81811015611540578061152c8882611163565b845260208401935050602081019050611519565b5050509392505050565b600082601f83011261155f5761155e61144e565b5b815161156f8482602086016114e1565b91505092915050565b60006020828403121561158e5761158d610e43565b5b600082015167ffffffffffffffff8111156115ac576115ab610e48565b5b6115b88482850161154a565b91505092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061161d6026836110e6565b9150611628826115c1565b604082019050919050565b6000602082019050818103600083015261164c81611610565b9050919050565b600061165e82610eab565b915061166983610eab565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561169e5761169d611233565b5b828201905092915050565b6000819050919050565b6000819050919050565b60006116d86116d36116ce846116a9565b6116b3565b610eab565b9050919050565b6116e8816116bd565b82525050565b600060a0820190506117036000830188610f5c565b61171060208301876116df565b818103604083015261172281866113c0565b905061173160608301856110bc565b61173e6080830184610f5c565b969550505050505056fea26469706673582212206fdacfaa2c2bfbc4986da3c96071df342f378aa651d8735b45ecd8dce4f6afcb64736f6c634300080d0033",
}

// ArbABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbMetaData.ABI instead.
var ArbABI = ArbMetaData.ABI

// ArbBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArbMetaData.Bin instead.
var ArbBin = ArbMetaData.Bin

// DeployArb deploys a new Ethereum contract, binding an instance of Arb to it.
func DeployArb(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Arb, error) {
	parsed, err := ArbMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Arb{ArbCaller: ArbCaller{contract: contract}, ArbTransactor: ArbTransactor{contract: contract}, ArbFilterer: ArbFilterer{contract: contract}}, nil
}

// Arb is an auto generated Go binding around an Ethereum contract.
type Arb struct {
	ArbCaller     // Read-only binding to the contract
	ArbTransactor // Write-only binding to the contract
	ArbFilterer   // Log filterer for contract events
}

// ArbCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbSession struct {
	Contract     *Arb              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbCallerSession struct {
	Contract *ArbCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbTransactorSession struct {
	Contract     *ArbTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbRaw struct {
	Contract *Arb // Generic contract binding to access the raw methods on
}

// ArbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbCallerRaw struct {
	Contract *ArbCaller // Generic read-only contract binding to access the raw methods on
}

// ArbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbTransactorRaw struct {
	Contract *ArbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArb creates a new instance of Arb, bound to a specific deployed contract.
func NewArb(address common.Address, backend bind.ContractBackend) (*Arb, error) {
	contract, err := bindArb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arb{ArbCaller: ArbCaller{contract: contract}, ArbTransactor: ArbTransactor{contract: contract}, ArbFilterer: ArbFilterer{contract: contract}}, nil
}

// NewArbCaller creates a new read-only instance of Arb, bound to a specific deployed contract.
func NewArbCaller(address common.Address, caller bind.ContractCaller) (*ArbCaller, error) {
	contract, err := bindArb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbCaller{contract: contract}, nil
}

// NewArbTransactor creates a new write-only instance of Arb, bound to a specific deployed contract.
func NewArbTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbTransactor, error) {
	contract, err := bindArb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbTransactor{contract: contract}, nil
}

// NewArbFilterer creates a new log filterer instance of Arb, bound to a specific deployed contract.
func NewArbFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbFilterer, error) {
	contract, err := bindArb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbFilterer{contract: contract}, nil
}

// bindArb binds a generic wrapper to an already deployed contract.
func bindArb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.ArbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transact(opts, method, params...)
}

// EstimateDualDexTrade is a free data retrieval call binding the contract method 0x068e7ca1.
//
// Solidity: function estimateDualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) view returns(uint256)
func (_Arb *ArbCaller) EstimateDualDexTrade(opts *bind.CallOpts, _router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "estimateDualDexTrade", _router1, _router2, _token1, _token2, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateDualDexTrade is a free data retrieval call binding the contract method 0x068e7ca1.
//
// Solidity: function estimateDualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) view returns(uint256)
func (_Arb *ArbSession) EstimateDualDexTrade(_router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateDualDexTrade(&_Arb.CallOpts, _router1, _router2, _token1, _token2, _amount)
}

// EstimateDualDexTrade is a free data retrieval call binding the contract method 0x068e7ca1.
//
// Solidity: function estimateDualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) view returns(uint256)
func (_Arb *ArbCallerSession) EstimateDualDexTrade(_router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateDualDexTrade(&_Arb.CallOpts, _router1, _router2, _token1, _token2, _amount)
}

// EstimateTriDexTrade is a free data retrieval call binding the contract method 0x26a215ec.
//
// Solidity: function estimateTriDexTrade(address _router1, address _router2, address _router3, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbCaller) EstimateTriDexTrade(opts *bind.CallOpts, _router1 common.Address, _router2 common.Address, _router3 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "estimateTriDexTrade", _router1, _router2, _router3, _token1, _token2, _token3, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTriDexTrade is a free data retrieval call binding the contract method 0x26a215ec.
//
// Solidity: function estimateTriDexTrade(address _router1, address _router2, address _router3, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbSession) EstimateTriDexTrade(_router1 common.Address, _router2 common.Address, _router3 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateTriDexTrade(&_Arb.CallOpts, _router1, _router2, _router3, _token1, _token2, _token3, _amount)
}

// EstimateTriDexTrade is a free data retrieval call binding the contract method 0x26a215ec.
//
// Solidity: function estimateTriDexTrade(address _router1, address _router2, address _router3, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbCallerSession) EstimateTriDexTrade(_router1 common.Address, _router2 common.Address, _router3 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateTriDexTrade(&_Arb.CallOpts, _router1, _router2, _router3, _token1, _token2, _token3, _amount)
}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x43cf6f24.
//
// Solidity: function getAmountOutMin(address router, address _tokenIn, address _tokenOut, uint256 _amount) view returns(uint256)
func (_Arb *ArbCaller) GetAmountOutMin(opts *bind.CallOpts, router common.Address, _tokenIn common.Address, _tokenOut common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "getAmountOutMin", router, _tokenIn, _tokenOut, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x43cf6f24.
//
// Solidity: function getAmountOutMin(address router, address _tokenIn, address _tokenOut, uint256 _amount) view returns(uint256)
func (_Arb *ArbSession) GetAmountOutMin(router common.Address, _tokenIn common.Address, _tokenOut common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.GetAmountOutMin(&_Arb.CallOpts, router, _tokenIn, _tokenOut, _amount)
}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x43cf6f24.
//
// Solidity: function getAmountOutMin(address router, address _tokenIn, address _tokenOut, uint256 _amount) view returns(uint256)
func (_Arb *ArbCallerSession) GetAmountOutMin(router common.Address, _tokenIn common.Address, _tokenOut common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.GetAmountOutMin(&_Arb.CallOpts, router, _tokenIn, _tokenOut, _amount)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenContractAddress) view returns(uint256)
func (_Arb *ArbCaller) GetBalance(opts *bind.CallOpts, _tokenContractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "getBalance", _tokenContractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenContractAddress) view returns(uint256)
func (_Arb *ArbSession) GetBalance(_tokenContractAddress common.Address) (*big.Int, error) {
	return _Arb.Contract.GetBalance(&_Arb.CallOpts, _tokenContractAddress)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenContractAddress) view returns(uint256)
func (_Arb *ArbCallerSession) GetBalance(_tokenContractAddress common.Address) (*big.Int, error) {
	return _Arb.Contract.GetBalance(&_Arb.CallOpts, _tokenContractAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbSession) Owner() (common.Address, error) {
	return _Arb.Contract.Owner(&_Arb.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbCallerSession) Owner() (common.Address, error) {
	return _Arb.Contract.Owner(&_Arb.CallOpts)
}

// DualDexTrade is a paid mutator transaction binding the contract method 0x1aa51318.
//
// Solidity: function dualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) returns()
func (_Arb *ArbTransactor) DualDexTrade(opts *bind.TransactOpts, _router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "dualDexTrade", _router1, _router2, _token1, _token2, _amount)
}

// DualDexTrade is a paid mutator transaction binding the contract method 0x1aa51318.
//
// Solidity: function dualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) returns()
func (_Arb *ArbSession) DualDexTrade(_router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.DualDexTrade(&_Arb.TransactOpts, _router1, _router2, _token1, _token2, _amount)
}

// DualDexTrade is a paid mutator transaction binding the contract method 0x1aa51318.
//
// Solidity: function dualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) returns()
func (_Arb *ArbTransactorSession) DualDexTrade(_router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.DualDexTrade(&_Arb.TransactOpts, _router1, _router2, _token1, _token2, _amount)
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Arb *ArbTransactor) RecoverEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "recoverEth")
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Arb *ArbSession) RecoverEth() (*types.Transaction, error) {
	return _Arb.Contract.RecoverEth(&_Arb.TransactOpts)
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Arb *ArbTransactorSession) RecoverEth() (*types.Transaction, error) {
	return _Arb.Contract.RecoverEth(&_Arb.TransactOpts)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address tokenAddress) returns()
func (_Arb *ArbTransactor) RecoverTokens(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "recoverTokens", tokenAddress)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address tokenAddress) returns()
func (_Arb *ArbSession) RecoverTokens(tokenAddress common.Address) (*types.Transaction, error) {
	return _Arb.Contract.RecoverTokens(&_Arb.TransactOpts, tokenAddress)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address tokenAddress) returns()
func (_Arb *ArbTransactorSession) RecoverTokens(tokenAddress common.Address) (*types.Transaction, error) {
	return _Arb.Contract.RecoverTokens(&_Arb.TransactOpts, tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arb.Contract.RenounceOwnership(&_Arb.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arb.Contract.RenounceOwnership(&_Arb.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.TransferOwnership(&_Arb.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.TransferOwnership(&_Arb.TransactOpts, newOwner)
}

// ArbOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Arb contract.
type ArbOwnershipTransferredIterator struct {
	Event *ArbOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArbOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbOwnershipTransferred)
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
		it.Event = new(ArbOwnershipTransferred)
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
func (it *ArbOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbOwnershipTransferred represents a OwnershipTransferred event raised by the Arb contract.
type ArbOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArbOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArbOwnershipTransferredIterator{contract: _Arb.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbOwnershipTransferred)
				if err := _Arb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Arb *ArbFilterer) ParseOwnershipTransferred(log types.Log) (*ArbOwnershipTransferred, error) {
	event := new(ArbOwnershipTransferred)
	if err := _Arb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
