package handler

import (
	"fmt"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser/binance"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Binance interface {
	Get24HourTickerData(c *gin.Context)
	GetTickerForPair(c *gin.Context)
}

type BinanceImpl struct {
	parser parser.Parser
}

type TickerData []binance.TickerData
type Ticker binance.TickerData

// Get24HourTickerData
//
//	@Summary		Get 24-hour ticker data
//	@Description	Retrieve 24-hour ticker data for all trading pairs.
//	@Produce		json
//	@tags			Binance
//	@Success		200	{array}	TickerData
//	@Router			/binance/ticker/24hr [get]
func (h *BinanceImpl) Get24HourTickerData(c *gin.Context) {
	fmt.Println("TEss")
	tickerData, err := h.parser.Binance().Get24HourTickerData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickerData)
}

// GetTickerForPair
//
//	@Summary		Get ticker data for a specific trading pair
//	@Description	Retrieve ticker data for a specific trading pair.
//	@Produce		json
//	@tags			Binance
//	@Param			pair	path		string	true	"Trading pair symbol (e.g., BTCUSDT)"
//	@Success		200		{object}	Ticker
//	@Router			/binance/ticker/24hr/{pair} [get]
func (h *BinanceImpl) GetTickerForPair(c *gin.Context) {
	pair := c.Param("pair")
	ticker, err := h.parser.Binance().GetTickerForPair(pair)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticker)
}

func NewBinance(parser2 parser.Parser) *BinanceImpl {
	return &BinanceImpl{parser: parser2}
}
