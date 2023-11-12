package parser

type Config struct {
	Binance *Binance
	Bybit   *Bybit
}

type Binance struct {
	ApiKey    string
	SecretKey string
}

type Bybit struct {
	ApiKey    string
	SecretKey string
}
