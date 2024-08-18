package main

import (
	ws "centralserver/pkg"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", ws.HandleConnections)
	// handles all the messages that are broadcasted
	go ws.HandleMessages()
	fmt.Println("Central server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
