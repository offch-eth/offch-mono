// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"fmt"
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
	_ = fmt.Sprintf
)

// BlockHeader is an auto generated low-level Go binding around an user-defined struct.
type BlockHeader struct {
	ParentHash       [32]byte
	OmmersHash       [32]byte
	Proposer         common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        [8][32]byte
	Difficulty       *big.Int
	Height           *big.Int
	GasLimit         uint64
	GasUsed          uint64
	Timestamp        uint64
	ExtraData        []byte
	MixHash          [32]byte
	Nonce            uint64
	BaseFeePerGas    *big.Int
	WithdrawalsRoot  [32]byte
}

// String BlockHeader returns a readable string representing the user-defined struct.
func (st *BlockHeader) String() string {
	s := "User defined struct: BlockHeader {\n"

	s += fmt.Sprintf("ParentHash: %v,\n", st.ParentHash)
	s += fmt.Sprintf("OmmersHash: %v,\n", st.OmmersHash)
	s += fmt.Sprintf("Proposer: %v,\n", st.Proposer)
	s += fmt.Sprintf("StateRoot: %v,\n", st.StateRoot)
	s += fmt.Sprintf("TransactionsRoot: %v,\n", st.TransactionsRoot)
	s += fmt.Sprintf("ReceiptsRoot: %v,\n", st.ReceiptsRoot)
	s += fmt.Sprintf("LogsBloom: %v,\n", st.LogsBloom)
	s += fmt.Sprintf("Difficulty: %v,\n", st.Difficulty)
	s += fmt.Sprintf("Height: %v,\n", st.Height)
	s += fmt.Sprintf("GasLimit: %v,\n", st.GasLimit)
	s += fmt.Sprintf("GasUsed: %v,\n", st.GasUsed)
	s += fmt.Sprintf("Timestamp: %v,\n", st.Timestamp)
	s += fmt.Sprintf("ExtraData: %v,\n", st.ExtraData)
	s += fmt.Sprintf("MixHash: %v,\n", st.MixHash)
	s += fmt.Sprintf("Nonce: %v,\n", st.Nonce)
	s += fmt.Sprintf("BaseFeePerGas: %v,\n", st.BaseFeePerGas)
	s += fmt.Sprintf("WithdrawalsRoot: %v,\n", st.WithdrawalsRoot)
	s += "}"

	return s
}

// BlockHeaderToRebuild is an auto generated low-level Go binding around an user-defined struct.
type BlockHeaderToRebuild struct {
	BlockNumber uint64
	BlockHeader BlockHeader
}

// String BlockHeaderToRebuild returns a readable string representing the user-defined struct.
func (st *BlockHeaderToRebuild) String() string {
	s := "User defined struct: BlockHeaderToRebuild {\n"

	s += fmt.Sprintf("BlockNumber: %v,\n", st.BlockNumber)
	s += fmt.Sprintf("BlockHeader: %v,\n", st.BlockHeader)
	s += "}"

	return s
}

// ProofsToVerify is an auto generated low-level Go binding around an user-defined struct.
type ProofsToVerify struct {
	AccountAddress  []byte
	AccountState    []byte
	AccountProof    []byte
	StateRoot       [32]byte
	StorageKey      []byte
	TokenAmount     []byte
	StorageProof    []byte
	StorageRootHash [32]byte
}

