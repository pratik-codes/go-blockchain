package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// whitelisting every origin to connect with the ws request
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan []byte)
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("Number of clients connected: ", len(clients))

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}

	defer ws.Close()

	clients[ws] = true

	// reading messages from the client
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			delete(clients, ws)
			break
		}

		broadcast <- message
	}
}

// this function will handle the messages received from the client
func HandleMessages() {
	log.Println("HandleMessages started", clients, broadcast)

	for {
		msg := <-broadcast

		go ProcessMessage(msg)
	}
}

// this function will process the message received from the client and broadcast it to all the connected miners
// and in the end it will send the response back to the client
func ProcessMessage(msg []byte) {
	// Unmarshal the JSON message
	var jsonMsg map[string]interface{}
	dataUnMarshalErr := json.Unmarshal(msg, &jsonMsg)
	if dataUnMarshalErr != nil {
		log.Println("Error unmarshaling JSON:", dataUnMarshalErr)
	}

	fmt.Println("Broadcasting message in JSON: ", jsonMsg)

	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}
