package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"github.com/shopspring/decimal"
)

// ChangeLosscutPrice ...
type ChangeLosscutPrice interface {
	ChangeLosscutPrice(positionID int64, losscutPrice decimal.Decimal) (*model.ChangeLosscutPriceRes, error)
}

func newChangeLosscutPrice(apiKey, secretKey string) changeLosscutPrice {
	return changeLosscutPrice{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type changeLosscutPrice struct {
	api.RestAPIBase
}

func (c *changeLosscutPrice) ChangeLosscutPrice(positionID int64, losscutPrice decimal.Decimal) (*model.ChangeLosscutPriceRes, error) {
	request := model.ChangeLosscutPriceReq{
		PositionID:   positionID,
		LosscutPrice: losscutPrice,
	}

	res, err := c.Post(request, "/v1/changeLosscutPrice")
	if err != nil {
		return nil, err
	}

	response := new(model.ChangeLosscutPriceRes)
	err = json.Unmarshal(res, response)
	if err != nil {
		return nil, fmt.Errorf("[ChangeLosscutPrice]error:%v,body:%s", err, res)
	}

	if len(response.Messages) != 0 {
		return nil, fmt.Errorf("%v", response.Messages)
	}
	return response, nil
}
