package state

import "centralserver/internal/datatypes"

type BlockChainState struct {
	Blockchain []datatypes.Block
	Balances   map[string]int
	Mempool    []datatypes.Transaction
}

var (
	Blockchain = make([]datatypes.Block, 0)
	Balances   = make(map[string]int)
	MemPool    = make([]datatypes.Transaction, 0)
)

func NewState() *BlockChainState {
	return &BlockChainState{
		Blockchain: Blockchain,
		Balances:   Balances,
		Mempool:    MemPool,
	}
}
