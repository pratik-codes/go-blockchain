package ws

import (
	"centralserver/internal/constants"
	"centralserver/internal/datatypes"
	log "centralserver/pkg/logger"
	"centralserver/service/operations"
	"centralserver/utils"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	userClients  map[string]*datatypes.Client // Track user clients
	minerClients map[string]*datatypes.Client // Track miner clients
	memPool      chan *datatypes.Message      // memPool for broadcasting messages to miners
	mu           sync.Mutex                   // Protect the clients map
	Ops          *operations.Operations

	// utitlity
	log *log.Logger
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// NewWebSocketServer initializes a new WebSocket server
func NewWebSocketServer() *WebSocketServer {
	userClients := make(map[string]*datatypes.Client)
	minerClients := make(map[string]*datatypes.Client)
	memPool := make(chan *datatypes.Message)
	o := &operations.Operations{
		UserClients:  userClients,
		MinerClients: minerClients,
		Broadcast:    memPool,
	}
	log := log.NewLogger()

	return &WebSocketServer{
		userClients:  userClients,
		minerClients: minerClients,
		memPool:      memPool,
		Ops:          o,
		log:          log,
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
	s.log.Info("Client connected, type: %s", clientType)

	// Upgrade the HTTP connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.Error("Error upgrading to WebSocket (%s): %v", clientType, err)
		return
	}

	// creating current user connected data
	_uuid := uuid.New().String()
	user := &datatypes.Client{
		Conn: ws,
		Id:   _uuid,
	}

	// defer cleanup for the current user
	defer s.clearClients(ws, clientType, user)

	utils.WithLock(&s.mu, func() {
		if clientType == constants.USER_CLIENT {
			s.userClients[_uuid] = user
			s.log.Info("User connected. Total users: %d", len(s.userClients))
		} else {
			// TODO: fix this when implementing proper miner side logic
			s.minerClients[_uuid] = &datatypes.Client{}
			s.log.Info("Miner connected. Total miners: %d", len(s.minerClients))
		}
	})

	// Handle client messages
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			s.log.Error("Error reading WebSocket message (%s): %v", clientType, err)
			break
		}

		// Handle messages based on client type
		if clientType == constants.USER_CLIENT {
			msg := &datatypes.Message{
				Client:  user,
				Content: message,
			}
			s.memPool <- msg
		}

		if clientType == constants.MINER_CLIENT {
			s.HandleMinersMessage(message)
		}
	}
}

func (s *WebSocketServer) HandleMinersMessage(msg []byte) {
	s.log.Info("message from miner: ", string(msg))
}

// Handle Transactions listens for transactions from the mempool and sends them to all miners
func (s *WebSocketServer) HandleTransactions() {
	s.log.Info("Listening for broadcast messages...")

	for {
		msg := <-s.memPool
		for miner := range s.minerClients {
			s.Ops.Miner = s.minerClients[miner]
			go s.Ops.HandleMinersBrodcast(msg)
		}
	}
}

// Cleanup when a client disconnects
func (s *WebSocketServer) clearClients(ws *websocket.Conn, clientType string, user *datatypes.Client) {
	ws.Close()

	if clientType == "user" {
		uuid := uuid.New()
		user := &datatypes.UserClient{
			Id:   uuid.String(),
			Conn: ws,
		}
		utils.WithLock(&s.mu, func() {
			delete(s.userClients, user.Id)
		})
		s.log.Info("User disconnected. Total users: %d", len(s.userClients))
	} else {
		utils.WithLock(&s.mu, func() {
			delete(s.minerClients, user.Id)
		})
		s.log.Info("Miner disconnected. Total miners: %d", len(s.minerClients))
	}
}
