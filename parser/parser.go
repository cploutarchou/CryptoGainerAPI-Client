package parser

import (
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser/binance"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser/bybit"
)

type Parser interface {
	Binance() *binance.Client
	Bybit() *bybit.Client
}

type parserImp struct {
	binanceKey    string
	binanceSecret string
	bybitKey      string
	bybitSecret   string
}

func NewBinance(apiKey, apiSecret string) *binance.Client {
	return binance.NewClient(apiKey, apiSecret)
}
func NewBybit(apiKey, apiSecret string) *bybit.Client {
	return bybit.NewClient(apiKey, apiSecret)
}

func (p *parserImp) Binance() *binance.Client {
	return NewBinance(p.binanceKey, p.binanceSecret)
}
func (p *parserImp) Bybit() *bybit.Client {
	return NewBybit(p.binanceKey, p.binanceSecret)
}

func New(config Config) (Parser, error) {
	var parser = new(parserImp)
	if config.Binance != nil {
		parser.binanceKey = config.Binance.ApiKey
		parser.binanceSecret = config.Binance.SecretKey
	}
	if config.Bybit != nil {
		parser.bybitKey = config.Bybit.ApiKey
		parser.bybitSecret = config.Bybit.SecretKey
	}
	return parser, nil
}
