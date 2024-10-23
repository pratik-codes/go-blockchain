package operations

import (
	"centralserver/internal/datatypes"
	log "centralserver/pkg/logger"
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type Operations struct {
	UserClients  map[string]*datatypes.Client // Track user clients
	MinerClients map[string]*datatypes.Client // Track miner clients
	Broadcast    chan *datatypes.Message      // Channel for broadcasting messages to miners
	Mu           sync.Mutex                   // Protect the clients map
	Miner        *datatypes.Client
	log          *log.Logger
}

// checks if a transaciton sent to the central server is valid
func (o *Operations) CheckValidTransaction() (bool, error) {
	return true, nil
}

// handles the logic to brodcast a valid transactions to all mineres
func (o *Operations) HandleMinersBrodcast(msg *datatypes.Message) {
	o.log.Info("Broadcasting message to miners:", string(msg.Content))

	msgData := map[string]interface{}{
		"message":  msg.Content,
		"clientID": msg.Client.Id,
	}

	msgDataBytes, marshalErr := json.Marshal(msgData)
	if marshalErr != nil {
		o.log.Error("Error marshalling message data: %v", marshalErr)
		return
	}

	err := o.Miner.Conn.WriteMessage(websocket.TextMessage, msgDataBytes)
	if err != nil {
		o.log.Error("Error sending message to miner: %v", err)
		o.Miner.Conn.Close()
		delete(o.MinerClients, msg.Client.Id)
	}
}
