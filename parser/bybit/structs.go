package bybit

type Response struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     Result                 `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type Result struct {
	Category string       `json:"category"`
	List     []TickerData `json:"list"`
}

type TickerData struct {
	Symbol            string  `json:"symbol"`
	Bid1Price         string  `json:"bid1Price"`
	Bid1Size          string  `json:"bid1Size"`
	Ask1Price         string  `json:"ask1Price"`
	Ask1Size          string  `json:"ask1Size"`
	LastPrice         string  `json:"lastPrice"`
	PrevPrice24h      string  `json:"prevPrice24h"`
	Price24hPcnt      string  `json:"price24hPcnt"`
	HighPrice24h      string  `json:"highPrice24h"`
	LowPrice24h       string  `json:"lowPrice24h"`
	Turnover24h       string  `json:"turnover24h"`
	Volume24h         string  `json:"volume24h"`
	UsdIndexPrice     string  `json:"usdIndexPrice"`
	Price24hPcntFloat float64 `json:"price_24_h_pcnt_float"`
}
