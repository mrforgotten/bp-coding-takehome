package service_type

type BoardcastInput struct {
	Symbol    string
	Price     string
	Timestamp string
}

type BoardcastOutput struct {
	TxHash string `json:"tx_hash"`
}
