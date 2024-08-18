package dataypes

// difficulty for mining and proof of work
const MINE_DIFFICULTY = 2

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

type Blockchain struct {
	Blocks []*Block
}
