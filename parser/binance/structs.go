package binance

// TickerData represents the 24-hour ticker data for a trading pair.
type TickerData struct {
	AskPrice                string  `json:"askPrice"`
	AskQuantity             string  `json:"askQty"`
	BidPrice                string  `json:"bidPrice"`
	BidQuantity             string  `json:"bidQty"`
	CloseTime               float64 `json:"closeTime"`
	TradeCount              int     `json:"count"`
	FirstTradeID            float64 `json:"firstId"`
	HighPrice               string  `json:"highPrice"`
	LastTradeID             float64 `json:"lastId"`
	LastPrice               string  `json:"lastPrice"`
	LastQuantity            string  `json:"lastQty"`
	LowPrice                string  `json:"lowPrice"`
	OpenPrice               string  `json:"openPrice"`
	OpenTime                float64 `json:"openTime"`
	PreviousClosePrice      string  `json:"prevClosePrice"`
	PriceChange             string  `json:"priceChange"`
	PriceChangePercent      string  `json:"priceChangePercent"`
	PriceChangePercentFloat float64
	QuoteVolume             string `json:"quoteVolume"`
	Symbol                  string `json:"symbol"`
	Volume                  string `json:"volume"`
	WeightedAvgPrice        string `json:"weightedAvgPrice"`
}
