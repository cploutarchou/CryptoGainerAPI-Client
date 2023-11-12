package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://api.bybit.com"
)

type Client struct {
	httpClient *http.Client
}

func NewBybitClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) GetBybitTickers() (*Response, error) {
	url := fmt.Sprintf("%s/v5/market/tickers?category=spot", baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var bybitResponse Response
	if err := json.NewDecoder(res.Body).Decode(&bybitResponse); err != nil {
		return nil, err
	}

	return &bybitResponse, nil
}
