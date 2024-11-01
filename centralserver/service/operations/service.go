package operations

import (
	"centralserver/internal/datatypes"
	log "centralserver/pkg/logger"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Operations struct {
	UserClients  map[string]*datatypes.Client // Track user clients
	MinerClients map[string]*datatypes.Client // Track miner clients
	Broadcast    chan *datatypes.Message      // Channel for broadcasting messages to miners
	Mu           sync.Mutex                   // Protect the clients map
	Miner        *datatypes.Client
	log          *log.Logger
}

// checks if a transaciton sent to the central server is valid
func (o *Operations) CheckValidTransaction() (bool, error) {
	return true, nil
}

// handles the logic to brodcast a valid transactions to all mineres
func (o *Operations) HandleMinersBrodcast(msg *datatypes.Message) {
	o.log.Info("Broadcasting message to miners:", string(msg.Content))

	msgData := map[string]interface{}{
		"message":  msg.Content,
		"clientID": msg.Client.Id,
	}

	msgDataBytes, marshalErr := json.Marshal(msgData)
	if marshalErr != nil {
		o.log.Error("Error marshalling message data: %v", marshalErr)
		return
	}

	err := o.Miner.Conn.WriteMessage(websocket.TextMessage, msgDataBytes)
	if err != nil {
		o.log.Error("Error sending message to miner: %v", err)
		o.Miner.Conn.Close()
		delete(o.MinerClients, msg.Client.Id)
	}
}

// CreateBlockHash creates a SHA-256 hash of a single block
func CreateBlockHash(block *datatypes.Block) string {
	// Concatenate block fields into a single string
	data := strconv.Itoa(block.Index) +
		strconv.FormatInt(block.Timestamp, 10) +
		concatTransactions(block.Transactions) +
		strconv.Itoa(block.Nonce) +
		block.PreviousHash

	// Compute SHA-256 hash
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// concatTransactions concatenates the transactions' relevant fields into a single string
func concatTransactions(transactions []datatypes.Transaction) string {
	var sb strings.Builder
	for _, tx := range transactions {
		sb.WriteString(tx.PublicKey)
		sb.WriteString(tx.Recipient)
		sb.WriteString(strconv.Itoa(tx.Amount))
	}
	return sb.String()
}

// VerifyBlockHash verifies that the stored hash matches the computed hash for a block
func VerifyBlockHash(block *datatypes.Block) bool {
	// Recalculate the hash using the same method as CreateBlockHash
	calculatedHash := CreateBlockHash(block)
	// Compare the recalculated hash to the stored hash
	return calculatedHash == block.Hash
}

// createWallet generates a new key pair, extracts the public and private keys in hexadecimal format, and generates an address from the public key.
func CreateWallet() (*datatypes.Wallet, error) {
	wallet := &datatypes.Wallet{}

	// Generate ECDSA key pair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return wallet, err
	}

	// Encode private key in hexadecimal
	privateKeyHex := hex.EncodeToString(privateKey.D.Bytes())

	// Encode public key in hexadecimal
	publicKeyBytes := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)

	// // Generate address from the public key (this function would need to be implemented)
	// address := PublicKeyToAddress(publicKeyBytes)

	wallet.PrivateKey = privateKeyHex
	wallet.PublicKey = publicKeyHex
	wallet.Amount = 0

	return wallet, nil
}

// // Example function to convert public key to address (you might need to customize this function)
// func PublicKeyToAddress(publicKey []byte) string {
// 	return hex.EncodeToString(publicKey[:20])
// }