// String ProofsToVerify returns a readable string representing the user-defined struct.
func (st *ProofsToVerify) String() string {
	s := "User defined struct: ProofsToVerify {\n"

	s += fmt.Sprintf("AccountAddress: %v,\n", st.AccountAddress)
	s += fmt.Sprintf("AccountState: %v,\n", st.AccountState)
	s += fmt.Sprintf("AccountProof: %v,\n", st.AccountProof)
	s += fmt.Sprintf("StateRoot: %v,\n", st.StateRoot)
	s += fmt.Sprintf("StorageKey: %v,\n", st.StorageKey)
	s += fmt.Sprintf("TokenAmount: %v,\n", st.TokenAmount)
	s += fmt.Sprintf("StorageProof: %v,\n", st.StorageProof)
	s += fmt.Sprintf("StorageRootHash: %v,\n", st.StorageRootHash)
	s += "}"

	return s
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blockHashValid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"accountProofValid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"storageProofValid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1BlockHash\",\"type\":\"bytes32\"}],\"name\":\"ProofValidity\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"accountAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"storageKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"tokenAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"storageProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"storageRootHash\",\"type\":\"bytes32\"}],\"internalType\":\"structProofsToVerify\",\"name\":\"content\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"ommersHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[8]\",\"name\":\"logsBloom\",\"type\":\"bytes32[8]\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"height\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasUsed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"baseFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalsRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBlockHeader\",\"name\":\"blockHeader\",\"type\":\"tuple\"}],\"internalType\":\"structBlockHeaderToRebuild\",\"name\":\"blockHeader\",\"type\":\"tuple\"}],\"name\":\"verifyProofWithBlockHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"blockHashValid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"accountProofValid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"storageProofValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506124fc8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80634f7142ad146100385780635af3d05c14610060575b5f80fd5b61004b610046366004611df9565b610092565b60405190151581526020015b60405180910390f35b61007361006e36600461202b565b610145565b6040805193151584529115156020840152151590820152606001610057565b5f61013988888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050604080516020601f8c018190048102820181019092528a815292508a91508990819084018382808284375f9201919091525050604080516020601f8b0181900481028201810190925289815292508991508890819084018382808284375f920191909152508892506102bb915050565b98975050505050505050565b8051604051635eeb5e1b60e11b81526001600160401b0390911660048201525f9081908190731000777700000000000000000000000000000001908290829063bdd6bc3690602401602060405180830381865afa1580156101a8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101cc91906121e7565b5f955090506102026101de88806121fe565b6101eb60208b018b6121fe565b6101f860408d018d6121fe565b8d60600135610092565b935061023861021460808901896121fe565b61022160a08b018b6121fe565b61022e60c08d018d6121fe565b8d60e00135610092565b92508380156102445750825b15610267576020860180516060808a013591015251610262906102de565b811494505b6040805186151581528515156020820152841515818301526060810183905290517f97c4a4e08295bc7f3760c9c4d51c41205ba626d35940818e7f7fcbe381c5f5589181900360800190a150509250925092565b5f806102c686610301565b90506102d481868686610333565b9695505050505050565b5f806102f26102ed845f61036e565b61078e565b80516020909101209392505050565b6060818051906020012060405160200161031d91815260200190565b6040516020818303038152906040529050919050565b5f805f6103418786866107d1565b9150915081801561036357508051602080830191909120875191880191909120145b979650505050505050565b610200820151606090156103d957610387826011612254565b6001600160401b0381111561039e5761039e611e93565b6040519080825280602002602001820160405280156103d157816020015b60608152602001906001900390816103bc5790505b509050610448565b6101e0830151156103ef57610387826010612254565b6103fa82600f612254565b6001600160401b0381111561041157610411611e93565b60405190808252806020026020018201604052801561044457816020015b606081526020019060019003908161042f5790505b5090505b8251610453906108ab565b815f8151811061046557610465612267565b602002602001018190525061047d83602001516108ab565b8160018151811061049057610490612267565b60200260200101819052506104a883604001516108be565b816002815181106104bb576104bb612267565b60200260200101819052506104d383606001516108ab565b816003815181106104e6576104e6612267565b60200260200101819052506104fe83608001516108ab565b8160048151811061051157610511612267565b60200260200101819052506105298360a001516108ab565b8160058151811061053c5761053c612267565b60200260200101819052506105738360c0015160405160200161055f919061227b565b6040516020818303038152906040526108e7565b8160068151811061058657610586612267565b602002602001018190525061059e8360e00151610955565b816007815181106105b1576105b1612267565b60200260200101819052506105d38361010001516001600160801b0316610955565b816008815181106105e6576105e6612267565b60200260200101819052506105ff836101200151610963565b8160098151811061061257610612612267565b602002602001018190525061062b836101400151610963565b81600a8151811061063e5761063e612267565b6020026020010181905250610657836101600151610963565b81600b8151811061066a5761066a612267565b60200260200101819052506106838361018001516108e7565b81600c8151811061069657610696612267565b60200260200101819052506106af836101a001516108ab565b81600d815181106106c2576106c2612267565b60200260200101819052506106f9836101c0015160405160200161055f919060c09190911b6001600160c01b031916815260080190565b81600e8151811061070c5761070c612267565b6020026020010181905250826101e001515f1461075057610731836101e00151610955565b81600f8151811061074457610744612267565b60200260200101819052505b61020083015115610788576107698361020001516108ab565b8160108151811061077c5761077c612267565b60200260200101819052505b92915050565b60605f61079a8361097a565b90506107a8815160c0610a9d565b816040516020016107ba9291906122dc565b604051602081830303815290604052915050919050565b5f60605f6107de85610c40565b90505f805f6107ee848a89610d2b565b815192955090935091501580806108025750815b6108535760405162461bcd60e51b815260206004820152601a60248201527f50726f76696465642070726f6f6620697320696e76616c69642e00000000000060448201526064015b60405180910390fd5b5f8161086d5760405180602001604052805f815250610899565b6108998661087c6001886122f8565b8151811061088c5761088c612267565b6020026020010151611131565b919b919a509098505050505050505050565b60606107886108b983611158565b6108e7565b604051606082811b6bffffffffffffffffffffffff19166020830152906107889060340161055f565b6060808251600114801561091457506080835f8151811061090a5761090a612267565b016020015160f81c105b15610920575081610788565b61092c83516080610a9d565b8360405160200161093e9291906122dc565b604051602081830303815290604052905092915050565b60606107886108b9836111fe565b60606107886108b9836001600160401b03166111fe565b606081515f0361099c57604080515f80825260208201909252905b5092915050565b5f805b83518110156109d8578381815181106109ba576109ba612267565b602002602001015151826109ce9190612254565b915060010161099f565b5f826001600160401b038111156109f1576109f1611e93565b6040519080825280601f01601f191660200182016040528015610a1b576020820181803683370190505b505f92509050602081015b8551831015610a94575f868481518110610a4257610a42612267565b602002602001015190505f602082019050610a5f83828451611312565b878581518110610a7157610a71612267565b60200260200101515183610a859190612254565b92505050826001019250610a26565b50949350505050565b6060806038841015610b025760408051600180825281830190925290602082018180368337019050509050610ad2838561230b565b60f81b815f81518110610ae757610ae7612267565b60200101906001600160f81b03191690815f1a905350610c39565b5f60015b610b108187612338565b15610b3357610b1e8261234b565b9150610b2c61010082612363565b9050610b06565b610b3e826001612254565b6001600160401b03811115610b5557610b55611e93565b6040519080825280601f01601f191660200182016040528015610b7f576020820181803683370190505b509250610b8c858361230b565b610b9790603761230b565b60f81b835f81518110610bac57610bac612267565b60200101906001600160f81b03191690815f1a905350600190505b818111610c3657610100610bdb82846122f8565b610be79061010061245a565b610bf19088612338565b610bfb9190612465565b60f81b838281518110610c1057610c10612267565b60200101906001600160f81b03191690815f1a905350610c2f8161234b565b9050610bc7565b50505b9392505050565b60605f610c4c8361136e565b90505f81516001600160401b03811115610c6857610c68611e93565b604051908082528060200260200182016040528015610cad57816020015b6040805180820190915260608082526020820152815260200190600190039081610c865790505b5090505f5b8251811015610d23575f610cde848381518110610cd157610cd1612267565b60200260200101516113a0565b90506040518060400160405280610cf48361136e565b815260200182815250838381518110610d0f57610d0f612267565b602090810291909101015250600101610cb2565b509392505050565b5f60605f805f610d3a8761142d565b90505f8690505f80610d5f604051806040016040528060608152602001606081525090565b5f5b8c51811015611109578c8181518110610d7c57610d7c612267565b602002602001015191508284610d929190612254565b9350610d9f600188612254565b9650835f03610dfb578482602001518051906020012014610df65760405162461bcd60e51b8152602060048201526011602482015270092dcecc2d8d2c840e4dedee840d0c2e6d607b1b604482015260640161084a565b610ebd565b602082602001515110610e62578482602001518051906020012014610df65760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206c6172676520696e7465726e616c20686173680000000000604482015260640161084a565b84610e708360200151611552565b14610ebd5760405162461bcd60e51b815260206004820152601a60248201527f496e76616c696420696e7465726e616c206e6f64652068617368000000000000604482015260640161084a565b610ec96010600161230b565b60ff16825f01515103610f345785518414611109575f868581518110610ef157610ef1612267565b01602001518351805160f89290921c92505f9183908110610f1457610f14612267565b60200260200101519050610f2781611579565b9650600194505050611101565b815151600119016110b9575f610f49836115ad565b90505f815f81518110610f5e57610f5e612267565b016020015160f81c90505f610f74600283612478565b610f7f906002612499565b90505f610f8f848360ff166115cf565b90505f610f9c8b8a6115cf565b90505f610fa98383611604565b905060ff851660021480610fc0575060ff85166003145b15610ffa57808351148015610fd55750808251145b15610fe757610fe4818b612254565b99505b50600160ff1b9950611109945050505050565b60ff8516158061100d575060ff85166001145b15611062578251811461102d5750600160ff1b9950611109945050505050565b611053885f015160018151811061104657611046612267565b6020026020010151611579565b9a509750611101945050505050565b60405162461bcd60e51b815260206004820152602660248201527f52656365697665642061206e6f6465207769746820616e20756e6b6e6f776e206044820152650e0e4caccd2f60d31b606482015260840161084a565b60405162461bcd60e51b815260206004820152601d60248201527f526563656976656420616e20756e706172736561626c65206e6f64652e000000604482015260640161084a565b600101610d61565b50600160ff1b84148661111c87866115cf565b909e909d50909b509950505050505050505050565b8051805160609161078891611148906001906122f8565b81518110610cd157610cd1612267565b60605f8260405160200161116e91815260200190565b60408051808303601f19018152602080845283830190925292505f9182916020820181803683370190505090505f5b8151811015610a945783836111b18161234b565b9450815181106111c3576111c3612267565b602001015160f81c60f81b8282815181106111e0576111e0612267565b60200101906001600160f81b03191690815f1a90535060010161119d565b60605f8260405160200161121491815260200190565b60405160208183030381529060405290505f5b602081101561125f5781818151811061124257611242612267565b01602001516001600160f81b0319165f0361125f57600101611227565b5f61126b8260206122f8565b6001600160401b0381111561128257611282611e93565b6040519080825280601f01601f1916602001820160405280156112ac576020820181803683370190505b5090505f5b8151811015610a945783836112c58161234b565b9450815181106112d7576112d7612267565b602001015160f81c60f81b8282815181106112f4576112f4612267565b60200101906001600160f81b03191690815f1a9053506001016112b1565b8282825b6020811061134e578151835261132d602084612254565b925061133a602083612254565b91506113476020826122f8565b9050611316565b905182516020929092036101000a5f190180199091169116179052505050565b6040805180820182525f808252602091820152815180830190925282518252808301908201526060906107889061167d565b60605f805f6113ae85611863565b919450925090505f8160018111156113c8576113c86124b2565b146114155760405162461bcd60e51b815260206004820152601860248201527f496e76616c696420524c502062797465732076616c75652e0000000000000000604482015260640161084a565b61142485602001518484611b9d565b95945050505050565b60605f8251600261143e9190612363565b6001600160401b0381111561145557611455611e93565b6040519080825280601f01601f19166020018201604052801561147f576020820181803683370190505b5090505f5b83518110156109955760048482815181106114a1576114a1612267565b01602001516001600160f81b031916901c826114be836002612363565b815181106114ce576114ce612267565b60200101906001600160f81b03191690815f1a90535060108482815181106114f8576114f8612267565b016020015161150a919060f81c612478565b60f81b82611519836002612363565b611524906001612254565b8151811061153457611534612267565b60200101906001600160f81b03191690815f1a905350600101611484565b5f60208251101561156557506020015190565b8180602001905181019061078891906121e7565b5f60606020835f015110156115985761159183611c41565b90506115a4565b6115a1836113a0565b90505b610c3981611552565b60606107886115ca835f01515f81518110610cd157610cd1612267565b61142d565b6060825182106115ed575060408051602081019091525f8152610788565b610c3983838486516115ff91906122f8565b611c4c565b5f805b8084511180156116175750808351115b8015611668575082818151811061163057611630612267565b602001015160f81c60f81b6001600160f81b03191684828151811061165757611657612267565b01602001516001600160f81b031916145b15610c39576116768161234b565b9050611607565b60605f8061168a84611863565b919350909150600190508160018111156116a6576116a66124b2565b146116f35760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420524c50206c6973742076616c75652e000000000000000000604482015260640161084a565b60408051602080825261042082019092525f91816020015b604080518082019091525f808252602082015281526020019060019003908161170b5790505090505f835b865181101561185857602082106117a25760405162461bcd60e51b815260206004820152602a60248201527f50726f766964656420524c50206c6973742065786365656473206d6178206c6960448201526939ba103632b733ba341760b11b606482015260840161084a565b5f806117dd6040518060400160405280858c5f01516117c191906122f8565b8152602001858c602001516117d69190612254565b9052611863565b5091509150604051806040016040528083836117f99190612254565b8152602001848b6020015161180e9190612254565b81525085858151811061182357611823612267565b6020908102919091010152611839600185612254565b93506118458183612254565b61184f9084612254565b92505050611736565b508152949350505050565b5f805f80845f0151116118b85760405162461bcd60e51b815260206004820152601860248201527f524c50206974656d2063616e6e6f74206265206e756c6c2e0000000000000000604482015260640161084a565b602084015180515f1a607f81116118da575f60015f9450945094505050611b96565b60b78111611953575f6118ee6080836122f8565b905080875f0151116119425760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420524c502073686f727420737472696e672e00000000000000604482015260640161084a565b6001955093505f9250611b96915050565b60bf8111611a3f575f61196760b7836122f8565b905080875f0151116119bb5760405162461bcd60e51b815260206004820152601f60248201527f496e76616c696420524c50206c6f6e6720737472696e67206c656e6774682e00604482015260640161084a565b600183015160208290036101000a90046119d58183612254565b885111611a245760405162461bcd60e51b815260206004820152601860248201527f496e76616c696420524c50206c6f6e6720737472696e672e0000000000000000604482015260640161084a565b611a2f826001612254565b965094505f9350611b9692505050565b60f78111611ab8575f611a5360c0836122f8565b905080875f015111611aa75760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420524c502073686f7274206c6973742e000000000000000000604482015260640161084a565b600195509350849250611b96915050565b5f611ac460f7836122f8565b905080875f015111611b185760405162461bcd60e51b815260206004820152601d60248201527f496e76616c696420524c50206c6f6e67206c697374206c656e6774682e000000604482015260640161084a565b600183015160208290036101000a9004611b328183612254565b885111611b7a5760405162461bcd60e51b815260206004820152601660248201527524b73b30b634b210292628103637b733903634b9ba1760511b604482015260640161084a565b611b85826001612254565b9650945060019350611b9692505050565b9193909250565b60605f826001600160401b03811115611bb857611bb8611e93565b6040519080825280601f01601f191660200182016040528015611be2576020820181803683370190505b50905080515f03611bf4579050610c39565b848401602082015f5b85811015611c15578281015182820152602001611bfd565b505f6001602087066020036101000a039050808251168119845116178252839450505050509392505050565b606061078882611da1565b606081611c5a81601f612254565b1015611c995760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015260640161084a565b82611ca48382612254565b1015611ce35760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015260640161084a565b611ced8284612254565b84511015611d315760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b604482015260640161084a565b606082158015611d4f5760405191505f825260208201604052610a94565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015611d88578051835260209283019201611d70565b5050858452601f01601f19166040525050949350505050565b606061078882602001515f845f0151611b9d565b5f8083601f840112611dc5575f80fd5b5081356001600160401b03811115611ddb575f80fd5b602083019150836020828501011115611df2575f80fd5b9250929050565b5f805f805f805f6080888a031215611e0f575f80fd5b87356001600160401b0380821115611e25575f80fd5b611e318b838c01611db5565b909950975060208a0135915080821115611e49575f80fd5b611e558b838c01611db5565b909750955060408a0135915080821115611e6d575f80fd5b50611e7a8a828b01611db5565b989b979a50959894979596606090950135949350505050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611ec957611ec9611e93565b60405290565b60405161022081016001600160401b0381118282101715611ec957611ec9611e93565b80356001600160401b0381168114611f08575f80fd5b919050565b80356001600160a01b0381168114611f08575f80fd5b5f82601f830112611f32575f80fd5b6040516101008082018281106001600160401b0382111715611f5657611f56611e93565b60405283018185821115611f68575f80fd5b845b82811015611f82578035825260209182019101611f6a565b509195945050505050565b80356001600160801b0381168114611f08575f80fd5b5f82601f830112611fb2575f80fd5b81356001600160401b0380821115611fcc57611fcc611e93565b604051601f8301601f19908116603f01168101908282118183101715611ff457611ff4611e93565b8160405283815286602085880101111561200c575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f806040838503121561203c575f80fd5b82356001600160401b0380821115612052575f80fd5b8185019150610100808388031215612068575f80fd5b9193506020850135918183111561207d575f80fd5b9185019160408388031215612090575f80fd5b612098611ea7565b6120a184611ef2565b81526020840135838111156120b4575f80fd5b939093019261030084890312156120c9575f80fd5b6120d1611ecf565b84358152602085013560208201526120eb60408601611f0d565b6040820152606085013560608201526080850135608082015260a085013560a082015261211b8960c08701611f23565b60c08201526101c08086013560e08301526101e061213a818801611f8d565b85840152610200945061214e858801611ef2565b6101208401526121616102208801611ef2565b6101408401526121746102408801611ef2565b6101608401526102608701358681111561218c575f80fd5b6121988c828a01611fa3565b610180850152506102808701356101a08401526121b86102a08801611ef2565b828401526102c08701358184015250506102e08501358382015280602083015250809450505050509250929050565b5f602082840312156121f7575f80fd5b5051919050565b5f808335601e19843603018112612213575f80fd5b8301803591506001600160401b0382111561222c575f80fd5b602001915036819003821315611df2575f80fd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561078857610788612240565b634e487b7160e01b5f52603260045260245ffd5b5f8183825b600881101561229f578151835260209283019290910190600101612280565b5050506101008201905092915050565b5f81515f5b818110156122ce57602081850181015186830152016122b4565b505f93019283525090919050565b5f6122f06122ea83866122af565b846122af565b949350505050565b8181038181111561078857610788612240565b60ff818116838216019081111561078857610788612240565b634e487b7160e01b5f52601260045260245ffd5b5f8261234657612346612324565b500490565b5f6001820161235c5761235c612240565b5060010190565b808202811582820484141761078857610788612240565b600181815b808511156123b457815f190482111561239a5761239a612240565b808516156123a757918102915b93841c939080029061237f565b509250929050565b5f826123ca57506001610788565b816123d657505f610788565b81600181146123ec57600281146123f657612412565b6001915050610788565b60ff84111561240757612407612240565b50506001821b610788565b5060208310610133831016604e8410600b8410161715612435575081810a610788565b61243f838361237a565b805f190482111561245257612452612240565b029392505050565b5f610c3983836123bc565b5f8261247357612473612324565b500690565b5f60ff83168061248a5761248a612324565b8060ff84160691505092915050565b60ff828116828216039081111561078857610788612240565b634e487b7160e01b5f52602160045260245ffdfea26469706673582212201028d2977c315b24d4b6ad26b298f93f06aff5e388141b559f9b0f97c7a77a3764736f6c63430008180033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x4f7142ad.
//
// Solidity: function verifyProof(bytes key, bytes value, bytes proof, bytes32 root) pure returns(bool valid)
func (_Contracts *ContractsCaller) VerifyProof(opts *bind.CallOpts, key []byte, value []byte, proof []byte, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "verifyProof", key, value, proof, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x4f7142ad.
//
// Solidity: function verifyProof(bytes key, bytes value, bytes proof, bytes32 root) pure returns(bool valid)
func (_Contracts *ContractsSession) VerifyProof(key []byte, value []byte, proof []byte, root [32]byte) (bool, error) {
	return _Contracts.Contract.VerifyProof(&_Contracts.CallOpts, key, value, proof, root)
}

// VerifyProof is a free data retrieval call binding the contract method 0x4f7142ad.
//
// Solidity: function verifyProof(bytes key, bytes value, bytes proof, bytes32 root) pure returns(bool valid)
func (_Contracts *ContractsCallerSession) VerifyProof(key []byte, value []byte, proof []byte, root [32]byte) (bool, error) {
	return _Contracts.Contract.VerifyProof(&_Contracts.CallOpts, key, value, proof, root)
}

// VerifyProofWithBlockHash is a paid mutator transaction binding the contract method 0x5af3d05c.
//
// Solidity: function verifyProofWithBlockHash((bytes,bytes,bytes,bytes32,bytes,bytes,bytes,bytes32) content, (uint64,(bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32[8],uint256,uint128,uint64,uint64,uint64,bytes,bytes32,uint64,uint256,bytes32)) blockHeader) returns(bool blockHashValid, bool accountProofValid, bool storageProofValid)
func (_Contracts *ContractsTransactor) VerifyProofWithBlockHash(opts *bind.TransactOpts, content ProofsToVerify, blockHeader BlockHeaderToRebuild) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "verifyProofWithBlockHash", content, blockHeader)
}

