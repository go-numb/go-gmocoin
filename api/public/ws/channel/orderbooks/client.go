package orderbooks

import (
	"encoding/json"
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"log"

	"github.com/ijufumi/gogmocoin/api/public/ws/internal/connect"
)

// Client ...
type Client interface {
	Subscribe() error
	Unsubscribe() error
	Receive() <-chan *Response
}

type client struct {
	conn *connect.Connection
}

// New ...
func New(symbol consts.Symbol) Client {
	conn := connect.New()
	conn.SetSubscribeFunc(func() interface{} {
		return Request{
			Command: consts.WebSocketCommandSubscribe,
			Channel: consts.WebSocketChannelOrderBooks,
			Symbol:  symbol,
		}
	})
	conn.SetUnsubscribeFunc(func() interface{} {
		return Request{
			Command: consts.WebSocketCommandUnsubscribe,
			Channel: consts.WebSocketChannelOrderBooks,
			Symbol:  symbol,
		}
	})
	c := &client{
		conn: conn,
	}
	return c
}

func (c *client) Subscribe() error {
	return c.conn.Subscribe()
}

func (c *client) Unsubscribe() error {
	return c.conn.Unsubscribe()
}

func (c *client) Receive() <-chan *Response {
	stream := make(chan *Response, 10)
	go func() {
		for {
			v := <-c.conn.Stream()
			if v == nil {
				return
			}
			log.Printf("received:%v", string(v))
			res := new(Response)
			err := json.Unmarshal(v, res)
			if err != nil {
				log.Printf("[Ticker] unmarshal error:%v", err)
				continue
			}
			stream <- res
		}
	}()
	return stream
}
