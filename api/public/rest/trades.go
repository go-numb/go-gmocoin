package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/api"
	"net/url"
	"strconv"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/public/rest/model"
)

// Trades ...
type Trades interface {
	Trades(symbol configuration.Symbol, page, count int64) (*model.TradesRes, error)
}

func newTrades() trades {
	return trades{
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type trades struct {
	api.RestAPIBase
}

func (t *trades) Trades(symbol configuration.Symbol, page, count int64) (*model.TradesRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	if page > 0 {
		param.Set("page", strconv.FormatInt(page, 10))
	}

	if count > 0 {
		param.Set("count", strconv.FormatInt(count, 10))
	}

	res, err := t.Get(param, "/v1/trades")
	if err != nil {
		return nil, err
	}

	tradesRes := new(model.TradesRes)
	err = json.Unmarshal(res, tradesRes)
	if err != nil {
		return nil, fmt.Errorf("[Trades]error:%v,body:%s", err, res)
	}
	if len(tradesRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", tradesRes.Messages)
	}

	return tradesRes, nil
}
