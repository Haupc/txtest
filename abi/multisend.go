// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// MultisendMetaData contains all meta data concerning the Multisend contract.
var MultisendMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escapeHatchCaller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"}],\"name\":\"multiTransferTightlyPacked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_addresses\",\"type\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiERC20Transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_dac\",\"type\":\"address\"}],\"name\":\"removeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwnerCandidate\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isTokenEscapable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"escapeHatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"}],\"name\":\"multiCallTightlyPacked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwnerCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newEscapeHatchCaller\",\"type\":\"address\"}],\"name\":\"changeHatchEscapeCaller\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escapeHatchDestination\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"}],\"name\":\"multiERC20TransferTightlyPacked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"MultiTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"MultiCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"MultiERC20Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"EscapeHatchBlackistedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EscapeHatchCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"by\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OwnershipRemoved\",\"type\":\"event\"}]",
	Bin: "0x608060405273839395e20bbb182fa440d08f850e6c7a8f6f0780600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550738ff920020c8ad673661c8117f2855c384758c572600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503480156100ba57600080fd5b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506121d4806101d46000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631476e40f146100f65780631e89d545146101aa5780631f6eb6e71461025e5780632a17e397146102b55780632af4c31e1461032657806335a2172814610369578063666a342714610432578063710bf3221461047557806379ba5097146104b8578063892db057146104cf5780638da5cb5b1461052a578063a142d60814610581578063ac66777f146105c4578063d091b55014610635578063d836fbe81461068c578063f5b61230146106cf578063f8f1d92714610726575b600080fd5b61019060048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192905050506107ac565b604051808215151515815260200191505060405180910390f35b61024460048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192905050506108f1565b604051808215151515815260200191505060405180910390f35b34801561026a57600080fd5b50610273610a36565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61030c60048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610a5c565b604051808215151515815260200191505060405180910390f35b34801561033257600080fd5b50610367600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bb1565b005b34801561037557600080fd5b50610430600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610e2b565b005b34801561043e57600080fd5b50610473600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f6b565b005b34801561048157600080fd5b506104b6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611170565b005b3480156104c457600080fd5b506104cd6112f4565b005b3480156104db57600080fd5b50610510600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611501565b604051808215151515815260200191505060405180910390f35b34801561053657600080fd5b5061053f611558565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561058d57600080fd5b506105c2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061157d565b005b61061b60048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050611b36565b604051808215151515815260200191505060405180910390f35b34801561064157600080fd5b5061064a611c8b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561069857600080fd5b506106cd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611cb1565b005b3480156106db57600080fd5b506106e4611e11565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561073257600080fd5b506107aa600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050611e37565b005b6000806000349150600090505b84518110156108db576107fa85828151811015156107d357fe5b9060200190602002015185838151811015156107eb57fe5b90602001906020020151611f75565b61081b82858381518110151561080c57fe5b90602001906020020151611fd6565b9150343373ffffffffffffffffffffffffffffffffffffffff167fa5c03fd81f783c38e4b274cd539e517bacfc824ee513ec4dad385269fa9684d7878481518110151561086457fe5b90602001906020020151878581518110151561087c57fe5b90602001906020020151604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a380806001019150506107b9565b6108e53383611fef565b60019250505092915050565b6000806000349150600090505b8451811015610a205761093f858281518110151561091857fe5b90602001906020020151858381518110151561093057fe5b90602001906020020151611fef565b61096082858381518110151561095157fe5b90602001906020020151611fd6565b9150343373ffffffffffffffffffffffffffffffffffffffff167f319e0008dcdeba1f31169497fb0f35d31b2b2f481c714d1f50640e86ac6c3bb087848151811015156109a957fe5b9060200190602002015187858151811015156109c157fe5b90602001906020020151604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a380806001019150506108fe565b610a2a3383611fef565b60019250505092915050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000806000349350600092505b8551831015610b9a5760608684815181101515610a8557fe5b90602001906020020151600019169060020a90046001900491508583815181101515610aad57fe5b90602001906020020151600190046bffffffffffffffffffffffff169050610afe828785815181101515610add57fe5b90602001906020020151600190046bffffffffffffffffffffffff16611fef565b610b088482611fd6565b9350343373ffffffffffffffffffffffffffffffffffffffff167f319e0008dcdeba1f31169497fb0f35d31b2b2f481c714d1f50640e86ac6c3bb08484604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a38280600101935050610a6c565b610ba43385611fef565b6001945050505050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c77576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6572725f6f776e65644e6f744f776e657200000000000000000000000000000081525060200191505060405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff1614151515610d06576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6572725f6f776e6564496e76616c69644164647265737300000000000000000081525060200191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008090505b8251811015610f6557610e73848483815181101515610e4c57fe5b906020019060200201518484815181101515610e6457fe5b90602001906020020151612060565b343373ffffffffffffffffffffffffffffffffffffffff167ffc01e439ca3c7015e18b8adea39e270034ba32f41fc788a4ff659842f0f37a938584815181101515610eba57fe5b906020019060200201518585815181101515610ed257fe5b9060200190602002015188604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a38080600101915050610e31565b50505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561102f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6572725f6f776e65644e6f744f776e657200000000000000000000000000000081525060200191505060405180910390fd5b610dac8173ffffffffffffffffffffffffffffffffffffffff161415156110be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6572725f6f776e6564496e76616c69644461630000000000000000000000000081525060200191505060405180910390fd5b60008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f94e8b32e01b9eedfddd778ffbd051a7718cdc14781702884561162dca6f74dbb60405160405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611234576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6572725f6f776e65644e6f744f776e657200000000000000000000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f13a4b3bc0d5234dd3d87c9f1557d8faefa37986da62c36ba49309e2fb2c9aec460405160405180910390a350565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156113bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6572725f6f776e65644e6f7443616e646964617465000000000000000000000081525060200191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16159050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061162857506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561169c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f6572725f657363617061626c65496e76616c696443616c6c657200000000000081525060200191505060405180910390fd5b60001515600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515611764576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f6572725f657363617061626c65426c61636b6c6973746564546f6b656e00000081525060200191505060405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff161415611876573073ffffffffffffffffffffffffffffffffffffffff16319150600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015611805573d6000803e3d6000fd5b507fa50dde912fa22ea0d215a0236093ac45b4d55d6ef0c604c319f900029c5d10f28383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1611b31565b8290508073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561191457600080fd5b505af1158015611928573d6000803e3d6000fd5b505050506040513d602081101561193e57600080fd5b810190808051906020019092919050505091508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611a1657600080fd5b505af1158015611a2a573d6000803e3d6000fd5b505050506040513d6020811015611a4057600080fd5b81019080805190602001909291905050501515611ac5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6572725f657363617061626c655472616e73666572000000000000000000000081525060200191505060405180910390fd5b7fa50dde912fa22ea0d215a0236093ac45b4d55d6ef0c604c319f900029c5d10f28383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15b505050565b6000806000806000349350600092505b8551831015611c745760608684815181101515611b5f57fe5b90602001906020020151600019169060020a90046001900491508583815181101515611b8757fe5b90602001906020020151600190046bffffffffffffffffffffffff169050611baf8282611f75565b611be2848785815181101515611bc157fe5b90602001906020020151600190046bffffffffffffffffffffffff16611fd6565b9350343373ffffffffffffffffffffffffffffffffffffffff167fa5c03fd81f783c38e4b274cd539e517bacfc824ee513ec4dad385269fa9684d78484604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a38280600101935050611b46565b611c7e3385611fef565b6001945050505050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480611d5957506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611dcd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f6572725f657363617061626c65496e76616c696443616c6c657200000000000081525060200191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060008092505b8351831015611f6e5760608484815181101515611e5957fe5b90602001906020020151600019169060020a90046001900491508383815181101515611e8157fe5b90602001906020020151600190046bffffffffffffffffffffffff169050611eaa858383612060565b343373ffffffffffffffffffffffffffffffffffffffff167ffc01e439ca3c7015e18b8adea39e270034ba32f41fc788a4ff659842f0f37a93848489604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a38280600101935050611e40565b5050505050565b60008273ffffffffffffffffffffffffffffffffffffffff1614151515611f9b57600080fd5b8173ffffffffffffffffffffffffffffffffffffffff168160405160006040518083038185875af1925050501515611fd257600080fd5b5050565b6000828211151515611fe457fe5b818303905092915050565b60008273ffffffffffffffffffffffffffffffffffffffff161415151561201557600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561205b573d6000803e3d6000fd5b505050565b60008273ffffffffffffffffffffffffffffffffffffffff161415151561208657600080fd5b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3384846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561215d57600080fd5b505af1158015612171573d6000803e3d6000fd5b505050506040513d602081101561218757600080fd5b810190808051906020019092919050505015156121a357600080fd5b5050505600a165627a7a72305820801d9f567c5fcaecc0a8905b08bfc2af619a9701c7faf4da67b09d36471439e10029",
}

