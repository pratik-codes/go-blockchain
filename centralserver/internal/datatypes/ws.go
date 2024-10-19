package datatypes

import "github.com/gorilla/websocket"

type Message struct {
	UserClient *websocket.Conn // The user client who sent the message
	Content    []byte          // The message content to broadcast
}
