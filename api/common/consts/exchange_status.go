package consts

// ExchangeStatus ...
type ExchangeStatus string

const (
	// ExchangeStatusOpen ...
	ExchangeStatusOpen = ExchangeStatus("OPEN")
	// ExchangeStatusPreOpen ...
	ExchangeStatusPreOpen = ExchangeStatus("PREOPEN")
	// ExchangeStatusMaintenance ...
	ExchangeStatusMaintenance = ExchangeStatus("MAINTENANCE")
)