// MultisendABI is the input ABI used to generate the binding from.
// Deprecated: Use MultisendMetaData.ABI instead.
var MultisendABI = MultisendMetaData.ABI

// MultisendBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultisendMetaData.Bin instead.
var MultisendBin = MultisendMetaData.Bin

// DeployMultisend deploys a new Ethereum contract, binding an instance of Multisend to it.
func DeployMultisend(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multisend, error) {
	parsed, err := MultisendMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultisendBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multisend{MultisendCaller: MultisendCaller{contract: contract}, MultisendTransactor: MultisendTransactor{contract: contract}, MultisendFilterer: MultisendFilterer{contract: contract}}, nil
}

// Multisend is an auto generated Go binding around an Ethereum contract.
type Multisend struct {
	MultisendCaller     // Read-only binding to the contract
	MultisendTransactor // Write-only binding to the contract
	MultisendFilterer   // Log filterer for contract events
}

// MultisendCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisendSession struct {
	Contract     *Multisend        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisendCallerSession struct {
	Contract *MultisendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultisendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisendTransactorSession struct {
	Contract     *MultisendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultisendRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisendRaw struct {
	Contract *Multisend // Generic contract binding to access the raw methods on
}

// MultisendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisendCallerRaw struct {
	Contract *MultisendCaller // Generic read-only contract binding to access the raw methods on
}

// MultisendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisendTransactorRaw struct {
	Contract *MultisendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisend creates a new instance of Multisend, bound to a specific deployed contract.
func NewMultisend(address common.Address, backend bind.ContractBackend) (*Multisend, error) {
	contract, err := bindMultisend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisend{MultisendCaller: MultisendCaller{contract: contract}, MultisendTransactor: MultisendTransactor{contract: contract}, MultisendFilterer: MultisendFilterer{contract: contract}}, nil
}

// NewMultisendCaller creates a new read-only instance of Multisend, bound to a specific deployed contract.
func NewMultisendCaller(address common.Address, caller bind.ContractCaller) (*MultisendCaller, error) {
	contract, err := bindMultisend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisendCaller{contract: contract}, nil
}

// NewMultisendTransactor creates a new write-only instance of Multisend, bound to a specific deployed contract.
func NewMultisendTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisendTransactor, error) {
	contract, err := bindMultisend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisendTransactor{contract: contract}, nil
}

