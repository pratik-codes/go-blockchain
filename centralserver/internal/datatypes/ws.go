package datatypes

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
	Id   string
}

type Message struct {
	Client  *Client // The user client who sent the message
	Content []byte  // The message content to broadcast
}

type MinerClient struct {
	Conn *websocket.Conn
	Id   string
}

type UserClient struct {
	Conn *websocket.Conn
	Id   string
}