// VerifyProofWithBlockHash is a paid mutator transaction binding the contract method 0x5af3d05c.
//
// Solidity: function verifyProofWithBlockHash((bytes,bytes,bytes,bytes32,bytes,bytes,bytes,bytes32) content, (uint64,(bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32[8],uint256,uint128,uint64,uint64,uint64,bytes,bytes32,uint64,uint256,bytes32)) blockHeader) returns(bool blockHashValid, bool accountProofValid, bool storageProofValid)
func (_Contracts *ContractsSession) VerifyProofWithBlockHash(content ProofsToVerify, blockHeader BlockHeaderToRebuild) (*types.Transaction, error) {
	return _Contracts.Contract.VerifyProofWithBlockHash(&_Contracts.TransactOpts, content, blockHeader)
}

// VerifyProofWithBlockHash is a paid mutator transaction binding the contract method 0x5af3d05c.
//
// Solidity: function verifyProofWithBlockHash((bytes,bytes,bytes,bytes32,bytes,bytes,bytes,bytes32) content, (uint64,(bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes32[8],uint256,uint128,uint64,uint64,uint64,bytes,bytes32,uint64,uint256,bytes32)) blockHeader) returns(bool blockHashValid, bool accountProofValid, bool storageProofValid)
func (_Contracts *ContractsTransactorSession) VerifyProofWithBlockHash(content ProofsToVerify, blockHeader BlockHeaderToRebuild) (*types.Transaction, error) {
	return _Contracts.Contract.VerifyProofWithBlockHash(&_Contracts.TransactOpts, content, blockHeader)
}

