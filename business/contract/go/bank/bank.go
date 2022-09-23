// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"Drain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"betId\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"bettors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nonce\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"PlaceBetsSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"betId\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"winners\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ReconcileSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506125a3806100606000396000f3fe60806040526004361061007b5760003560e01c8063a8ffed961161004e578063a8ffed96146100fa578063b4a99a4e14610123578063e63f341f1461014e578063ed21248c1461018b5761007b565b80630ef678871461008057806357ea89b6146100ab57806382156760146100b557806393f28237146100de575b600080fd5b34801561008c57600080fd5b50610095610195565b6040516100a291906110ab565b60405180910390f35b6100b36101dc565b005b3480156100c157600080fd5b506100dc60048036038101906100d791906113e1565b610371565b005b6100f860048036038101906100f391906114bb565b61060d565b005b34801561010657600080fd5b50610121600480360381019061011c9190611731565b610719565b005b34801561012f57600080fd5b50610138610a84565b60405161014591906118b2565b60405180910390f35b34801561015a57600080fd5b50610175600480360381019061017091906114bb565b610aa8565b60405161018291906110ab565b60405180910390f35b610193610b4a565b005b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008103610268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025f9061192a565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156102ae573d6000803e3d6000fd5b5080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102fe9190611979565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61032f33610c49565b61033883610e0c565b604051602001610349929190611a90565b6040516020818303038152906040526040516103659190611b1a565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103c957600080fd5b60006103d788888888610f94565b90506000600182868686604051600081526020016040526040516103fe9493929190611b5a565b6020604051602081039080840390855afa158015610420573d6000803e3d6000fd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff161461049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049190611beb565b60405180910390fd5b6000885160028b6040516104ae9190611c0b565b9081526020016040518091039020546104c79190611c51565b905060005b89518110156105dc5781600160008c84815181106104ed576104ec611c82565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461053e9190611cb1565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a8b61058a8c848151811061057d5761057c611c82565b5b6020026020010151610c49565b61059385610e0c565b6040516020016105a593929190611d57565b6040516020818303038152906040526040516105c19190611b1a565b60405180910390a180806105d490611dc4565b9150506104cc565b50600060028b6040516105ef9190611c0b565b90815260200160405180910390208190555050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066557600080fd5b600047905060008290508073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156106b5573d6000803e3d6000fd5b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106e083610e0c565b6040516020016106f09190611e32565b60405160208183030381529060405260405161070c9190611b1a565b60405180910390a1505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461077157600080fd5b60005b8851811015610a785760006107c18b8b848151811061079657610795611c82565b5b60200260200101518b8b8b8b88815181106107b4576107b3611c82565b5b6020026020010151610fd5565b905060006001828785815181106107db576107da611c82565b5b60200260200101518786815181106107f6576107f5611c82565b5b602002602001015187878151811061081157610810611c82565b5b6020026020010151604051600081526020016040526040516108369493929190611b5a565b6020604051602081039080840390855afa158015610858573d6000803e3d6000fd5b5050506020604051035190508a838151811061087757610876611c82565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146108ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e390611eb3565b60405180910390fd5b88600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561096e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096590611f1f565b60405180910390fd5b8860028d60405161097f9190611c0b565b9081526020016040518091039020600082825461099c9190611cb1565b9250508190555088600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109f29190611979565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a8c610a2483610c49565b610a2d8c610e0c565b604051602001610a3f93929190611f65565b604051602081830303815290604052604051610a5b9190611b1a565b60405180910390a150508080610a7090611dc4565b915050610774565b50505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b0357600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b999190611cb1565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610bca33610c49565b610c12600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e0c565b604051602001610c2392919061201e565b604051602081830303815290604052604051610c3f9190611b1a565b60405180910390a1565b60606000602867ffffffffffffffff811115610c6857610c676110f5565b5b6040519080825280601f01601f191660200182016040528015610c9a5781602001600182028036833780820191505090505b50905060005b6014811015610e02576000816013610cb89190611979565b6008610cc4919061206f565b6002610cd091906121e4565b8573ffffffffffffffffffffffffffffffffffffffff16610cf19190611c51565b60f81b9050600060108260f81c610d08919061222f565b60f81b905060008160f81c6010610d1f9190612260565b8360f81c610d2d919061229d565b60f81b9050610d3b8261101c565b85856002610d49919061206f565b81518110610d5a57610d59611c82565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610d928161101c565b856001866002610da2919061206f565b610dac9190611cb1565b81518110610dbd57610dbc611c82565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610dfa90611dc4565b915050610ca0565b5080915050919050565b606060008203610e53576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610f8f565b600082905060005b60008214610e85578080610e6e90611dc4565b915050600a82610e7e9190611c51565b9150610e5b565b60008167ffffffffffffffff811115610ea157610ea06110f5565b5b6040519080825280601f01601f191660200182016040528015610ed35781602001600182028036833780820191505090505b50905060008290505b60008614610f8757600181610ef19190611979565b90506000600a8088610f039190611c51565b610f0d919061206f565b87610f189190611979565b6030610f2491906122d2565b905060008160f81b905080848481518110610f4257610f41611c82565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610f7e9190611c51565b97505050610edc565b819450505050505b919050565b6000610fcb85858585604051602001610fb09493929190612428565b60405160208183030381529060405280519060200120611062565b9050949350505050565b6000611010878787878787604051602001610ff59695949392919061246e565b60405160208183030381529060405280519060200120611062565b90509695505050505050565b6000600a8260f81c60ff1610156110475760308260f81c61103d91906122d2565b60f81b905061105d565b60578260f81c61105791906122d2565b60f81b90505b919050565b6000816040516020016110759190612547565b604051602081830303815290604052805190602001209050919050565b6000819050919050565b6110a581611092565b82525050565b60006020820190506110c0600083018461109c565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61112d826110e4565b810181811067ffffffffffffffff8211171561114c5761114b6110f5565b5b80604052505050565b600061115f6110c6565b905061116b8282611124565b919050565b600067ffffffffffffffff82111561118b5761118a6110f5565b5b611194826110e4565b9050602081019050919050565b82818337600083830152505050565b60006111c36111be84611170565b611155565b9050828152602081018484840111156111df576111de6110df565b5b6111ea8482856111a1565b509392505050565b600082601f830112611207576112066110da565b5b81356112178482602086016111b0565b91505092915050565b600067ffffffffffffffff82111561123b5761123a6110f5565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061127c82611251565b9050919050565b61128c81611271565b811461129757600080fd5b50565b6000813590506112a981611283565b92915050565b60006112c26112bd84611220565b611155565b905080838252602082019050602084028301858111156112e5576112e461124c565b5b835b8181101561130e57806112fa888261129a565b8452602084019350506020810190506112e7565b5050509392505050565b600082601f83011261132d5761132c6110da565b5b813561133d8482602086016112af565b91505092915050565b61134f81611092565b811461135a57600080fd5b50565b60008135905061136c81611346565b92915050565b600060ff82169050919050565b61138881611372565b811461139357600080fd5b50565b6000813590506113a58161137f565b92915050565b6000819050919050565b6113be816113ab565b81146113c957600080fd5b50565b6000813590506113db816113b5565b92915050565b600080600080600080600060e0888a031215611400576113ff6110d0565b5b600088013567ffffffffffffffff81111561141e5761141d6110d5565b5b61142a8a828b016111f2565b975050602088013567ffffffffffffffff81111561144b5761144a6110d5565b5b6114578a828b01611318565b96505060406114688a828b0161129a565b95505060606114798a828b0161135d565b945050608061148a8a828b01611396565b93505060a061149b8a828b016113cc565b92505060c06114ac8a828b016113cc565b91505092959891949750929550565b6000602082840312156114d1576114d06110d0565b5b60006114df8482850161129a565b91505092915050565b600067ffffffffffffffff821115611503576115026110f5565b5b602082029050602081019050919050565b6000611527611522846114e8565b611155565b9050808382526020820190506020840283018581111561154a5761154961124c565b5b835b81811015611573578061155f888261135d565b84526020840193505060208101905061154c565b5050509392505050565b600082601f830112611592576115916110da565b5b81356115a2848260208601611514565b91505092915050565b600067ffffffffffffffff8211156115c6576115c56110f5565b5b602082029050602081019050919050565b60006115ea6115e5846115ab565b611155565b9050808382526020820190506020840283018581111561160d5761160c61124c565b5b835b8181101561163657806116228882611396565b84526020840193505060208101905061160f565b5050509392505050565b600082601f830112611655576116546110da565b5b81356116658482602086016115d7565b91505092915050565b600067ffffffffffffffff821115611689576116886110f5565b5b602082029050602081019050919050565b60006116ad6116a88461166e565b611155565b905080838252602082019050602084028301858111156116d0576116cf61124c565b5b835b818110156116f957806116e588826113cc565b8452602084019350506020810190506116d2565b5050509392505050565b600082601f830112611718576117176110da565b5b813561172884826020860161169a565b91505092915050565b60008060008060008060008060006101208a8c031215611754576117536110d0565b5b60008a013567ffffffffffffffff811115611772576117716110d5565b5b61177e8c828d016111f2565b99505060208a013567ffffffffffffffff81111561179f5761179e6110d5565b5b6117ab8c828d01611318565b98505060406117bc8c828d0161129a565b97505060606117cd8c828d0161135d565b96505060806117de8c828d0161135d565b95505060a08a013567ffffffffffffffff8111156117ff576117fe6110d5565b5b61180b8c828d0161157d565b94505060c08a013567ffffffffffffffff81111561182c5761182b6110d5565b5b6118388c828d01611640565b93505060e08a013567ffffffffffffffff811115611859576118586110d5565b5b6118658c828d01611703565b9250506101008a013567ffffffffffffffff811115611887576118866110d5565b5b6118938c828d01611703565b9150509295985092959850929598565b6118ac81611271565b82525050565b60006020820190506118c760008301846118a3565b92915050565b600082825260208201905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b60006119146012836118cd565b915061191f826118de565b602082019050919050565b6000602082019050818103600083015261194381611907565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061198482611092565b915061198f83611092565b92508282039050818111156119a7576119a661194a565b5b92915050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b83811015611a075780820151818401526020810190506119ec565b60008484015250505050565b6000611a1e826119d3565b611a2881856119de565b9350611a388185602086016119e9565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000611a9b826119ad565b600982019150611aab8285611a13565b9150611ab682611a44565b600982019150611ac68284611a13565b9150611ad182611a6a565b6001820191508190509392505050565b6000611aec826119d3565b611af681856118cd565b9350611b068185602086016119e9565b611b0f816110e4565b840191505092915050565b60006020820190508181036000830152611b348184611ae1565b905092915050565b611b45816113ab565b82525050565b611b5481611372565b82525050565b6000608082019050611b6f6000830187611b3c565b611b7c6020830186611b4b565b611b896040830185611b3c565b611b966060830184611b3c565b95945050505050565b7f696e76616c6964206d6f64657261746f72207369676e61747572650000000000600082015250565b6000611bd5601b836118cd565b9150611be082611b9f565b602082019050919050565b60006020820190508181036000830152611c0481611bc8565b9050919050565b6000611c178284611a13565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611c5c82611092565b9150611c6783611092565b925082611c7757611c76611c22565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611cbc82611092565b9150611cc783611092565b9250828201905080821115611cdf57611cde61194a565b5b92915050565b7f62657449645b0000000000000000000000000000000000000000000000000000815250565b7f5d20626574746f725b0000000000000000000000000000000000000000000000815250565b7f5d2077696e6e696e67735b000000000000000000000000000000000000000000815250565b6000611d6282611ce5565b600682019150611d728286611a13565b9150611d7d82611d0b565b600982019150611d8d8285611a13565b9150611d9882611d31565b600b82019150611da88284611a13565b9150611db382611a6a565b600182019150819050949350505050565b6000611dcf82611092565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e0157611e0061194a565b5b600182019050919050565b7f7472616e736665725b0000000000000000000000000000000000000000000000815250565b6000611e3d82611e0c565b600982019150611e4d8284611a13565b9150611e5882611a6a565b60018201915081905092915050565b7f696e76616c696420626574746f72000000000000000000000000000000000000600082015250565b6000611e9d600e836118cd565b9150611ea882611e67565b602082019050919050565b60006020820190508181036000830152611ecc81611e90565b9050919050565b7f696e73756666696369656e742066756e64730000000000000000000000000000600082015250565b6000611f096012836118cd565b9150611f1482611ed3565b602082019050919050565b60006020820190508181036000830152611f3881611efc565b9050919050565b7f5d206265745b0000000000000000000000000000000000000000000000000000815250565b6000611f7082611ce5565b600682019150611f808286611a13565b9150611f8b82611d0b565b600982019150611f9b8285611a13565b9150611fa682611f3f565b600682019150611fb68284611a13565b9150611fc182611a6a565b600182019150819050949350505050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b600061202982611fd2565b6008820191506120398285611a13565b915061204482611ff8565b600a820191506120548284611a13565b915061205f82611a6a565b6001820191508190509392505050565b600061207a82611092565b915061208583611092565b925082820261209381611092565b915082820484148315176120aa576120a961194a565b5b5092915050565b60008160011c9050919050565b6000808291508390505b6001851115612108578086048111156120e4576120e361194a565b5b60018516156120f35780820291505b8081029050612101856120b1565b94506120c8565b94509492505050565b60008261212157600190506121dd565b8161212f57600090506121dd565b8160018114612145576002811461214f5761217e565b60019150506121dd565b60ff8411156121615761216061194a565b5b8360020a9150848211156121785761217761194a565b5b506121dd565b5060208310610133831016604e8410600b84101617156121b35782820a9050838111156121ae576121ad61194a565b5b6121dd565b6121c084848460016120be565b925090508184048111156121d7576121d661194a565b5b81810290505b9392505050565b60006121ef82611092565b91506121fa83611092565b92506122277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612111565b905092915050565b600061223a82611372565b915061224583611372565b92508261225557612254611c22565b5b828204905092915050565b600061226b82611372565b915061227683611372565b925082820261228481611372565b91508082146122965761229561194a565b5b5092915050565b60006122a882611372565b91506122b383611372565b9250828203905060ff8111156122cc576122cb61194a565b5b92915050565b60006122dd82611372565b91506122e883611372565b9250828201905060ff8111156123015761230061194a565b5b92915050565b600081519050919050565b600081905092915050565b6000819050602082019050919050565b61233681611271565b82525050565b6000612348838361232d565b60208301905092915050565b6000602082019050919050565b600061236c82612307565b6123768185612312565b93506123818361231d565b8060005b838110156123b2578151612399888261233c565b97506123a483612354565b925050600181019050612385565b5085935050505092915050565b60008160601b9050919050565b60006123d7826123bf565b9050919050565b60006123e9826123cc565b9050919050565b6124016123fc82611271565b6123de565b82525050565b6000819050919050565b61242261241d82611092565b612407565b82525050565b60006124348287611a13565b91506124408286612361565b915061244c82856123f0565b60148201915061245c8284612411565b60208201915081905095945050505050565b600061247a8289611a13565b915061248682886123f0565b60148201915061249682876123f0565b6014820191506124a68286612411565b6020820191506124b68285612411565b6020820191506124c68284612411565b602082019150819050979650505050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b6000612510601c836119de565b915061251b826124da565b601c82019050919050565b6000819050919050565b61254161253c826113ab565b612526565b82525050565b600061255282612503565b915061255e8284612530565b6020820191508190509291505056fea2646970667358221220c28890db26a855a4e882704492112d87f046c0b27a30f0106241b51052e2dce864736f6c63430008110033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCallerSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCallerSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Drain is a paid mutator transaction binding the contract method 0x93f28237.
//
// Solidity: function Drain(address target) payable returns()
func (_Bank *BankTransactor) Drain(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Drain", target)
}

