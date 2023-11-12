package bybit

import (
	"fmt"
	"net/http"
)

const baseURL = "https://api.bybit.com"

// Market represents a custom type for market categories.
type Market string

// Constants for a Market type resembling an enum.
const (
	Spot    Market = "spot"
	Linear  Market = "linear"
	Option  Market = "option"
	Inverse Market = "inverse"
)

// Client represents a HTTP client for Bybit API.
type Client struct {
	httpClient *http.Client
}

// NewBybitClient creates a new Bybit API client.
func NewBybitClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{httpClient: httpClient}
}

// Get24HourTickerData gets tickers from Bybit API for a given market.
func (c *Client) Get24HourTickerData(market Market) (*Response, error) {
	if !IsValidMarket(market) {
		return nil, fmt.Errorf("invalid market type: %s", market)
	}

	url := fmt.Sprintf("%s/v5/market/tickers?category=%s", baseURL, market)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request failed: %v", err)
	}
	defer res.Body.Close()

	return parseResponse(res)
}

// Get24HourTickerDataSymbol gets tickers from Bybit API for a given market.
func (c *Client) Get24HourTickerDataSymbol(market Market, symbol string) (*Response, error) {
	if !IsValidMarket(market) {
		return nil, fmt.Errorf("invalid market type: %s", market)
	}

	if len(symbol) < 3 {
		return nil, fmt.Errorf("invalid symbol")
	}
	url := fmt.Sprintf("%s/v5/market/tickers?category=%s&symbol=%s", baseURL, market, symbol)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request failed: %v", err)
	}
	defer res.Body.Close()

	return parseResponse(res)
}
