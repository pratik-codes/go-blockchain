package datatypes

type TransactionPayloadData struct {
	Sender           string `json:"sender"`
	Reciever         string `json:"receiver"`
	Amount           int    `json:"amount"`
	SenderPrivateKey string `json:"privateKey"`
}

type UserTransactionPayload struct {
	Type string                 `json:"type"`
	Data TransactionPayloadData `json:"data"`
}
