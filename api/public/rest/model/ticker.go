package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"time"

	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// TickerRes ...
type TickerRes struct {
	model.ResponseCommon
	Data []struct {
		Ask       decimal.Decimal `json:"ask"`
		Bid       decimal.Decimal `json:"bid"`
		High      decimal.Decimal `json:"high"`
		Low       decimal.Decimal `json:"low"`
		Last      decimal.Decimal `json:"last"`
		Symbol    consts.Symbol   `json:"symbol"`
		Timestamp time.Time       `json:"timestamp"`
		Volume    decimal.Decimal `json:"volume"`
	}
}
