package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"github.com/shopspring/decimal"
)

// CloseBulkOrder ...
type CloseBulkOrder interface {
	CloseBulkOrder(symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.CloseBulkOrderRes, error)
}

func newCloseBulkOrder(apiKey, secretKey string) closeBulkOrder {
	return closeBulkOrder{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type closeBulkOrder struct {
	api.RestAPIBase
}

// CloseBulkOrder ...
func (c *closeBulkOrder) CloseBulkOrder(symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.CloseBulkOrderRes, error) {
	request := model.CloseBulkOrderReq{
		Symbol:        symbol,
		Side:          side,
		ExecutionType: executionType,
		Size:          size,
	}

	if executionType == consts.ExecutionTypeLIMIT {
		request.Price = &price
	}

	res, err := c.Post(request, "/v1/closeBulkOrder")
	if err != nil {
		return nil, err
	}

	response := new(model.CloseBulkOrderRes)
	err = json.Unmarshal(res, response)
	if err != nil {
		return nil, fmt.Errorf("[CloseBulkOrder]error:%v,body:%s", err, res)
	}

	if !response.Success() {
		return nil, response.Error()
	}
	return response, nil
}