// Drain is a paid mutator transaction binding the contract method 0x93f28237.
//
// Solidity: function Drain(address target) payable returns()
func (_Bank *BankSession) Drain(target common.Address) (*types.Transaction, error) {
	return _Bank.Contract.Drain(&_Bank.TransactOpts, target)
}

// Drain is a paid mutator transaction binding the contract method 0x93f28237.
//
// Solidity: function Drain(address target) payable returns()
func (_Bank *BankTransactorSession) Drain(target common.Address) (*types.Transaction, error) {
	return _Bank.Contract.Drain(&_Bank.TransactOpts, target)
}

// PlaceBetsSigned is a paid mutator transaction binding the contract method 0xa8ffed96.
//
// Solidity: function PlaceBetsSigned(string betId, address[] bettors, address moderator, uint256 amount, uint256 expiration, uint256[] nonce, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Bank *BankTransactor) PlaceBetsSigned(opts *bind.TransactOpts, betId string, bettors []common.Address, moderator common.Address, amount *big.Int, expiration *big.Int, nonce []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "PlaceBetsSigned", betId, bettors, moderator, amount, expiration, nonce, v, r, s)
}

// PlaceBetsSigned is a paid mutator transaction binding the contract method 0xa8ffed96.
//
// Solidity: function PlaceBetsSigned(string betId, address[] bettors, address moderator, uint256 amount, uint256 expiration, uint256[] nonce, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Bank *BankSession) PlaceBetsSigned(betId string, bettors []common.Address, moderator common.Address, amount *big.Int, expiration *big.Int, nonce []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Bank.Contract.PlaceBetsSigned(&_Bank.TransactOpts, betId, bettors, moderator, amount, expiration, nonce, v, r, s)
}

