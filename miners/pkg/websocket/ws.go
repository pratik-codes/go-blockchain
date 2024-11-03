package ws

import (
	log "miners/pkg/logger"

	"github.com/gorilla/websocket"
)

type MinerClient struct {
	conn *websocket.Conn
	log  *log.Logger
}

func NewMinerClient(serverURL string) (*MinerClient, error) {
	// Connect to the central WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	log := log.NewLogger()
	if err != nil {
		return nil, err
	}
	return &MinerClient{
		conn: conn,
		log:  log,
	}, nil
}

// Listen for incoming messages (transactions or blocks) from the central server
func (m *MinerClient) ListenForMessages() {
	defer m.conn.Close()
	for {
		_, message, err := m.conn.ReadMessage()
		if err != nil {
			m.log.Error("Error reading message: %s", err)
			return
		}

		m.log.Info("Received message: %s", message)

		// Deserialize and handle the received block
		// receivedBlock := DeserializeBlock(message)
		// if receivedBlock != nil {
		// 	// Validate and add to the blockchain if valid
		// 	m.handleReceivedBlock(receivedBlock)
		// }
	}
}

//
// // Handle received block from other miners or transactions from users
// func (m *MinerClient) handleReceivedBlock(block *Block) {
// 	// Example blockchain validation (customize with specific logic)
// 	if blockchain.IsValid() && block.Index > blockchain.Blocks[len(blockchain.Blocks)-1].Index {
// 		blockchain.Blocks = append(blockchain.Blocks, block)
// 	}
// }
