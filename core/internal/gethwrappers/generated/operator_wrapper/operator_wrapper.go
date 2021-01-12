// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operator_wrapper

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OperatorABI is the input ABI used to generate the binding from.
const OperatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"CancelOracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"specId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cancelExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dataVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"OracleResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EXPIRY_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunc\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"cancelOracleRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"forward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"fulfillOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"fulfillOracleRequest2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainlinkToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"isAuthorizedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"specId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataVersion\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"oracleRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setAuthorizedSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OperatorBin is the compiled bytecode used for deploying new contracts.
var OperatorBin = "0x60a06040526001600455600160055534801561001a57600080fd5b506040516126133803806126138339818101604052604081101561003d57600080fd5b508051602090910151600080546001600160a01b0319166001600160a01b03928316178155606083901b6001600160601b031916608052911690612567906100ac90398061066c52806107325280610d2a528061132452806114085280611d33528061226b52506125676000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80636fadcf72116100b2578063a4c0ed3611610081578063f3dfc2a911610066578063f3dfc2a9146105c3578063f3fef3a3146105fe578063fa00763a146106375761011b565b8063a4c0ed36146104c8578063f2fde38b146105905761011b565b80636fadcf721461042357806379ba5097146104b057806380e673da146104b85780638da5cb5b146104c05761011b565b806350188301116100ee57806350188301146102ca5780635705ae43146102d25780636ae0bc761461030b5780636ee4d553146103d25761011b565b8063165d35e11461012057806340429946146101515780634ab0d190146102295780634b602282146102b0575b600080fd5b61012861066a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b610227600480360361010081101561016857600080fd5b73ffffffffffffffffffffffffffffffffffffffff8235811692602081013592604082013592606083013516917fffffffff000000000000000000000000000000000000000000000000000000006080820135169160a08201359160c081013591810190610100810160e08201356401000000008111156101e857600080fd5b8201836020820111156101fa57600080fd5b8035906020019184600183028401116401000000008311171561021c57600080fd5b50909250905061068e565b005b61029c600480360360c081101561023f57600080fd5b5080359060208101359073ffffffffffffffffffffffffffffffffffffffff604082013516907fffffffff000000000000000000000000000000000000000000000000000000006060820135169060808101359060a00135610910565b604080519115158252519081900360200190f35b6102b8610c0f565b60408051918252519081900360200190f35b6102b8610c15565b610227600480360360408110156102e857600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610c2b565b61029c600480360360c081101561032157600080fd5b81359160208101359173ffffffffffffffffffffffffffffffffffffffff604083013516917fffffffff00000000000000000000000000000000000000000000000000000000606082013516916080820135919081019060c0810160a082013564010000000081111561039357600080fd5b8201836020820111156103a557600080fd5b803590602001918460018302840111640100000000831117156103c757600080fd5b509092509050610dea565b610227600480360360808110156103e857600080fd5b508035906020810135907fffffffff00000000000000000000000000000000000000000000000000000000604082013516906060013561117a565b6102276004803603604081101561043957600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516919081019060408101602082013564010000000081111561047157600080fd5b82018360208201111561048357600080fd5b803590602001918460018302840111640100000000831117156104a557600080fd5b50909250905061139e565b610227611591565b6102b8611693565b61012861169d565b610227600480360360608110156104de57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8235169160208101359181019060608101604082013564010000000081111561051b57600080fd5b82018360208201111561052d57600080fd5b8035906020019184600183028401116401000000008311171561054f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506116b9945050505050565b610227600480360360208110156105a657600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166119d5565b610227600480360360408110156105d957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001351515611b56565b6102276004803603604081101561061457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135611c32565b61029c6004803603602081101561064d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16611df9565b7f000000000000000000000000000000000000000000000000000000000000000090565b61069661066a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461072f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4d75737420757365204c494e4b20746f6b656e00000000000000000000000000604482015290519081900360640190fd5b857f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156107eb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f742063616c6c6261636b20746f204c494e4b000000000000000000604482015290519081900360640190fd5b6000806107fc8c8c8b8b8b8b611e24565b91509150897fd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c658d848e8d8d878d8d8d604051808a73ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018881526020018773ffffffffffffffffffffffffffffffffffffffff168152602001867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001858152602001848152602001806020018281038252848482818152602001925080828437600083820152604051601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169092018290039c50909a5050505050505050505050a2505050505050505050505050565b3360009081526003602052604081205460ff16610978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612508602a913960400191505060405180910390fd5b600087815260026020526040902054879060081b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016610a1957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4d757374206861766520612076616c6964207265717565737449640000000000604482015290519081900360640190fd5b610a288888888888600161200d565b60405188907f9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a6490600090a262061a805a1015610ac557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4d7573742070726f7669646520636f6e73756d657220656e6f75676820676173604482015290519081900360640190fd5b60408051602481018a9052604480820186905282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000089161781529151815160009373ffffffffffffffffffffffffffffffffffffffff8b169392918291908083835b60208310610b9857805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610b5b565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610bfa576040519150601f19603f3d011682016040523d82523d6000602084013e610bff565b606091505b50909a9950505050505050505050565b61012c81565b600554600090610c269060016121ab565b905090565b60005473ffffffffffffffffffffffffffffffffffffffff163314610cb157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015290519081900360640190fd5b80610cba612222565b811115610d2857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4e6f7420656e6f756768207265636f76657261626c652066756e647300000000604482015290519081900360640190fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610db957600080fd5b505af1158015610dcd573d6000803e3d6000fd5b505050506040513d6020811015610de357600080fd5b5050505050565b3360009081526003602052604081205460ff16610e52576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612508602a913960400191505060405180910390fd5b600088815260026020526040902054889060081b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016610ef357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4d757374206861766520612076616c6964207265717565737449640000000000604482015290519081900360640190fd5b8884848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506020810151828114610f9d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f466972737420776f7264206d7573742062652072657175657374496400000000604482015290519081900360640190fd5b610fac8c8c8c8c8c600261200d565b6040518c907f9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a6490600090a262061a805a101561104957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4d7573742070726f7669646520636f6e73756d657220656e6f75676820676173604482015290519081900360640190fd5b60008a73ffffffffffffffffffffffffffffffffffffffff168a898960405160200180847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526004018383808284378083019250505093505050506040516020818303038152906040526040518082805190602001908083835b602083106110ff57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016110c2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611161576040519150601f19603f3d011682016040523d82523d6000602084013e611166565b606091505b50909e9d5050505050505050505050505050565b60006111888433858561232a565b60008681526002602052604090205490915060081b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009081169082161461123057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f506172616d7320646f206e6f74206d6174636820726571756573742049440000604482015290519081900360640190fd5b4282111561129f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f52657175657374206973206e6f74206578706972656400000000000000000000604482015290519081900360640190fd5b6000858152600260205260408082208290555186917fa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e9391a2604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815233600482015260248101869052905173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163a9059cbb9160448083019260209291908290030181600087803b15801561136c57600080fd5b505af1158015611380573d6000803e3d6000fd5b505050506040513d602081101561139657600080fd5b5051610de357fe5b3360009081526003602052604090205460ff16611406576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612508602a913960400191505060405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156114ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806124a16032913960400191505060405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff168383604051808383808284376040519201945060009350909150508083038183865af19150503d8060008114611516576040519150601f19603f3d011682016040523d82523d6000602084013e61151b565b606091505b505090508061158b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f466f727761726465642063616c6c206661696c65642e00000000000000000000604482015290519081900360640190fd5b50505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461161757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015290519081900360640190fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6000610c26612222565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6116c161066a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461175a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4d75737420757365204c494e4b20746f6b656e00000000000000000000000000604482015290519081900360640190fd5b80518190604411156117cd57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e76616c69642072657175657374206c656e67746800000000000000000000604482015290519081900360640190fd5b602082015182907fffffffff0000000000000000000000000000000000000000000000000000000081167f40429946000000000000000000000000000000000000000000000000000000001461188457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4d757374207573652077686974656c69737465642066756e6374696f6e730000604482015290519081900360640190fd5b85602485015284604485015260003073ffffffffffffffffffffffffffffffffffffffff16856040518082805190602001908083835b602083106118f757805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016118ba565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114611957576040519150601f19603f3d011682016040523d82523d6000602084013e61195c565b606091505b50509050806119cc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f556e61626c6520746f2063726561746520726571756573740000000000000000604482015290519081900360640190fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611a5b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8116331415611ae057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015290519081900360640190fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005473ffffffffffffffffffffffffffffffffffffffff163314611bdc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260036020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff163314611cb857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015290519081900360640190fd5b80611cc48160016123af565b6005541015611d1e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260358152602001806124d36035913960400191505060405180910390fd5b600554611d2b90836121ab565b6005819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611dc257600080fd5b505af1158015611dd6573d6000803e3d6000fd5b505050506040513d6020811015611dec57600080fd5b5051611df457fe5b505050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205460ff1690565b60408051606088901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208083019190915260348083018690528351808403909101815260549092018352815191810191909120600081815260029092529181205460081b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001615611f1957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f4d75737420757365206120756e69717565204944000000000000000000000000604482015290519081900360640190fd5b611f254261012c6123af565b90506000611f358888888561232a565b905060405180604001604052808260ff19168152602001611f558661242a565b60ff90811690915260008581526002602090815260409091208351815494909201519092167f01000000000000000000000000000000000000000000000000000000000000000260089190911c7fff00000000000000000000000000000000000000000000000000000000000000909316929092177effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff16919091179055600454611ffe90896123af565b60045550965096945050505050565b600061201b8686868661232a565b60008881526002602052604090205490915060081b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00908116908216146120c357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f506172616d7320646f206e6f74206d6174636820726571756573742049440000604482015290519081900360640190fd5b6120cc8261242a565b60008881526002602052604090205460ff9182167f0100000000000000000000000000000000000000000000000000000000000000909104909116111561217457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f446174612076657273696f6e73206d757374206d617463680000000000000000604482015290519081900360640190fd5b60045461218190876121ab565b60045560055461219190876123af565b600555505050600093845250506002602052506040812055565b60008282111561221c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b60008061223b60016004546121ab90919063ffffffff16565b9050600061225560016005546121ab90919063ffffffff16565b9050600061226383836123af565b9050612322817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156122f057600080fd5b505afa158015612304573d6000803e3d6000fd5b505050506040513d602081101561231a57600080fd5b5051906121ab565b935050505090565b6040805160208082019690965260609490941b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016848201527fffffffff000000000000000000000000000000000000000000000000000000009290921660548401526058808401919091528151808403909101815260789092019052805191012090565b60008282018381101561242357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000610100821061249c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f6e756d62657220746f6f2062696720746f206361737400000000000000000000604482015290519081900360640190fd5b509056fe43616e6e6f74207573652023666f727761726420746f2073656e64206d6573736167657320746f204c696e6b20746f6b656e416d6f756e74207265717565737465642069732067726561746572207468616e20776974686472617761626c652062616c616e63654e6f7420616e20617574686f72697a6564206e6f646520746f2066756c66696c6c207265717565737473a264697066735822beefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeef64736f6c6343decafe0033"

