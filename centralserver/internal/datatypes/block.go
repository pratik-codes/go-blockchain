package datatypes

type Transaction struct {
	PublicKey string
	Recipient string
	Amount    int
}

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	Nonce        int
	Hash         string
	PreviousHash string
}

type Wallet struct {
	PublicKey  string
	PrivateKey string
	Amount     float32
}
