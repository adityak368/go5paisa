package go5paisa

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	// "log"
)

type MarketFeedQuery struct {
	Exch        string
	ExchType    string
	Symbol      string
	Expiry      string
	StrikePrice float64
	OptionType  string
}

type MarketFeedData struct {
	Exch     string
	ExchType string
	High     float64
	LastRate float64
	Low      float64
	PClose   float64
	TickDt   string
	Time     int
	Token    float64
	TotalQty float64
}

type MarketFeedResponse struct {
	Data      []MarketFeedData
	Message   string
	Status    int
	CacheTime int
	TimeStamp string
}

// MarketFeed fetches market feed of a scrip
func (c *Client) MarketFeed(marketFeedQuery []MarketFeedQuery) ([]MarketFeedData, error) {
	var marketFeedResponse MarketFeedResponse
	c.appConfig.head.RequestCode = marketFeedRequestCode
	payloadBody := marketFeedPayloadBody{
		MarketFeedData:  marketFeedQuery,
		ClientLoginType: 0,
	}
	payload := marketFeedPayload{
		Head: c.appConfig.head,
		Body: payloadBody,
	}
	jsonValue, _ := json.Marshal(payload)
	res, err := c.connection.Post(baseURL+marketFeedRoute, contentType, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	parseResBody(resBody, &marketFeedResponse)
	if marketFeedResponse.Status != 0 {
		return nil, errors.New(marketFeedResponse.Message)
	}
	return marketFeedResponse.Data, nil
}
