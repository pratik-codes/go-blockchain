package cmd

import (
	log "centralserver/pkg/logger"
	ws "centralserver/pkg/websocket"
	"flag"
	"net/http"
)

func Execute() {
	// take port number as input flag to run the server on
	port := flag.String("port", "8080", "Port on which the server will run")
	flag.Parse() // Parse the command-line flags

	wsServer := ws.NewWebSocketServer()

	// handling the websocket connections
	http.HandleFunc("/ws/user", wsServer.HandleUserConnections)
	http.HandleFunc("/ws/miner", wsServer.HandleMinerConnections)

	// a function that runs and always consumes new messages from a user client
	go wsServer.HandleUserMessages()

	log := log.NewLogger()

	log.Info("Starting server on port :%s", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Error("Error starting server: %s", err)
	}
}
