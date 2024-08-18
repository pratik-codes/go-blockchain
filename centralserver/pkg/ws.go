package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	// reading messages from the client
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			delete(clients, ws)
			break
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}

		broadcast <- message
	}
}

func HandleMessages() {
	log.Println("HandleMessages started", clients, broadcast)

	for {
		msg := <-broadcast

		// Unmarshal the JSON message
		var jsonMsg map[string]interface{}
		dataUnMarshalErr := json.Unmarshal(msg, &jsonMsg)
		if dataUnMarshalErr != nil {
			log.Println("Error unmarshaling JSON:", dataUnMarshalErr)
			continue
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
}
