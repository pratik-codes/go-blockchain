package cmd

import (
	ws "miners/pkg/websocket"

	log "miners/pkg/logger"
)

func Execute() {
	serverURL := "ws://localhost:8080/ws/miner"
	log := log.NewLogger()
	miner, err := ws.NewMinerClient(serverURL)
	if err != nil {
		log.Fatalf("Failed to connect to central server: %v", err)
	}

	go miner.ListenForMessages()
}
