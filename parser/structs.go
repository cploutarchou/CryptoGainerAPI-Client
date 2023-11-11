package parser

type Config struct {
	Binance *Binance
}

type Binance struct {
	ApiKey    string
	SecretKey string
}