// DeployOperator deploys a new Ethereum contract, binding an instance of Operator to it.
func DeployOperator(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, owner common.Address) (common.Address, *types.Transaction, *Operator, error) {
	parsed, err := abi.JSON(strings.NewReader(OperatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OperatorBin), backend, link, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Operator{OperatorCaller: OperatorCaller{contract: contract}, OperatorTransactor: OperatorTransactor{contract: contract}, OperatorFilterer: OperatorFilterer{contract: contract}}, nil
}

// Operator is an auto generated Go binding around an Ethereum contract.
type Operator struct {
	OperatorCaller     // Read-only binding to the contract
	OperatorTransactor // Write-only binding to the contract
	OperatorFilterer   // Log filterer for contract events
}

// OperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorSession struct {
	Contract     *Operator         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorCallerSession struct {
	Contract *OperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorTransactorSession struct {
	Contract     *OperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRaw struct {
	Contract *Operator // Generic contract binding to access the raw methods on
}

// OperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorCallerRaw struct {
	Contract *OperatorCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorTransactorRaw struct {
	Contract *OperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperator creates a new instance of Operator, bound to a specific deployed contract.
func NewOperator(address common.Address, backend bind.ContractBackend) (*Operator, error) {
	contract, err := bindOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Operator{OperatorCaller: OperatorCaller{contract: contract}, OperatorTransactor: OperatorTransactor{contract: contract}, OperatorFilterer: OperatorFilterer{contract: contract}}, nil
}

// NewOperatorCaller creates a new read-only instance of Operator, bound to a specific deployed contract.
func NewOperatorCaller(address common.Address, caller bind.ContractCaller) (*OperatorCaller, error) {
	contract, err := bindOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorCaller{contract: contract}, nil
}

// NewOperatorTransactor creates a new write-only instance of Operator, bound to a specific deployed contract.
func NewOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorTransactor, error) {
	contract, err := bindOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTransactor{contract: contract}, nil
}

// NewOperatorFilterer creates a new log filterer instance of Operator, bound to a specific deployed contract.
func NewOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorFilterer, error) {
	contract, err := bindOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorFilterer{contract: contract}, nil
}

// bindOperator binds a generic wrapper to an already deployed contract.
func bindOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OperatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.OperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transact(opts, method, params...)
}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_Operator *OperatorCaller) EXPIRYTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "EXPIRY_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_Operator *OperatorSession) EXPIRYTIME() (*big.Int, error) {
	return _Operator.Contract.EXPIRYTIME(&_Operator.CallOpts)
}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_Operator *OperatorCallerSession) EXPIRYTIME() (*big.Int, error) {
	return _Operator.Contract.EXPIRYTIME(&_Operator.CallOpts)
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_Operator *OperatorCaller) GetChainlinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "getChainlinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_Operator *OperatorSession) GetChainlinkToken() (common.Address, error) {
	return _Operator.Contract.GetChainlinkToken(&_Operator.CallOpts)
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_Operator *OperatorCallerSession) GetChainlinkToken() (common.Address, error) {
	return _Operator.Contract.GetChainlinkToken(&_Operator.CallOpts)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address node) view returns(bool)
