package ws

import (
	"centralserver/internal/constants"
	"centralserver/internal/datatypes"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	userClients          map[*websocket.Conn]bool // Track user clients
	minerClients         map[*websocket.Conn]bool // Track miner clients
	broadcast            chan datatypes.Message   // Channel for broadcasting messages to miners
	mu                   sync.Mutex               // Protect the clients map
	handleMinersBrodcast func() (msg datatypes.Message)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// NewWebSocketServer initializes a new WebSocket server
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		userClients:  make(map[*websocket.Conn]bool),
		minerClients: make(map[*websocket.Conn]bool),
		broadcast:    make(chan datatypes.Message),
		// handleMinersBrodcast: operations.HandleMinersBrodcast,
	}
}

// HandleUserConnections upgrades the HTTP connection to a WebSocket for users (to send transactions)
func (s *WebSocketServer) HandleUserConnections(w http.ResponseWriter, r *http.Request) {
	s.handleConnections(w, r, constants.USER_CLIENT)
}

// HandleMinerConnections upgrades the HTTP connection to a WebSocket for miners (to receive transactions, mine blocks)
func (s *WebSocketServer) HandleMinerConnections(w http.ResponseWriter, r *http.Request) {
	s.handleConnections(w, r, constants.MINER_CLIENT)
}

// handleConnections is a helper function to upgrade the connection and manage the client map
func (s *WebSocketServer) handleConnections(w http.ResponseWriter, r *http.Request, clientType string) {
	// Upgrade the HTTP connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket (%s): %v", clientType, err)
		return
	}
	defer ws.Close()

	s.mu.Lock()
	if clientType == constants.USER_CLIENT {
		s.userClients[ws] = true
		log.Printf("User connected. Total users: %d", len(s.userClients))
	} else {
		s.minerClients[ws] = true
		log.Printf("Miner connected. Total miners: %d", len(s.minerClients))
	}
	s.mu.Unlock()

	defer func() {
		// clear clients
		s.clearClients(ws, clientType)
	}()

	// Handle client messages
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading WebSocket message (%s): %v", clientType, err)
			break
		}

		// Handle messages based on client type
		if clientType == constants.USER_CLIENT {
			msg := datatypes.Message{
				// Broadcast the user's transaction to miners
				UserClient: ws,
				Content:    message,
			}
			s.broadcast <- msg
		}

		// TODO: Miner-specific message handling (e.g., mined blocks)
	}
}

// HandleMessages listens for broadcast messages and sends them to all miners
func (s *WebSocketServer) HandleMessages() {
	log.Println("Listening for broadcast messages...")

	for {
		// Lock only when accessing shared resources
		s.mu.Lock()
		for miner := range s.minerClients {
			fmt.Println(miner)
			// o := &operations.Operations{
			// 	Miner:       miner,
			// 	UserClients: s.userClients,
			// 	Broadcast:   s.broadcast,
			// 	Mu:          s.mu,
			// }
			// go o.HandleMinersBroadcast()
		}
		s.mu.Unlock()
	}
}

// Cleanup when a client disconnects
func (s *WebSocketServer) clearClients(ws *websocket.Conn, clientType string) {
	s.mu.Lock()
	if clientType == "user" {
		delete(s.userClients, ws)
		log.Printf("User disconnected. Total users: %d", len(s.userClients))
	} else {
		delete(s.minerClients, ws)
		log.Printf("Miner disconnected. Total miners: %d", len(s.minerClients))
	}
	s.mu.Unlock()
}
