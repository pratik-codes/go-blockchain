package datatypes

import "github.com/gorilla/websocket"

type Message struct {
	UserClient *websocket.Conn // The user client who sent the message
	Content    []byte          // The message content to broadcast
}

type Transaction struct{}

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	nonce        int
	PreviousHash string
	Hash         string
}

const (
	DIFFICULTY         = 4 // requires hash to start with '0000'
	DESIRED_BLOCK_TIME = 30000
)
