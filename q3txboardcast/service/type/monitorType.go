package service_type

type MonitorResult struct {
	Result string
	Err    error
}

type MonitorStatusOutput struct {
	TxStatus string `json:"tx_status"`
}
