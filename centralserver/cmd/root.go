package cmd

import (
	ws "centralserver/pkg/websocket"
	"log"
	"net/http"
)

func Execute() {
	wsServer := ws.NewWebSocketServer()

	http.HandleFunc("/ws/user", wsServer.HandleUserConnections)
	http.HandleFunc("/ws/miner", wsServer.HandleMinerConnections)

	go wsServer.HandleMessages()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