func (_Operator *OperatorCaller) IsAuthorizedSender(opts *bind.CallOpts, node common.Address) (bool, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "isAuthorizedSender", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address node) view returns(bool)
func (_Operator *OperatorSession) IsAuthorizedSender(node common.Address) (bool, error) {
	return _Operator.Contract.IsAuthorizedSender(&_Operator.CallOpts, node)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address node) view returns(bool)
func (_Operator *OperatorCallerSession) IsAuthorizedSender(node common.Address) (bool, error) {
	return _Operator.Contract.IsAuthorizedSender(&_Operator.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorSession) Owner() (common.Address, error) {
	return _Operator.Contract.Owner(&_Operator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorCallerSession) Owner() (common.Address, error) {
	return _Operator.Contract.Owner(&_Operator.CallOpts)
}

// Recoverable is a free data retrieval call binding the contract method 0x80e673da.
//
// Solidity: function recoverable() view returns(uint256)
func (_Operator *OperatorCaller) Recoverable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "recoverable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Recoverable is a free data retrieval call binding the contract method 0x80e673da.
//
// Solidity: function recoverable() view returns(uint256)
func (_Operator *OperatorSession) Recoverable() (*big.Int, error) {
	return _Operator.Contract.Recoverable(&_Operator.CallOpts)
}

// Recoverable is a free data retrieval call binding the contract method 0x80e673da.
//
// Solidity: function recoverable() view returns(uint256)
func (_Operator *OperatorCallerSession) Recoverable() (*big.Int, error) {
	return _Operator.Contract.Recoverable(&_Operator.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() view returns(uint256)
func (_Operator *OperatorCaller) Withdrawable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "withdrawable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() view returns(uint256)
func (_Operator *OperatorSession) Withdrawable() (*big.Int, error) {
	return _Operator.Contract.Withdrawable(&_Operator.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() view returns(uint256)
func (_Operator *OperatorCallerSession) Withdrawable() (*big.Int, error) {
	return _Operator.Contract.Withdrawable(&_Operator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Operator *OperatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Operator *OperatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Operator.Contract.AcceptOwnership(&_Operator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Operator *OperatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Operator.Contract.AcceptOwnership(&_Operator.TransactOpts)
}

// CancelOracleRequest is a paid mutator transaction binding the contract method 0x6ee4d553.
//
// Solidity: function cancelOracleRequest(bytes32 requestId, uint256 payment, bytes4 callbackFunc, uint256 expiration) returns()
func (_Operator *OperatorTransactor) CancelOracleRequest(opts *bind.TransactOpts, requestId [32]byte, payment *big.Int, callbackFunc [4]byte, expiration *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "cancelOracleRequest", requestId, payment, callbackFunc, expiration)
}

// CancelOracleRequest is a paid mutator transaction binding the contract method 0x6ee4d553.
//
// Solidity: function cancelOracleRequest(bytes32 requestId, uint256 payment, bytes4 callbackFunc, uint256 expiration) returns()
func (_Operator *OperatorSession) CancelOracleRequest(requestId [32]byte, payment *big.Int, callbackFunc [4]byte, expiration *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.CancelOracleRequest(&_Operator.TransactOpts, requestId, payment, callbackFunc, expiration)
}

// CancelOracleRequest is a paid mutator transaction binding the contract method 0x6ee4d553.
//
// Solidity: function cancelOracleRequest(bytes32 requestId, uint256 payment, bytes4 callbackFunc, uint256 expiration) returns()
func (_Operator *OperatorTransactorSession) CancelOracleRequest(requestId [32]byte, payment *big.Int, callbackFunc [4]byte, expiration *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.CancelOracleRequest(&_Operator.TransactOpts, requestId, payment, callbackFunc, expiration)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address to, bytes data) returns()
func (_Operator *OperatorTransactor) Forward(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "forward", to, data)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address to, bytes data) returns()
func (_Operator *OperatorSession) Forward(to common.Address, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.Forward(&_Operator.TransactOpts, to, data)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address to, bytes data) returns()
func (_Operator *OperatorTransactorSession) Forward(to common.Address, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.Forward(&_Operator.TransactOpts, to, data)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0x4ab0d190.
//
// Solidity: function fulfillOracleRequest(bytes32 requestId, uint256 payment, address callbackAddress, bytes4 callbackFunctionId, uint256 expiration, bytes32 data) returns(bool)
func (_Operator *OperatorTransactor) FulfillOracleRequest(opts *bind.TransactOpts, requestId [32]byte, payment *big.Int, callbackAddress common.Address, callbackFunctionId [4]byte, expiration *big.Int, data [32]byte) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "fulfillOracleRequest", requestId, payment, callbackAddress, callbackFunctionId, expiration, data)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0x4ab0d190.
//
// Solidity: function fulfillOracleRequest(bytes32 requestId, uint256 payment, address callbackAddress, bytes4 callbackFunctionId, uint256 expiration, bytes32 data) returns(bool)
func (_Operator *OperatorSession) FulfillOracleRequest(requestId [32]byte, payment *big.Int, callbackAddress common.Address, callbackFunctionId [4]byte, expiration *big.Int, data [32]byte) (*types.Transaction, error) {
	return _Operator.Contract.FulfillOracleRequest(&_Operator.TransactOpts, requestId, payment, callbackAddress, callbackFunctionId, expiration, data)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0x4ab0d190.
//
// Solidity: function fulfillOracleRequest(bytes32 requestId, uint256 payment, address callbackAddress, bytes4 callbackFunctionId, uint256 expiration, bytes32 data) returns(bool)
func (_Operator *OperatorTransactorSession) FulfillOracleRequest(requestId [32]byte, payment *big.Int, callbackAddress common.Address, callbackFunctionId [4]byte, expiration *big.Int, data [32]byte) (*types.Transaction, error) {
	return _Operator.Contract.FulfillOracleRequest(&_Operator.TransactOpts, requestId, payment, callbackAddress, callbackFunctionId, expiration, data)
}

// FulfillOracleRequest2 is a paid mutator transaction binding the contract method 0x6ae0bc76.
//
// Solidity: function fulfillOracleRequest2(bytes32 requestId, uint256 payment, address callbackAddress, bytes4 callbackFunctionId, uint256 expiration, bytes data) returns(bool)
func (_Operator *OperatorTransactor) FulfillOracleRequest2(opts *bind.TransactOpts, requestId [32]byte, payment *big.Int, callbackAddress common.Address, callbackFunctionId [4]byte, expiration *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "fulfillOracleRequest2", requestId, payment, callbackAddress, callbackFunctionId, expiration, data)
}

// FulfillOracleRequest2 is a paid mutator transaction binding the contract method 0x6ae0bc76.
//
// Solidity: function fulfillOracleRequest2(bytes32 requestId, uint256 payment, address callbackAddress, bytes4 callbackFunctionId, uint256 expiration, bytes data) returns(bool)
func (_Operator *OperatorSession) FulfillOracleRequest2(requestId [32]byte, payment *big.Int, callbackAddress common.Address, callbackFunctionId [4]byte, expiration *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.FulfillOracleRequest2(&_Operator.TransactOpts, requestId, payment, callbackAddress, callbackFunctionId, expiration, data)
}

// FulfillOracleRequest2 is a paid mutator transaction binding the contract method 0x6ae0bc76.
//
// Solidity: function fulfillOracleRequest2(bytes32 requestId, uint256 payment, address callbackAddress, bytes4 callbackFunctionId, uint256 expiration, bytes data) returns(bool)
func (_Operator *OperatorTransactorSession) FulfillOracleRequest2(requestId [32]byte, payment *big.Int, callbackAddress common.Address, callbackFunctionId [4]byte, expiration *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.FulfillOracleRequest2(&_Operator.TransactOpts, requestId, payment, callbackAddress, callbackFunctionId, expiration, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_Operator *OperatorTransactor) OnTokenTransfer(opts *bind.TransactOpts, sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "onTokenTransfer", sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_Operator *OperatorSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.OnTokenTransfer(&_Operator.TransactOpts, sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_Operator *OperatorTransactorSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.OnTokenTransfer(&_Operator.TransactOpts, sender, amount, data)
}

// OracleRequest is a paid mutator transaction binding the contract method 0x40429946.
//
// Solidity: function oracleRequest(address sender, uint256 payment, bytes32 specId, address callbackAddress, bytes4 callbackFunctionId, uint256 nonce, uint256 dataVersion, bytes data) returns()
func (_Operator *OperatorTransactor) OracleRequest(opts *bind.TransactOpts, sender common.Address, payment *big.Int, specId [32]byte, callbackAddress common.Address, callbackFunctionId [4]byte, nonce *big.Int, dataVersion *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "oracleRequest", sender, payment, specId, callbackAddress, callbackFunctionId, nonce, dataVersion, data)
}

// OracleRequest is a paid mutator transaction binding the contract method 0x40429946.
//
// Solidity: function oracleRequest(address sender, uint256 payment, bytes32 specId, address callbackAddress, bytes4 callbackFunctionId, uint256 nonce, uint256 dataVersion, bytes data) returns()
func (_Operator *OperatorSession) OracleRequest(sender common.Address, payment *big.Int, specId [32]byte, callbackAddress common.Address, callbackFunctionId [4]byte, nonce *big.Int, dataVersion *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.OracleRequest(&_Operator.TransactOpts, sender, payment, specId, callbackAddress, callbackFunctionId, nonce, dataVersion, data)
}

// OracleRequest is a paid mutator transaction binding the contract method 0x40429946.
//
// Solidity: function oracleRequest(address sender, uint256 payment, bytes32 specId, address callbackAddress, bytes4 callbackFunctionId, uint256 nonce, uint256 dataVersion, bytes data) returns()
func (_Operator *OperatorTransactorSession) OracleRequest(sender common.Address, payment *big.Int, specId [32]byte, callbackAddress common.Address, callbackFunctionId [4]byte, nonce *big.Int, dataVersion *big.Int, data []byte) (*types.Transaction, error) {
	return _Operator.Contract.OracleRequest(&_Operator.TransactOpts, sender, payment, specId, callbackAddress, callbackFunctionId, nonce, dataVersion, data)
}

// Recover is a paid mutator transaction binding the contract method 0x5705ae43.
//
// Solidity: function recover(address recipient, uint256 amount) returns()
func (_Operator *OperatorTransactor) Recover(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "recover", recipient, amount)
}

// Recover is a paid mutator transaction binding the contract method 0x5705ae43.
//
// Solidity: function recover(address recipient, uint256 amount) returns()
func (_Operator *OperatorSession) Recover(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.Recover(&_Operator.TransactOpts, recipient, amount)
}

// Recover is a paid mutator transaction binding the contract method 0x5705ae43.
//
// Solidity: function recover(address recipient, uint256 amount) returns()
func (_Operator *OperatorTransactorSession) Recover(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.Recover(&_Operator.TransactOpts, recipient, amount)
}

// SetAuthorizedSender is a paid mutator transaction binding the contract method 0xf3dfc2a9.
//
// Solidity: function setAuthorizedSender(address node, bool allowed) returns()
func (_Operator *OperatorTransactor) SetAuthorizedSender(opts *bind.TransactOpts, node common.Address, allowed bool) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "setAuthorizedSender", node, allowed)
}

// SetAuthorizedSender is a paid mutator transaction binding the contract method 0xf3dfc2a9.
//
// Solidity: function setAuthorizedSender(address node, bool allowed) returns()
func (_Operator *OperatorSession) SetAuthorizedSender(node common.Address, allowed bool) (*types.Transaction, error) {
	return _Operator.Contract.SetAuthorizedSender(&_Operator.TransactOpts, node, allowed)
}

// SetAuthorizedSender is a paid mutator transaction binding the contract method 0xf3dfc2a9.
//
// Solidity: function setAuthorizedSender(address node, bool allowed) returns()
func (_Operator *OperatorTransactorSession) SetAuthorizedSender(node common.Address, allowed bool) (*types.Transaction, error) {
	return _Operator.Contract.SetAuthorizedSender(&_Operator.TransactOpts, node, allowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Operator *OperatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Operator *OperatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Operator.Contract.TransferOwnership(&_Operator.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Operator *OperatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Operator.Contract.TransferOwnership(&_Operator.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Operator *OperatorTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "withdraw", recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Operator *OperatorSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.Withdraw(&_Operator.TransactOpts, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Operator *OperatorTransactorSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Operator.Contract.Withdraw(&_Operator.TransactOpts, recipient, amount)
}

// OperatorCancelOracleRequestIterator is returned from FilterCancelOracleRequest and is used to iterate over the raw logs and unpacked data for CancelOracleRequest events raised by the Operator contract.
type OperatorCancelOracleRequestIterator struct {
	Event *OperatorCancelOracleRequest // Event containing the contract specifics and raw log

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
func (it *OperatorCancelOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorCancelOracleRequest)
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
		it.Event = new(OperatorCancelOracleRequest)
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
func (it *OperatorCancelOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorCancelOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorCancelOracleRequest represents a CancelOracleRequest event raised by the Operator contract.
type OperatorCancelOracleRequest struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCancelOracleRequest is a free log retrieval operation binding the contract event 0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93.
//
// Solidity: event CancelOracleRequest(bytes32 indexed requestId)
func (_Operator *OperatorFilterer) FilterCancelOracleRequest(opts *bind.FilterOpts, requestId [][32]byte) (*OperatorCancelOracleRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "CancelOracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OperatorCancelOracleRequestIterator{contract: _Operator.contract, event: "CancelOracleRequest", logs: logs, sub: sub}, nil
}

// WatchCancelOracleRequest is a free log subscription operation binding the contract event 0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93.
//
// Solidity: event CancelOracleRequest(bytes32 indexed requestId)
func (_Operator *OperatorFilterer) WatchCancelOracleRequest(opts *bind.WatchOpts, sink chan<- *OperatorCancelOracleRequest, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "CancelOracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorCancelOracleRequest)
				if err := _Operator.contract.UnpackLog(event, "CancelOracleRequest", log); err != nil {
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

// ParseCancelOracleRequest is a log parse operation binding the contract event 0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93.
//
// Solidity: event CancelOracleRequest(bytes32 indexed requestId)
func (_Operator *OperatorFilterer) ParseCancelOracleRequest(log types.Log) (*OperatorCancelOracleRequest, error) {
	event := new(OperatorCancelOracleRequest)
	if err := _Operator.contract.UnpackLog(event, "CancelOracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOracleRequestIterator is returned from FilterOracleRequest and is used to iterate over the raw logs and unpacked data for OracleRequest events raised by the Operator contract.
type OperatorOracleRequestIterator struct {
	Event *OperatorOracleRequest // Event containing the contract specifics and raw log

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
func (it *OperatorOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOracleRequest)
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
		it.Event = new(OperatorOracleRequest)
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
func (it *OperatorOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOracleRequest represents a OracleRequest event raised by the Operator contract.
type OperatorOracleRequest struct {
	SpecId             [32]byte
	Requester          common.Address
	RequestId          [32]byte
	Payment            *big.Int
	CallbackAddr       common.Address
	CallbackFunctionId [4]byte
	CancelExpiration   *big.Int
	DataVersion        *big.Int
	Data               []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOracleRequest is a free log retrieval operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_Operator *OperatorFilterer) FilterOracleRequest(opts *bind.FilterOpts, specId [][32]byte) (*OperatorOracleRequestIterator, error) {

	var specIdRule []interface{}
	for _, specIdItem := range specId {
		specIdRule = append(specIdRule, specIdItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OracleRequest", specIdRule)
	if err != nil {
		return nil, err
	}
	return &OperatorOracleRequestIterator{contract: _Operator.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

// WatchOracleRequest is a free log subscription operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_Operator *OperatorFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *OperatorOracleRequest, specId [][32]byte) (event.Subscription, error) {

	var specIdRule []interface{}
	for _, specIdItem := range specId {
		specIdRule = append(specIdRule, specIdItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OracleRequest", specIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOracleRequest)
				if err := _Operator.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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

// ParseOracleRequest is a log parse operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_Operator *OperatorFilterer) ParseOracleRequest(log types.Log) (*OperatorOracleRequest, error) {
	event := new(OperatorOracleRequest)
	if err := _Operator.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOracleResponseIterator is returned from FilterOracleResponse and is used to iterate over the raw logs and unpacked data for OracleResponse events raised by the Operator contract.
type OperatorOracleResponseIterator struct {
	Event *OperatorOracleResponse // Event containing the contract specifics and raw log

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
func (it *OperatorOracleResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOracleResponse)
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
		it.Event = new(OperatorOracleResponse)
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
func (it *OperatorOracleResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOracleResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOracleResponse represents a OracleResponse event raised by the Operator contract.
type OperatorOracleResponse struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleResponse is a free log retrieval operation binding the contract event 0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64.
//
// Solidity: event OracleResponse(bytes32 indexed requestId)
func (_Operator *OperatorFilterer) FilterOracleResponse(opts *bind.FilterOpts, requestId [][32]byte) (*OperatorOracleResponseIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OracleResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OperatorOracleResponseIterator{contract: _Operator.contract, event: "OracleResponse", logs: logs, sub: sub}, nil
}

// WatchOracleResponse is a free log subscription operation binding the contract event 0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64.
//
// Solidity: event OracleResponse(bytes32 indexed requestId)
func (_Operator *OperatorFilterer) WatchOracleResponse(opts *bind.WatchOpts, sink chan<- *OperatorOracleResponse, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OracleResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOracleResponse)
				if err := _Operator.contract.UnpackLog(event, "OracleResponse", log); err != nil {
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

// ParseOracleResponse is a log parse operation binding the contract event 0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64.
//
// Solidity: event OracleResponse(bytes32 indexed requestId)
func (_Operator *OperatorFilterer) ParseOracleResponse(log types.Log) (*OperatorOracleResponse, error) {
	event := new(OperatorOracleResponse)
	if err := _Operator.contract.UnpackLog(event, "OracleResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the Operator contract.
type OperatorOwnershipTransferRequestedIterator struct {
	Event *OperatorOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *OperatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOwnershipTransferRequested)
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
		it.Event = new(OperatorOwnershipTransferRequested)
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
func (it *OperatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the Operator contract.
type OperatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Operator *OperatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OperatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OperatorOwnershipTransferRequestedIterator{contract: _Operator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Operator *OperatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OperatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOwnershipTransferRequested)
				if err := _Operator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Operator *OperatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*OperatorOwnershipTransferRequested, error) {
	event := new(OperatorOwnershipTransferRequested)
	if err := _Operator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Operator contract.
type OperatorOwnershipTransferredIterator struct {
	Event *OperatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OperatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOwnershipTransferred)
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
		it.Event = new(OperatorOwnershipTransferred)
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
func (it *OperatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOwnershipTransferred represents a OwnershipTransferred event raised by the Operator contract.
type OperatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Operator *OperatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OperatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OperatorOwnershipTransferredIterator{contract: _Operator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Operator *OperatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OperatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOwnershipTransferred)
				if err := _Operator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Operator *OperatorFilterer) ParseOwnershipTransferred(log types.Log) (*OperatorOwnershipTransferred, error) {
	event := new(OperatorOwnershipTransferred)
	if err := _Operator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
