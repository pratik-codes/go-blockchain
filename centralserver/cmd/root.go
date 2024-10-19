package cmd

import (
	ws "centralserver/pkg/websocket"
	"fmt"
	"net/http"
)

func Execute() {
	http.HandleFunc("/ws", ws.HandleConnections)

	// handles all the messages that are broadcasted
	go ws.HandleMessages()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	fmt.Println("Central server started on :8080")
}
