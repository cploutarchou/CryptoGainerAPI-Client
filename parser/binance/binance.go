package binance

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://api.binance.com/api/v3"
)

// Binance is a struct representing the Binance API client.
type Binance struct {
	apiKey    string
	apiSecret string
	client    *http.Client
}

// New creates a new instance of the Binance API client.
func New(apiKey, apiSecret string) *Binance {
	return &Binance{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		client:    &http.Client{},
	}
}

// Get24HourTickerData returns 24-hour price statistics mapped to TickerData for all trading pairs.
func (b *Binance) Get24HourTickerData() ([]TickerData, error) {
	url := fmt.Sprintf("%s/ticker/24hr", baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-MBX-APIKEY", b.apiKey)

	resp, err := b.client.Do(req)
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
func (b *Binance) GetTickerForPair(pairSymbol string) (TickerData, error) {
	url := fmt.Sprintf("%s/ticker/24hr?symbol=%s", baseURL, pairSymbol)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TickerData{}, err
	}

	req.Header.Set("X-MBX-APIKEY", b.apiKey)

	resp, err := b.client.Do(req)
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
