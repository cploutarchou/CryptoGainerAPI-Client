package handler

import (
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser"
)

type Handlers interface {
	Binance() Binance
	Bybit() Bybit
}

type HandlersImpl struct {
	parser parser.Parser
}

func New(parser2 parser.Parser) *HandlersImpl {
	return &HandlersImpl{parser: parser2}
}

func (h *HandlersImpl) Binance() Binance {
	return NewBinance(h.parser)
}

func (h *HandlersImpl) Bybit() Bybit {
	return NewBybit(h.parser)
}
