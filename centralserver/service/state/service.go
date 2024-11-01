package state

import (
	log "centralserver/pkg/logger"

	"centralserver/internal/constants"
	"centralserver/internal/datatypes"
	"centralserver/service/operations"
)

type BlockChainState struct {
	Blockchain []*datatypes.Block
	Balances   map[string]int
}

type StateService struct {
	BlockChainState *BlockChainState
	log             *log.Logger
}

var (
	Blockchain     = make([]*datatypes.Block, 0)
	Balances       = make(map[string]int)
	GenisisWallets = make([]*datatypes.Wallet, 0)
)

func NewState() *StateService {
	blockChainState := &BlockChainState{
		Blockchain: Blockchain,
		Balances:   Balances,
	}

	s := &StateService{
		BlockChainState: blockChainState,
		log:             log.NewLogger(),
	}

	// TODO: decide if we want to keep this here
	s.BlockChainGenisis()

	return s
}

func (s *StateService) BlockChainGenisis() *datatypes.Block {
	// genesis block
	genBlock := &datatypes.Block{
		Index:        0,
		Timestamp:    0,
		Transactions: nil,
		PreviousHash: "0",
	}

	genBlock.Hash = operations.CreateBlockHash(genBlock)

	AddBlock(genBlock)

	// create genesis wallets
	s.createGenesisWallets()

	return nil
}

func AddBlock(block *datatypes.Block) {
	Blockchain = append(Blockchain, block)
}

func GetBlockChain() []*datatypes.Block {
	return Blockchain
}

// private functions

// createGenesisWallets creates the initial wallets for the blockchain with minted token
// NOTE: In ideal scenario this should not be done in this way
func (s *StateService) createGenesisWallets() error {
	for i := 0; i < 4; i++ {
		wallet, err := operations.CreateWallet()
		if err != nil {
			s.log.Error("Error creating wallet while genisis: %v", err)
			return err
		} else {
			wallet.Amount = constants.GENISIS_WALLET_AMOUNT
			GenisisWallets = append(GenisisWallets, wallet)
			Balances[wallet.PublicKey] = constants.GENISIS_WALLET_AMOUNT
		}
	}

	return nil
}
