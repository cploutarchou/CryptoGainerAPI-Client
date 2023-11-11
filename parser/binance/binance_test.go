package binance

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGet24HourTickerDataForPair(t *testing.T) {
	// Create a mock HTTP server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define a sample JSON response for the 24-hour ticker
		jsonResponse := `{
			"symbol": "BTCUSDT",
			"priceChangePercent": "2.5",
			"askPrice": "0.00000000",
			"askQty": "0.00000000",
			"bidPrice": "0.00000000",
			"bidQty": "0.00000000",
			"closeTime": 0,
			"count": 0,
			"firstId": 0,
			"highPrice": "0.00000000",
			"lastId": 0,
			"lastPrice": "0.00000000",
			"lastQty": "0.00000000",
			"lowPrice": "0.00000000",
			"openPrice": "0.00000000",
			"openTime": 0,
			"prevClosePrice": "0.00000000",
			"priceChange": "0.00000000",
			"quoteVolume": "0.00000000",
			"volume": "0.00000000",
			"weightedAvgPrice": "0.00000000"
		}`

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Write the JSON response to the client
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	key := os.Getenv("BINANCE_KEY")
	secret := os.Getenv("BINANCE_SECRET")
	// Create a Binance instance using the mock server's URL
	binanceClient := New(key, secret)
	binanceClient.client = server.Client() // Set the client to the mock server's client

	// Make the API call to the mock server and retrieve BinanceTickerData
	tickerData, err := binanceClient.GetTickerForPair("BTCUSDT")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check the response data if needed
	if tickerData.Symbol != "BTCUSDT" {
		t.Errorf("Expected symbol 'BTCUSDT', but got %s", tickerData.Symbol)
	}

}

func TestGet24HourTickerData(t *testing.T) {
	// Create a mock HTTP server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define a sample JSON response for the 24-hour ticker
		jsonResponse := `{
			"symbol": "BTCUSDT",
			"priceChangePercent": "2.5",
			"askPrice": "0.00000000",
			"askQty": "0.00000000",
			"bidPrice": "0.00000000",
			"bidQty": "0.00000000",
			"closeTime": 0,
			"count": 0,
			"firstId": 0,
			"highPrice": "0.00000000",
			"lastId": 0,
			"lastPrice": "0.00000000",
			"lastQty": "0.00000000",
			"lowPrice": "0.00000000",
			"openPrice": "0.00000000",
			"openTime": 0,
			"prevClosePrice": "0.00000000",
			"priceChange": "0.00000000",
			"quoteVolume": "0.00000000",
			"volume": "0.00000000",
			"weightedAvgPrice": "0.00000000"
		}`

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Write the JSON response to the client
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	key := os.Getenv("BINANCE_KEY")
	secret := os.Getenv("BINANCE_SECRET")
	// Create a Binance instance using the mock server's URL
	binanceClient := New(key, secret)
	binanceClient.client = server.Client() // Set the client to the mock server's client

	// Make the API call to the mock server and retrieve BinanceTickerData
	tickerData, err := binanceClient.Get24HourTickerData()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if len(tickerData) < 1000 {
		t.Errorf("Expecter more than 1000 pair's")
	}

}
