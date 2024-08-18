package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	operations "miners/service"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var blockchain = &operations.Blockchain{[]*operations.Block{operations.CreateGenesisBlock()}}

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	log.Printf("Connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	go listenForMessages(conn)

	for {
		time.Sleep(10 * time.Second)
		blockchain.AddBlock(fmt.Sprintf("Block data at %s", time.Now().String()))
		sendBlock(conn, blockchain.Blocks[len(blockchain.Blocks)-1])
	}
}

func listenForMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		receivedBlock := DeserializeBlock(message)
		if receivedBlock != nil && receivedBlock.Index > blockchain.Blocks[len(blockchain.Blocks)-1].Index {
			if blockchain.IsValid() {
				blockchain.Blocks = append(blockchain.Blocks, receivedBlock)
			}
		}
	}
}

func sendBlock(conn *websocket.Conn, block *operations.Block) {
	err := conn.WriteMessage(websocket.BinaryMessage, SerializeBlock(block))
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func SerializeBlock(block *operations.Block) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(block)
	if err != nil {
		log.Println("Error encoding block:", err)
		return nil
	}
	return buf.Bytes()
}

func DeserializeBlock(data []byte) *operations.Block {
	var block operations.Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&block)
	if err != nil {
		log.Println("Error decoding block:", err)
		return nil
	}
	return &block
}