// PlaceBetsSigned is a paid mutator transaction binding the contract method 0xa8ffed96.
//
// Solidity: function PlaceBetsSigned(string betId, address[] bettors, address moderator, uint256 amount, uint256 expiration, uint256[] nonce, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Bank *BankTransactorSession) PlaceBetsSigned(betId string, bettors []common.Address, moderator common.Address, amount *big.Int, expiration *big.Int, nonce []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Bank.Contract.PlaceBetsSigned(&_Bank.TransactOpts, betId, bettors, moderator, amount, expiration, nonce, v, r, s)
}

// ReconcileSigned is a paid mutator transaction binding the contract method 0x82156760.
//
// Solidity: function ReconcileSigned(string betId, address[] winners, address moderator, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bank *BankTransactor) ReconcileSigned(opts *bind.TransactOpts, betId string, winners []common.Address, moderator common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "ReconcileSigned", betId, winners, moderator, nonce, v, r, s)
}

// ReconcileSigned is a paid mutator transaction binding the contract method 0x82156760.
//
// Solidity: function ReconcileSigned(string betId, address[] winners, address moderator, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bank *BankSession) ReconcileSigned(betId string, winners []common.Address, moderator common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bank.Contract.ReconcileSigned(&_Bank.TransactOpts, betId, winners, moderator, nonce, v, r, s)
}

// ReconcileSigned is a paid mutator transaction binding the contract method 0x82156760.
//
// Solidity: function ReconcileSigned(string betId, address[] winners, address moderator, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bank *BankTransactorSession) ReconcileSigned(betId string, winners []common.Address, moderator common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bank.Contract.ReconcileSigned(&_Bank.TransactOpts, betId, winners, moderator, nonce, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// BankEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bank contract.
type BankEventLogIterator struct {
	Event *BankEventLog // Event containing the contract specifics and raw log

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
func (it *BankEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventLog)
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
		it.Event = new(BankEventLog)
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
func (it *BankEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventLog represents a EventLog event raised by the Bank contract.
type BankEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankEventLogIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankEventLogIterator{contract: _Bank.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankEventLog) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventLog)
				if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
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

// ParseEventLog is a log parse operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) ParseEventLog(log types.Log) (*BankEventLog, error) {
	event := new(BankEventLog)
	if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