// NewMultisendFilterer creates a new log filterer instance of Multisend, bound to a specific deployed contract.
func NewMultisendFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisendFilterer, error) {
	contract, err := bindMultisend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisendFilterer{contract: contract}, nil
}

// bindMultisend binds a generic wrapper to an already deployed contract.
func bindMultisend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisend *MultisendRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisend.Contract.MultisendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisend *MultisendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisend *MultisendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisend *MultisendCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisend *MultisendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisend *MultisendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisend.Contract.contract.Transact(opts, method, params...)
}

// EscapeHatchCaller is a free data retrieval call binding the contract method 0x1f6eb6e7.
//
// Solidity: function escapeHatchCaller() view returns(address)
func (_Multisend *MultisendCaller) EscapeHatchCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multisend.contract.Call(opts, &out, "escapeHatchCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscapeHatchCaller is a free data retrieval call binding the contract method 0x1f6eb6e7.
//
// Solidity: function escapeHatchCaller() view returns(address)
func (_Multisend *MultisendSession) EscapeHatchCaller() (common.Address, error) {
	return _Multisend.Contract.EscapeHatchCaller(&_Multisend.CallOpts)
}

// EscapeHatchCaller is a free data retrieval call binding the contract method 0x1f6eb6e7.
//
// Solidity: function escapeHatchCaller() view returns(address)
func (_Multisend *MultisendCallerSession) EscapeHatchCaller() (common.Address, error) {
	return _Multisend.Contract.EscapeHatchCaller(&_Multisend.CallOpts)
}

// EscapeHatchDestination is a free data retrieval call binding the contract method 0xf5b61230.
//
// Solidity: function escapeHatchDestination() view returns(address)
func (_Multisend *MultisendCaller) EscapeHatchDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multisend.contract.Call(opts, &out, "escapeHatchDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscapeHatchDestination is a free data retrieval call binding the contract method 0xf5b61230.
//
// Solidity: function escapeHatchDestination() view returns(address)
func (_Multisend *MultisendSession) EscapeHatchDestination() (common.Address, error) {
	return _Multisend.Contract.EscapeHatchDestination(&_Multisend.CallOpts)
}

// EscapeHatchDestination is a free data retrieval call binding the contract method 0xf5b61230.
//
// Solidity: function escapeHatchDestination() view returns(address)
func (_Multisend *MultisendCallerSession) EscapeHatchDestination() (common.Address, error) {
	return _Multisend.Contract.EscapeHatchDestination(&_Multisend.CallOpts)
}

// IsTokenEscapable is a free data retrieval call binding the contract method 0x892db057.
//
// Solidity: function isTokenEscapable(address _token) view returns(bool)
func (_Multisend *MultisendCaller) IsTokenEscapable(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _Multisend.contract.Call(opts, &out, "isTokenEscapable", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenEscapable is a free data retrieval call binding the contract method 0x892db057.
//
// Solidity: function isTokenEscapable(address _token) view returns(bool)
func (_Multisend *MultisendSession) IsTokenEscapable(_token common.Address) (bool, error) {
	return _Multisend.Contract.IsTokenEscapable(&_Multisend.CallOpts, _token)
}

// IsTokenEscapable is a free data retrieval call binding the contract method 0x892db057.
//
// Solidity: function isTokenEscapable(address _token) view returns(bool)
func (_Multisend *MultisendCallerSession) IsTokenEscapable(_token common.Address) (bool, error) {
	return _Multisend.Contract.IsTokenEscapable(&_Multisend.CallOpts, _token)
}

// NewOwnerCandidate is a free data retrieval call binding the contract method 0xd091b550.
//
// Solidity: function newOwnerCandidate() view returns(address)
func (_Multisend *MultisendCaller) NewOwnerCandidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multisend.contract.Call(opts, &out, "newOwnerCandidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwnerCandidate is a free data retrieval call binding the contract method 0xd091b550.
//
// Solidity: function newOwnerCandidate() view returns(address)
func (_Multisend *MultisendSession) NewOwnerCandidate() (common.Address, error) {
	return _Multisend.Contract.NewOwnerCandidate(&_Multisend.CallOpts)
}

// NewOwnerCandidate is a free data retrieval call binding the contract method 0xd091b550.
//
// Solidity: function newOwnerCandidate() view returns(address)
func (_Multisend *MultisendCallerSession) NewOwnerCandidate() (common.Address, error) {
	return _Multisend.Contract.NewOwnerCandidate(&_Multisend.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Multisend *MultisendCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multisend.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Multisend *MultisendSession) Owner() (common.Address, error) {
	return _Multisend.Contract.Owner(&_Multisend.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Multisend *MultisendCallerSession) Owner() (common.Address, error) {
	return _Multisend.Contract.Owner(&_Multisend.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Multisend *MultisendTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Multisend *MultisendSession) AcceptOwnership() (*types.Transaction, error) {
	return _Multisend.Contract.AcceptOwnership(&_Multisend.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Multisend *MultisendTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Multisend.Contract.AcceptOwnership(&_Multisend.TransactOpts)
}

// ChangeHatchEscapeCaller is a paid mutator transaction binding the contract method 0xd836fbe8.
//
// Solidity: function changeHatchEscapeCaller(address _newEscapeHatchCaller) returns()
func (_Multisend *MultisendTransactor) ChangeHatchEscapeCaller(opts *bind.TransactOpts, _newEscapeHatchCaller common.Address) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "changeHatchEscapeCaller", _newEscapeHatchCaller)
}

// ChangeHatchEscapeCaller is a paid mutator transaction binding the contract method 0xd836fbe8.
//
// Solidity: function changeHatchEscapeCaller(address _newEscapeHatchCaller) returns()
func (_Multisend *MultisendSession) ChangeHatchEscapeCaller(_newEscapeHatchCaller common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.ChangeHatchEscapeCaller(&_Multisend.TransactOpts, _newEscapeHatchCaller)
}

// ChangeHatchEscapeCaller is a paid mutator transaction binding the contract method 0xd836fbe8.
//
// Solidity: function changeHatchEscapeCaller(address _newEscapeHatchCaller) returns()
func (_Multisend *MultisendTransactorSession) ChangeHatchEscapeCaller(_newEscapeHatchCaller common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.ChangeHatchEscapeCaller(&_Multisend.TransactOpts, _newEscapeHatchCaller)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(address _newOwner) returns()
func (_Multisend *MultisendTransactor) ChangeOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "changeOwnership", _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(address _newOwner) returns()
func (_Multisend *MultisendSession) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.ChangeOwnership(&_Multisend.TransactOpts, _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(address _newOwner) returns()
func (_Multisend *MultisendTransactorSession) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.ChangeOwnership(&_Multisend.TransactOpts, _newOwner)
}

// EscapeHatch is a paid mutator transaction binding the contract method 0xa142d608.
//
// Solidity: function escapeHatch(address _token) returns()
func (_Multisend *MultisendTransactor) EscapeHatch(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "escapeHatch", _token)
}

// EscapeHatch is a paid mutator transaction binding the contract method 0xa142d608.
//
// Solidity: function escapeHatch(address _token) returns()
func (_Multisend *MultisendSession) EscapeHatch(_token common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.EscapeHatch(&_Multisend.TransactOpts, _token)
}

// EscapeHatch is a paid mutator transaction binding the contract method 0xa142d608.
//
// Solidity: function escapeHatch(address _token) returns()
func (_Multisend *MultisendTransactorSession) EscapeHatch(_token common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.EscapeHatch(&_Multisend.TransactOpts, _token)
}

// MultiCall is a paid mutator transaction binding the contract method 0x1476e40f.
//
// Solidity: function multiCall(address[] _addresses, uint256[] _amounts) payable returns(bool)
func (_Multisend *MultisendTransactor) MultiCall(opts *bind.TransactOpts, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multiCall", _addresses, _amounts)
}

// MultiCall is a paid mutator transaction binding the contract method 0x1476e40f.
//
// Solidity: function multiCall(address[] _addresses, uint256[] _amounts) payable returns(bool)
func (_Multisend *MultisendSession) MultiCall(_addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.MultiCall(&_Multisend.TransactOpts, _addresses, _amounts)
}

// MultiCall is a paid mutator transaction binding the contract method 0x1476e40f.
//
// Solidity: function multiCall(address[] _addresses, uint256[] _amounts) payable returns(bool)
func (_Multisend *MultisendTransactorSession) MultiCall(_addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.MultiCall(&_Multisend.TransactOpts, _addresses, _amounts)
}

// MultiCallTightlyPacked is a paid mutator transaction binding the contract method 0xac66777f.
//
// Solidity: function multiCallTightlyPacked(bytes32[] _addressesAndAmounts) payable returns(bool)
func (_Multisend *MultisendTransactor) MultiCallTightlyPacked(opts *bind.TransactOpts, _addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multiCallTightlyPacked", _addressesAndAmounts)
}

// MultiCallTightlyPacked is a paid mutator transaction binding the contract method 0xac66777f.
//
// Solidity: function multiCallTightlyPacked(bytes32[] _addressesAndAmounts) payable returns(bool)
func (_Multisend *MultisendSession) MultiCallTightlyPacked(_addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.Contract.MultiCallTightlyPacked(&_Multisend.TransactOpts, _addressesAndAmounts)
}

// MultiCallTightlyPacked is a paid mutator transaction binding the contract method 0xac66777f.
//
// Solidity: function multiCallTightlyPacked(bytes32[] _addressesAndAmounts) payable returns(bool)
func (_Multisend *MultisendTransactorSession) MultiCallTightlyPacked(_addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.Contract.MultiCallTightlyPacked(&_Multisend.TransactOpts, _addressesAndAmounts)
}

// MultiERC20Transfer is a paid mutator transaction binding the contract method 0x35a21728.
//
// Solidity: function multiERC20Transfer(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_Multisend *MultisendTransactor) MultiERC20Transfer(opts *bind.TransactOpts, _token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multiERC20Transfer", _token, _addresses, _amounts)
}

// MultiERC20Transfer is a paid mutator transaction binding the contract method 0x35a21728.
//
// Solidity: function multiERC20Transfer(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_Multisend *MultisendSession) MultiERC20Transfer(_token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.MultiERC20Transfer(&_Multisend.TransactOpts, _token, _addresses, _amounts)
}

// MultiERC20Transfer is a paid mutator transaction binding the contract method 0x35a21728.
//
// Solidity: function multiERC20Transfer(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_Multisend *MultisendTransactorSession) MultiERC20Transfer(_token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.MultiERC20Transfer(&_Multisend.TransactOpts, _token, _addresses, _amounts)
}

// MultiERC20TransferTightlyPacked is a paid mutator transaction binding the contract method 0xf8f1d927.
//
// Solidity: function multiERC20TransferTightlyPacked(address _token, bytes32[] _addressesAndAmounts) returns()
func (_Multisend *MultisendTransactor) MultiERC20TransferTightlyPacked(opts *bind.TransactOpts, _token common.Address, _addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multiERC20TransferTightlyPacked", _token, _addressesAndAmounts)
}

// MultiERC20TransferTightlyPacked is a paid mutator transaction binding the contract method 0xf8f1d927.
//
// Solidity: function multiERC20TransferTightlyPacked(address _token, bytes32[] _addressesAndAmounts) returns()
func (_Multisend *MultisendSession) MultiERC20TransferTightlyPacked(_token common.Address, _addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.Contract.MultiERC20TransferTightlyPacked(&_Multisend.TransactOpts, _token, _addressesAndAmounts)
}

// MultiERC20TransferTightlyPacked is a paid mutator transaction binding the contract method 0xf8f1d927.
//
// Solidity: function multiERC20TransferTightlyPacked(address _token, bytes32[] _addressesAndAmounts) returns()
func (_Multisend *MultisendTransactorSession) MultiERC20TransferTightlyPacked(_token common.Address, _addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.Contract.MultiERC20TransferTightlyPacked(&_Multisend.TransactOpts, _token, _addressesAndAmounts)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _addresses, uint256[] _amounts) payable returns(bool)
func (_Multisend *MultisendTransactor) MultiTransfer(opts *bind.TransactOpts, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multiTransfer", _addresses, _amounts)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _addresses, uint256[] _amounts) payable returns(bool)
func (_Multisend *MultisendSession) MultiTransfer(_addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.MultiTransfer(&_Multisend.TransactOpts, _addresses, _amounts)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _addresses, uint256[] _amounts) payable returns(bool)
func (_Multisend *MultisendTransactorSession) MultiTransfer(_addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.MultiTransfer(&_Multisend.TransactOpts, _addresses, _amounts)
}

// MultiTransferTightlyPacked is a paid mutator transaction binding the contract method 0x2a17e397.
//
// Solidity: function multiTransferTightlyPacked(bytes32[] _addressesAndAmounts) payable returns(bool)
func (_Multisend *MultisendTransactor) MultiTransferTightlyPacked(opts *bind.TransactOpts, _addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multiTransferTightlyPacked", _addressesAndAmounts)
}

// MultiTransferTightlyPacked is a paid mutator transaction binding the contract method 0x2a17e397.
//
// Solidity: function multiTransferTightlyPacked(bytes32[] _addressesAndAmounts) payable returns(bool)
func (_Multisend *MultisendSession) MultiTransferTightlyPacked(_addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.Contract.MultiTransferTightlyPacked(&_Multisend.TransactOpts, _addressesAndAmounts)
}

// MultiTransferTightlyPacked is a paid mutator transaction binding the contract method 0x2a17e397.
//
// Solidity: function multiTransferTightlyPacked(bytes32[] _addressesAndAmounts) payable returns(bool)
func (_Multisend *MultisendTransactorSession) MultiTransferTightlyPacked(_addressesAndAmounts [][32]byte) (*types.Transaction, error) {
	return _Multisend.Contract.MultiTransferTightlyPacked(&_Multisend.TransactOpts, _addressesAndAmounts)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _newOwnerCandidate) returns()
func (_Multisend *MultisendTransactor) ProposeOwnership(opts *bind.TransactOpts, _newOwnerCandidate common.Address) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "proposeOwnership", _newOwnerCandidate)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _newOwnerCandidate) returns()
func (_Multisend *MultisendSession) ProposeOwnership(_newOwnerCandidate common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.ProposeOwnership(&_Multisend.TransactOpts, _newOwnerCandidate)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _newOwnerCandidate) returns()
func (_Multisend *MultisendTransactorSession) ProposeOwnership(_newOwnerCandidate common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.ProposeOwnership(&_Multisend.TransactOpts, _newOwnerCandidate)
}

// RemoveOwnership is a paid mutator transaction binding the contract method 0x666a3427.
//
// Solidity: function removeOwnership(address _dac) returns()
func (_Multisend *MultisendTransactor) RemoveOwnership(opts *bind.TransactOpts, _dac common.Address) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "removeOwnership", _dac)
}

// RemoveOwnership is a paid mutator transaction binding the contract method 0x666a3427.
//
// Solidity: function removeOwnership(address _dac) returns()
func (_Multisend *MultisendSession) RemoveOwnership(_dac common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.RemoveOwnership(&_Multisend.TransactOpts, _dac)
}

// RemoveOwnership is a paid mutator transaction binding the contract method 0x666a3427.
//
// Solidity: function removeOwnership(address _dac) returns()
func (_Multisend *MultisendTransactorSession) RemoveOwnership(_dac common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.RemoveOwnership(&_Multisend.TransactOpts, _dac)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Multisend *MultisendTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Multisend.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Multisend *MultisendSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Multisend.Contract.Fallback(&_Multisend.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Multisend *MultisendTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Multisend.Contract.Fallback(&_Multisend.TransactOpts, calldata)
}

// MultisendEscapeHatchBlackistedTokenIterator is returned from FilterEscapeHatchBlackistedToken and is used to iterate over the raw logs and unpacked data for EscapeHatchBlackistedToken events raised by the Multisend contract.
type MultisendEscapeHatchBlackistedTokenIterator struct {
	Event *MultisendEscapeHatchBlackistedToken // Event containing the contract specifics and raw log

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
func (it *MultisendEscapeHatchBlackistedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendEscapeHatchBlackistedToken)
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
		it.Event = new(MultisendEscapeHatchBlackistedToken)
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
func (it *MultisendEscapeHatchBlackistedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendEscapeHatchBlackistedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendEscapeHatchBlackistedToken represents a EscapeHatchBlackistedToken event raised by the Multisend contract.
type MultisendEscapeHatchBlackistedToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEscapeHatchBlackistedToken is a free log retrieval operation binding the contract event 0x6b44fa019116f0ba92616e27b01166e395abfb95f2d81476d0a179a538a16bb1.
//
// Solidity: event EscapeHatchBlackistedToken(address token)
func (_Multisend *MultisendFilterer) FilterEscapeHatchBlackistedToken(opts *bind.FilterOpts) (*MultisendEscapeHatchBlackistedTokenIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "EscapeHatchBlackistedToken")
	if err != nil {
		return nil, err
	}
	return &MultisendEscapeHatchBlackistedTokenIterator{contract: _Multisend.contract, event: "EscapeHatchBlackistedToken", logs: logs, sub: sub}, nil
}

// WatchEscapeHatchBlackistedToken is a free log subscription operation binding the contract event 0x6b44fa019116f0ba92616e27b01166e395abfb95f2d81476d0a179a538a16bb1.
//
// Solidity: event EscapeHatchBlackistedToken(address token)
func (_Multisend *MultisendFilterer) WatchEscapeHatchBlackistedToken(opts *bind.WatchOpts, sink chan<- *MultisendEscapeHatchBlackistedToken) (event.Subscription, error) {

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "EscapeHatchBlackistedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendEscapeHatchBlackistedToken)
				if err := _Multisend.contract.UnpackLog(event, "EscapeHatchBlackistedToken", log); err != nil {
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

// ParseEscapeHatchBlackistedToken is a log parse operation binding the contract event 0x6b44fa019116f0ba92616e27b01166e395abfb95f2d81476d0a179a538a16bb1.
//
// Solidity: event EscapeHatchBlackistedToken(address token)
func (_Multisend *MultisendFilterer) ParseEscapeHatchBlackistedToken(log types.Log) (*MultisendEscapeHatchBlackistedToken, error) {
	event := new(MultisendEscapeHatchBlackistedToken)
	if err := _Multisend.contract.UnpackLog(event, "EscapeHatchBlackistedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisendEscapeHatchCalledIterator is returned from FilterEscapeHatchCalled and is used to iterate over the raw logs and unpacked data for EscapeHatchCalled events raised by the Multisend contract.
type MultisendEscapeHatchCalledIterator struct {
	Event *MultisendEscapeHatchCalled // Event containing the contract specifics and raw log

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
func (it *MultisendEscapeHatchCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendEscapeHatchCalled)
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
		it.Event = new(MultisendEscapeHatchCalled)
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
func (it *MultisendEscapeHatchCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendEscapeHatchCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendEscapeHatchCalled represents a EscapeHatchCalled event raised by the Multisend contract.
type MultisendEscapeHatchCalled struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEscapeHatchCalled is a free log retrieval operation binding the contract event 0xa50dde912fa22ea0d215a0236093ac45b4d55d6ef0c604c319f900029c5d10f2.
//
// Solidity: event EscapeHatchCalled(address token, uint256 amount)
func (_Multisend *MultisendFilterer) FilterEscapeHatchCalled(opts *bind.FilterOpts) (*MultisendEscapeHatchCalledIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "EscapeHatchCalled")
	if err != nil {
		return nil, err
	}
	return &MultisendEscapeHatchCalledIterator{contract: _Multisend.contract, event: "EscapeHatchCalled", logs: logs, sub: sub}, nil
}

// WatchEscapeHatchCalled is a free log subscription operation binding the contract event 0xa50dde912fa22ea0d215a0236093ac45b4d55d6ef0c604c319f900029c5d10f2.
//
// Solidity: event EscapeHatchCalled(address token, uint256 amount)
func (_Multisend *MultisendFilterer) WatchEscapeHatchCalled(opts *bind.WatchOpts, sink chan<- *MultisendEscapeHatchCalled) (event.Subscription, error) {

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "EscapeHatchCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendEscapeHatchCalled)
				if err := _Multisend.contract.UnpackLog(event, "EscapeHatchCalled", log); err != nil {
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

// ParseEscapeHatchCalled is a log parse operation binding the contract event 0xa50dde912fa22ea0d215a0236093ac45b4d55d6ef0c604c319f900029c5d10f2.
//
// Solidity: event EscapeHatchCalled(address token, uint256 amount)
func (_Multisend *MultisendFilterer) ParseEscapeHatchCalled(log types.Log) (*MultisendEscapeHatchCalled, error) {
	event := new(MultisendEscapeHatchCalled)
	if err := _Multisend.contract.UnpackLog(event, "EscapeHatchCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisendMultiCallIterator is returned from FilterMultiCall and is used to iterate over the raw logs and unpacked data for MultiCall events raised by the Multisend contract.
type MultisendMultiCallIterator struct {
	Event *MultisendMultiCall // Event containing the contract specifics and raw log

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
func (it *MultisendMultiCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendMultiCall)
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
		it.Event = new(MultisendMultiCall)
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
func (it *MultisendMultiCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendMultiCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendMultiCall represents a MultiCall event raised by the Multisend contract.
type MultisendMultiCall struct {
	From   common.Address
	Value  *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMultiCall is a free log retrieval operation binding the contract event 0xa5c03fd81f783c38e4b274cd539e517bacfc824ee513ec4dad385269fa9684d7.
//
// Solidity: event MultiCall(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_Multisend *MultisendFilterer) FilterMultiCall(opts *bind.FilterOpts, _from []common.Address, _value []*big.Int) (*MultisendMultiCallIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "MultiCall", _fromRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return &MultisendMultiCallIterator{contract: _Multisend.contract, event: "MultiCall", logs: logs, sub: sub}, nil
}

// WatchMultiCall is a free log subscription operation binding the contract event 0xa5c03fd81f783c38e4b274cd539e517bacfc824ee513ec4dad385269fa9684d7.
//
// Solidity: event MultiCall(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_Multisend *MultisendFilterer) WatchMultiCall(opts *bind.WatchOpts, sink chan<- *MultisendMultiCall, _from []common.Address, _value []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "MultiCall", _fromRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendMultiCall)
				if err := _Multisend.contract.UnpackLog(event, "MultiCall", log); err != nil {
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

// ParseMultiCall is a log parse operation binding the contract event 0xa5c03fd81f783c38e4b274cd539e517bacfc824ee513ec4dad385269fa9684d7.
//
// Solidity: event MultiCall(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_Multisend *MultisendFilterer) ParseMultiCall(log types.Log) (*MultisendMultiCall, error) {
	event := new(MultisendMultiCall)
	if err := _Multisend.contract.UnpackLog(event, "MultiCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisendMultiERC20TransferIterator is returned from FilterMultiERC20Transfer and is used to iterate over the raw logs and unpacked data for MultiERC20Transfer events raised by the Multisend contract.
type MultisendMultiERC20TransferIterator struct {
	Event *MultisendMultiERC20Transfer // Event containing the contract specifics and raw log

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
func (it *MultisendMultiERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendMultiERC20Transfer)
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
		it.Event = new(MultisendMultiERC20Transfer)
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
func (it *MultisendMultiERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendMultiERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendMultiERC20Transfer represents a MultiERC20Transfer event raised by the Multisend contract.
type MultisendMultiERC20Transfer struct {
	From   common.Address
	Value  *big.Int
	To     common.Address
	Amount *big.Int
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMultiERC20Transfer is a free log retrieval operation binding the contract event 0xfc01e439ca3c7015e18b8adea39e270034ba32f41fc788a4ff659842f0f37a93.
//
// Solidity: event MultiERC20Transfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount, address _token)
func (_Multisend *MultisendFilterer) FilterMultiERC20Transfer(opts *bind.FilterOpts, _from []common.Address, _value []*big.Int) (*MultisendMultiERC20TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "MultiERC20Transfer", _fromRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return &MultisendMultiERC20TransferIterator{contract: _Multisend.contract, event: "MultiERC20Transfer", logs: logs, sub: sub}, nil
}

// WatchMultiERC20Transfer is a free log subscription operation binding the contract event 0xfc01e439ca3c7015e18b8adea39e270034ba32f41fc788a4ff659842f0f37a93.
//
// Solidity: event MultiERC20Transfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount, address _token)
func (_Multisend *MultisendFilterer) WatchMultiERC20Transfer(opts *bind.WatchOpts, sink chan<- *MultisendMultiERC20Transfer, _from []common.Address, _value []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "MultiERC20Transfer", _fromRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendMultiERC20Transfer)
				if err := _Multisend.contract.UnpackLog(event, "MultiERC20Transfer", log); err != nil {
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

// ParseMultiERC20Transfer is a log parse operation binding the contract event 0xfc01e439ca3c7015e18b8adea39e270034ba32f41fc788a4ff659842f0f37a93.
//
// Solidity: event MultiERC20Transfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount, address _token)
func (_Multisend *MultisendFilterer) ParseMultiERC20Transfer(log types.Log) (*MultisendMultiERC20Transfer, error) {
	event := new(MultisendMultiERC20Transfer)
	if err := _Multisend.contract.UnpackLog(event, "MultiERC20Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisendMultiTransferIterator is returned from FilterMultiTransfer and is used to iterate over the raw logs and unpacked data for MultiTransfer events raised by the Multisend contract.
type MultisendMultiTransferIterator struct {
	Event *MultisendMultiTransfer // Event containing the contract specifics and raw log

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
func (it *MultisendMultiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendMultiTransfer)
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
		it.Event = new(MultisendMultiTransfer)
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
func (it *MultisendMultiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendMultiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendMultiTransfer represents a MultiTransfer event raised by the Multisend contract.
type MultisendMultiTransfer struct {
	From   common.Address
	Value  *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMultiTransfer is a free log retrieval operation binding the contract event 0x319e0008dcdeba1f31169497fb0f35d31b2b2f481c714d1f50640e86ac6c3bb0.
//
// Solidity: event MultiTransfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_Multisend *MultisendFilterer) FilterMultiTransfer(opts *bind.FilterOpts, _from []common.Address, _value []*big.Int) (*MultisendMultiTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "MultiTransfer", _fromRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return &MultisendMultiTransferIterator{contract: _Multisend.contract, event: "MultiTransfer", logs: logs, sub: sub}, nil
}

// WatchMultiTransfer is a free log subscription operation binding the contract event 0x319e0008dcdeba1f31169497fb0f35d31b2b2f481c714d1f50640e86ac6c3bb0.
//
// Solidity: event MultiTransfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_Multisend *MultisendFilterer) WatchMultiTransfer(opts *bind.WatchOpts, sink chan<- *MultisendMultiTransfer, _from []common.Address, _value []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "MultiTransfer", _fromRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendMultiTransfer)
				if err := _Multisend.contract.UnpackLog(event, "MultiTransfer", log); err != nil {
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

// ParseMultiTransfer is a log parse operation binding the contract event 0x319e0008dcdeba1f31169497fb0f35d31b2b2f481c714d1f50640e86ac6c3bb0.
//
// Solidity: event MultiTransfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_Multisend *MultisendFilterer) ParseMultiTransfer(log types.Log) (*MultisendMultiTransfer, error) {
	event := new(MultisendMultiTransfer)
	if err := _Multisend.contract.UnpackLog(event, "MultiTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisendOwnershipRemovedIterator is returned from FilterOwnershipRemoved and is used to iterate over the raw logs and unpacked data for OwnershipRemoved events raised by the Multisend contract.
type MultisendOwnershipRemovedIterator struct {
	Event *MultisendOwnershipRemoved // Event containing the contract specifics and raw log

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
func (it *MultisendOwnershipRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendOwnershipRemoved)
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
		it.Event = new(MultisendOwnershipRemoved)
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
func (it *MultisendOwnershipRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendOwnershipRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendOwnershipRemoved represents a OwnershipRemoved event raised by the Multisend contract.
type MultisendOwnershipRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRemoved is a free log retrieval operation binding the contract event 0x94e8b32e01b9eedfddd778ffbd051a7718cdc14781702884561162dca6f74dbb.
//
// Solidity: event OwnershipRemoved()
func (_Multisend *MultisendFilterer) FilterOwnershipRemoved(opts *bind.FilterOpts) (*MultisendOwnershipRemovedIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "OwnershipRemoved")
	if err != nil {
		return nil, err
	}
	return &MultisendOwnershipRemovedIterator{contract: _Multisend.contract, event: "OwnershipRemoved", logs: logs, sub: sub}, nil
}

// WatchOwnershipRemoved is a free log subscription operation binding the contract event 0x94e8b32e01b9eedfddd778ffbd051a7718cdc14781702884561162dca6f74dbb.
//
// Solidity: event OwnershipRemoved()
func (_Multisend *MultisendFilterer) WatchOwnershipRemoved(opts *bind.WatchOpts, sink chan<- *MultisendOwnershipRemoved) (event.Subscription, error) {

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "OwnershipRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendOwnershipRemoved)
				if err := _Multisend.contract.UnpackLog(event, "OwnershipRemoved", log); err != nil {
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

// ParseOwnershipRemoved is a log parse operation binding the contract event 0x94e8b32e01b9eedfddd778ffbd051a7718cdc14781702884561162dca6f74dbb.
//
// Solidity: event OwnershipRemoved()
func (_Multisend *MultisendFilterer) ParseOwnershipRemoved(log types.Log) (*MultisendOwnershipRemoved, error) {
	event := new(MultisendOwnershipRemoved)
	if err := _Multisend.contract.UnpackLog(event, "OwnershipRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisendOwnershipRequestedIterator is returned from FilterOwnershipRequested and is used to iterate over the raw logs and unpacked data for OwnershipRequested events raised by the Multisend contract.
type MultisendOwnershipRequestedIterator struct {
	Event *MultisendOwnershipRequested // Event containing the contract specifics and raw log

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
func (it *MultisendOwnershipRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendOwnershipRequested)
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
		it.Event = new(MultisendOwnershipRequested)
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
func (it *MultisendOwnershipRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendOwnershipRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendOwnershipRequested represents a OwnershipRequested event raised by the Multisend contract.
type MultisendOwnershipRequested struct {
	By  common.Address
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRequested is a free log retrieval operation binding the contract event 0x13a4b3bc0d5234dd3d87c9f1557d8faefa37986da62c36ba49309e2fb2c9aec4.
//
// Solidity: event OwnershipRequested(address indexed by, address indexed to)
func (_Multisend *MultisendFilterer) FilterOwnershipRequested(opts *bind.FilterOpts, by []common.Address, to []common.Address) (*MultisendOwnershipRequestedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "OwnershipRequested", byRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultisendOwnershipRequestedIterator{contract: _Multisend.contract, event: "OwnershipRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipRequested is a free log subscription operation binding the contract event 0x13a4b3bc0d5234dd3d87c9f1557d8faefa37986da62c36ba49309e2fb2c9aec4.
//
// Solidity: event OwnershipRequested(address indexed by, address indexed to)
func (_Multisend *MultisendFilterer) WatchOwnershipRequested(opts *bind.WatchOpts, sink chan<- *MultisendOwnershipRequested, by []common.Address, to []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "OwnershipRequested", byRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendOwnershipRequested)
				if err := _Multisend.contract.UnpackLog(event, "OwnershipRequested", log); err != nil {
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

// ParseOwnershipRequested is a log parse operation binding the contract event 0x13a4b3bc0d5234dd3d87c9f1557d8faefa37986da62c36ba49309e2fb2c9aec4.
//
// Solidity: event OwnershipRequested(address indexed by, address indexed to)
func (_Multisend *MultisendFilterer) ParseOwnershipRequested(log types.Log) (*MultisendOwnershipRequested, error) {
	event := new(MultisendOwnershipRequested)
	if err := _Multisend.contract.UnpackLog(event, "OwnershipRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisendOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Multisend contract.
type MultisendOwnershipTransferredIterator struct {
	Event *MultisendOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MultisendOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendOwnershipTransferred)
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
		it.Event = new(MultisendOwnershipTransferred)
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
func (it *MultisendOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendOwnershipTransferred represents a OwnershipTransferred event raised by the Multisend contract.
type MultisendOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Multisend *MultisendFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultisendOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultisendOwnershipTransferredIterator{contract: _Multisend.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Multisend *MultisendFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultisendOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendOwnershipTransferred)
				if err := _Multisend.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Multisend *MultisendFilterer) ParseOwnershipTransferred(log types.Log) (*MultisendOwnershipTransferred, error) {
	event := new(MultisendOwnershipTransferred)
	if err := _Multisend.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
