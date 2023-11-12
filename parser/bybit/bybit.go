package bybit

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

const baseURL = "https://api.bybit.com"

type Market string

const (
	Spot    Market = "spot"
	Linear  Market = "linear"
	Option  Market = "option"
	Inverse Market = "inverse"
)

type PairListResponse struct {
	Pairs         []string `json:"pairs"`
	RefreshPeriod int      `json:"refresh_period"`
}

// Client is a struct representing the Client API client.
type Client struct {
	apiKey    string
	apiSecret string
	client    *http.Client
}

// NewClient creates a new instance of the Client API client.
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		client:    &http.Client{},
	}
}

// Get24HourTickerData gets tickers from Bybit API for a given market.
func (c *Client) Get24HourTickerData(market Market) (*[]TickerData, error) {
	if !IsValidMarket(market) {
		return nil, fmt.Errorf("invalid market type: %s", market)
	}

	url := fmt.Sprintf("%s/v5/market/tickers?category=%s", baseURL, market)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %v", err)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request failed: %v", err)
	}
	defer res.Body.Close()

	return parseResponse(res)
}

// Get24HourTickerDataSymbol gets tickers from Bybit API for a given market.
func (c *Client) Get24HourTickerDataSymbol(market Market, symbol string) (*[]TickerData, error) {
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

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request failed: %v", err)
	}
	defer res.Body.Close()

	return parseResponse(res)
}

// Get24HourGainersTickerData returns all trading pairs with a positive price change percent
// of more than +2% over the last 24 hours, sorted by performance (descending order).
func (c *Client) Get24HourGainersTickerData(market Market, limit int, endingFilter string) ([]TickerData, error) {
	resp, err := c.Get24HourTickerData(market)
	if err != nil {
		return nil, err
	}

	// Filter profitable pairs.
	gainerPairs := filterProfitablePairs(*resp)

	// Apply the ending filter.
	filteredPairs := filterPairsByEnding(gainerPairs, endingFilter)

	// Sort the pairs by performance.
	sortPairsByPerformance(filteredPairs)

	// Apply the limit if specified
	if limit > 0 && limit < len(filteredPairs) {
		filteredPairs = filteredPairs[:limit]
	}

	return filteredPairs, nil
}

// GetTickersGainerForPairs returns formatted trading pair symbols as strings.
func (c *Client) GetTickersGainerForPairs(market Market, limit int, endingFilter, excludeFilter string) (PairListResponse, error) {
	resp, err := c.Get24HourTickerData(market)
	if err != nil {
		return PairListResponse{}, err
	}

	// Filter profitable pairs.
	gainerPairs := filterProfitablePairs(*resp)

	// Apply the ending filter.
	filteredPairs := filterPairsByEnding(gainerPairs, endingFilter)

	// Apply the exclude filter.
	finalPairs := make([]TickerData, 0)
	for _, pair := range filteredPairs {
		if !strings.Contains(pair.Symbol, excludeFilter) {
			finalPairs = append(finalPairs, pair)
		}
	}

	// Sort the pairs by performance.
	sortPairsByPerformance(finalPairs)

	// Extract the symbols and apply the limit.
	tradingPairSymbols := extractTradingPairSymbols(finalPairs, limit)

	// Format the pairs.
	formattedPairs := formatPairs(tradingPairSymbols, endingFilter)
	response := PairListResponse{
		Pairs:         formattedPairs,
		RefreshPeriod: 43200, //12h
	}
	return response, nil
}

// formatPairs takes a slice of symbols and appends a "/" between the base currency and the endingFilter.
func formatPairs(symbols []string, endingFilter string) []string {
	var formattedPairs []string
	for _, symbol := range symbols {
		// Trim the endingFilter from the symbol and add it back with a "/" for formatting.
		base := strings.TrimSuffix(symbol, endingFilter)
		formattedPair := base + "/" + endingFilter
		formattedPairs = append(formattedPairs, formattedPair)
	}
	return formattedPairs
}

// FilterProfitablePairs filters out pairs that have a positive price change percent.
func filterProfitablePairs(data []TickerData) []TickerData {
	var profitablePairs []TickerData

	for _, pair := range data {
		if pair.Price24hPcntFloat > 0 {
			profitablePairs = append(profitablePairs, pair)
		}
	}

	return profitablePairs
}

// FilterPairsByEnding filters pairs by ending if endingFilter is provided.
func filterPairsByEnding(data []TickerData, endingFilter string) []TickerData {
	var filteredPairs []TickerData

	for _, pair := range data {
		if strings.HasSuffix(pair.Symbol, endingFilter) {
			filteredPairs = append(filteredPairs, pair)
		}
	}

	return filteredPairs
}

// SortPairsByPerformance sorts the pairs based on their price change percent in descending order.
func sortPairsByPerformance(data []TickerData) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Price24hPcntFloat > data[j].Price24hPcntFloat
	})
}
func extractTradingPairSymbols(data []TickerData, limit int) []string {
	var symbols []string
	for i, pair := range data {
		if limit > 0 && i >= limit {
			break
		}
		symbols = append(symbols, pair.Symbol)
	}
	return symbols
}