// ContractsProofValidityIterator is returned from FilterProofValidity and is used to iterate over the raw logs and unpacked data for ProofValidity events raised by the Contracts contract.
type ContractsProofValidityIterator struct {
	Event *ContractsProofValidity // Event containing the contract specifics and raw log

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
func (it *ContractsProofValidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProofValidity)
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
		it.Event = new(ContractsProofValidity)
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
func (it *ContractsProofValidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProofValidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProofValidity represents a ProofValidity event raised by the Contracts contract.
type ContractsProofValidity struct {
	BlockHashValid    bool
	AccountProofValid bool
	StorageProofValid bool
	L1BlockHash       [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// String ContractsProofValidity returns a readable string representing the event struct.
//
// Solidity: event ProofValidity(bool blockHashValid, bool accountProofValid, bool storageProofValid, bytes32 l1BlockHash)
func (e *ContractsProofValidity) String() string {
	s := "Event: ContractsProofValidity {\n"

	s += fmt.Sprintf("BlockHashValid: %v,\n", e.BlockHashValid)
	s += fmt.Sprintf("AccountProofValid: %v,\n", e.AccountProofValid)
	s += fmt.Sprintf("StorageProofValid: %v,\n", e.StorageProofValid)
	s += fmt.Sprintf("L1BlockHash: %v,\n", e.L1BlockHash)
	s += "}"

	return s
}

// FilterProofValidity is a free log retrieval operation binding the contract event 0x97c4a4e08295bc7f3760c9c4d51c41205ba626d35940818e7f7fcbe381c5f558.
//
// Solidity: event ProofValidity(bool blockHashValid, bool accountProofValid, bool storageProofValid, bytes32 l1BlockHash)
func (_Contracts *ContractsFilterer) FilterProofValidity(opts *bind.FilterOpts) (*ContractsProofValidityIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProofValidity")
	if err != nil {
		return nil, err
	}
	return &ContractsProofValidityIterator{contract: _Contracts.contract, event: "ProofValidity", logs: logs, sub: sub}, nil
}

// WatchProofValidity is a free log subscription operation binding the contract event 0x97c4a4e08295bc7f3760c9c4d51c41205ba626d35940818e7f7fcbe381c5f558.
//
// Solidity: event ProofValidity(bool blockHashValid, bool accountProofValid, bool storageProofValid, bytes32 l1BlockHash)
func (_Contracts *ContractsFilterer) WatchProofValidity(opts *bind.WatchOpts, sink chan<- *ContractsProofValidity) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProofValidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProofValidity)
				if err := _Contracts.contract.UnpackLog(event, "ProofValidity", log); err != nil {
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

// ParseProofValidity is a log parse operation binding the contract event 0x97c4a4e08295bc7f3760c9c4d51c41205ba626d35940818e7f7fcbe381c5f558.
//
// Solidity: event ProofValidity(bool blockHashValid, bool accountProofValid, bool storageProofValid, bytes32 l1BlockHash)
func (_Contracts *ContractsFilterer) ParseProofValidity(log types.Log) (*ContractsProofValidity, error) {
	event := new(ContractsProofValidity)
	if err := _Contracts.contract.UnpackLog(event, "ProofValidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
