package parser

import "github.com/cploutarchou/CryptoGainerAPI-Client/parser/binance"

type Parser interface {
	Binance(apiKey, apiSecret string) *binance.Binance
}

type parserImp struct {
}

func NewBinance(apiKey, apiSecret string) *binance.Binance {
	return binance.New(apiKey, apiSecret)

}

func (p *parserImp) Binance(apiKey, apiSecret string) *binance.Binance {
	return NewBinance(apiKey, apiSecret)
}
