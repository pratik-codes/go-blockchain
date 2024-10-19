package operations

import (
	"centralserver/internal/datatypes"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Operations struct {
	UserClients  map[*websocket.Conn]bool // Track user clients
	MinerClients map[*websocket.Conn]bool // Track miner clients
	Broadcast    chan datatypes.Message   // Channel for broadcasting messages to miners
	Mu           sync.Mutex               // Protect the clients map
	Miner        *websocket.Conn
}

func (o *Operations) HandleMinersBrodcast(msg datatypes.Message) {
	log.Println("Broadcasting message to miners:", string(msg.Content))

	err := o.Miner.WriteMessage(websocket.TextMessage, msg.Content)
	if err != nil {
		log.Printf("Error sending message to miner: %v", err)
		o.Miner.Close()
		delete(o.MinerClients, o.Miner)
	}
}
