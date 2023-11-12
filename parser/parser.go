package parser

import (
	"fmt"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser/binance"
)

type Parser interface {
	Binance() *binance.Binance
}

type parserImp struct {
	binanceKey    string
	binanceSecret string
}

func NewBinance(apiKey, apiSecret string) *binance.Binance {
	return binance.NewClient(apiKey, apiSecret)
}

func (p *parserImp) Binance() *binance.Binance {
	return NewBinance(p.binanceKey, p.binanceSecret)
}

func New(config Config) (Parser, error) {
	if config.Binance == nil {
		return nil, fmt.Errorf("binance API key and secret key are required")
	}

	var parser = new(parserImp)
	if config.Binance != nil {
		parser.binanceKey = config.Binance.ApiKey
		parser.binanceSecret = config.Binance.SecretKey
	}
	return parser, nil
}
