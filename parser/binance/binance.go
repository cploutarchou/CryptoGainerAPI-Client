package binance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	baseURL = "https://api.binance.com/api/v3"
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

// Get24HourTickerData returns 24-hour price statistics mapped to TickerData for all trading pairs.
func (c *Client) Get24HourTickerData() ([]TickerData, error) {
	url := fmt.Sprintf("%s/ticker/24hr", baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-MBX-APIKEY", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	var data []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var tickerData []TickerData
	for _, item := range data {
		var ticker TickerData
		err := mapToBinanceTickerData(item, &ticker)
		if err != nil {
			return nil, err
		}
		tickerData = append(tickerData, ticker)
	}

	return tickerData, nil
}

// GetTickerForPair returns 24-hour price statistics for a specific trading pair.
func (c *Client) GetTickerForPair(pairSymbol string) (TickerData, error) {
	url := fmt.Sprintf("%s/ticker/24hr?symbol=%s", baseURL, pairSymbol)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TickerData{}, err
	}

	req.Header.Set("X-MBX-APIKEY", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return TickerData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return TickerData{}, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return TickerData{}, err
	}

	var ticker TickerData
	err = mapToBinanceTickerData(data, &ticker)
	if err != nil {
		return TickerData{}, err
	}

	return ticker, nil
}

// GetTickersForPairs returns 24-hour price statistics for specific trading pairs.
func (c *Client) GetTickersForPairs(pairSymbols []string) ([]TickerData, error) {
	var tickerDataSlice []TickerData

	for _, pairSymbol := range pairSymbols {
		url := fmt.Sprintf("%s/ticker/24hr?symbol=%s", baseURL, pairSymbol)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("X-MBX-APIKEY", c.apiKey)

		resp, err := c.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("HTTP error: %s", resp.Status)
		}

		var data map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, err
		}

		var ticker TickerData
		err = mapToBinanceTickerData(data, &ticker)
		if err != nil {
			return nil, err
		}

		tickerDataSlice = append(tickerDataSlice, ticker)
	}

	return tickerDataSlice, nil
}

// Get24HourGainersTickerData returns all trading pairs with a positive price change percent
// of more than +2% over the last 24 hours, sorted by performance (descending order).
func (b *Client) Get24HourGainersTickerData(limit int, endingFilter string) ([]TickerData, error) {
	allTickers, err := b.Get24HourTickerData()
	if err != nil {
		return nil, err
	}

	// Filter profitable pairs.
	gainerPairs := filterProfitablePairs(allTickers)

	// Apply the ending filter.
	filteredPairs := filterPairsByEnding(gainerPairs, endingFilter)

	// Apply the exclude filter.
	finalPairs := make([]TickerData, 0)
	for _, pair := range filteredPairs {

		finalPairs = append(finalPairs, pair)

	}

	// Sort the pairs by performance.
	sortPairsByPerformance(finalPairs)

	// Apply the limit if specified
	if limit > 0 && limit < len(filteredPairs) {
		filteredPairs = filteredPairs[:limit]
	}

	return filteredPairs, nil
}

// GetTickersGainerForPairs returns formatted trading pair symbols as strings.
func (c *Client) GetTickersGainerForPairs(limit int, endingFilter, excludeFilter string) (PairListResponse, error) {
	// First, fetch all tickers.
	allTickers, err := c.Get24HourTickerData()
	if err != nil {
		return PairListResponse{}, err
	}

	// Filter profitable pairs.
	gainerPairs := filterProfitablePairs(allTickers)

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
		RefreshPeriod: 10800,
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

// Function to filter profitable pairs
func filterProfitablePairs(data []TickerData) []TickerData {
	var profitablePairs []TickerData

	for _, pair := range data {
		priceChangePercent := pair.PriceChangePercent
		if priceChangePercent[0] != '-' {
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

// Updated sorting function to sort based on the price change percent.
func sortPairsByPerformance(data []TickerData) {
	sort.Slice(data, func(i, j int) bool {
		// Convert price change percent to float for sorting.
		percentI, _ := strconv.ParseFloat(strings.TrimSuffix(data[i].PriceChangePercent, "%"), 64)
		percentJ, _ := strconv.ParseFloat(strings.TrimSuffix(data[j].PriceChangePercent, "%"), 64)
		return percentI > percentJ
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

// FilterPairsEndingWith returns pairs that end with the specified ending.
func (c *Client) FilterPairsEndingWith(ending string) ([]string, error) {
	tickerData, err := c.Get24HourTickerData()
	if err != nil {
		return nil, err
	}

	var filteredPairs []string

	for _, data := range tickerData {
		if strings.HasSuffix(data.Symbol, ending) {
			filteredPairs = append(filteredPairs, data.Symbol)
		}
	}

	return filteredPairs, nil
}

// HasMoreThanTwoConsecutiveZeros checks if a string has more than two consecutive zeros.
func hasMoreThanTwoConsecutiveZeros(s string) bool {
	count := 0
	for _, char := range s {
		if char == '0' {
			count++
			if count > 2 {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

// mapToBinanceTickerData maps JSON data to TickerData struct.
func mapToBinanceTickerData(data map[string]interface{}, ticker *TickerData) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytesData, ticker)
	if err != nil {
		return err
	}
	return nil
}
