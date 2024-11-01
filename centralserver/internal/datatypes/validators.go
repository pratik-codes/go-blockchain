package datatypes

type TransactionPayloadData struct {
	Sender           string `json:"from"`
	Reciever         string `json:"to"`
	Amount           int    `json:"amount"`
	SenderPrivateKey string `json:"privateKey"`
}

type UserTransactionPayload struct {
	Type string                 `json:"type"`
	Data TransactionPayloadData `json:"data"`
}
