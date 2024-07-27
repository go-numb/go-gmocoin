package configuration

// WebSocketChannel ...
type WebSocketChannel string

const (
	// WebSocketChannelTicker ...
	WebSocketChannelTicker = WebSocketChannel("ticker")
	// WebSocketChannelOrderBooks ...
	WebSocketChannelOrderBooks = WebSocketChannel("orderbooks")
	// WebSocketChannelTrades ...
	WebSocketChannelTrades = WebSocketChannel("trades")
)
