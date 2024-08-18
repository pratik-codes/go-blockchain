package operations

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	dataypes "miners/internal"
	"time"
)

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

func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp + b.Data + b.PrevHash + string(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) MineBlock(difficulty int) {
	for i := 0; ; i++ {
		b.Nonce = i
		hash := b.CalculateHash()
		if hash[:difficulty] == string(bytes.Repeat([]byte("0"), difficulty)) {
			b.Hash = hash
			break
		}
	}
}

func CreateBlock(prevBlock *Block, data string) *Block {
	block := &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
		Nonce:     0,
	}

	// difficulty 2, adjust as needed
	block.MineBlock(dataypes.MINE_DIFFICULTY)
	return block
}

func CreateGenesisBlock() *Block {
	return CreateBlock(&Block{Index: 0, Timestamp: time.Now().String(), Data: "Genesis Block", PrevHash: "", Hash: ""}, "Genesis Block")
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(prevBlock, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		prevBlock := bc.Blocks[i-1]
		currentBlock := bc.Blocks[i]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}
