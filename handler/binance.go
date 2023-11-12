package handler

import (
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser/binance"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PairListResponse struct {
	Pairs         []string `json:"pairs"`
	RefreshPeriod int      `json:"refresh_period"`
}
type Binance interface {
	Get24HourTickerData(c *gin.Context)
	GetTickerForPair(c *gin.Context)
	Get24HourGainersTickerData(c *gin.Context)
	Get24HourGainersPairs(c *gin.Context)
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
//	@tags			Client
//	@Success		200	{array}	TickerData
//	@Router			/binance/ticker/24hr [get]
func (h *BinanceImpl) Get24HourTickerData(c *gin.Context) {
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
//	@tags			Client
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

// Get24HourGainersTickerData
//
//	@Summary		Get the top gainers with a specified limit, filtered by ending.
//	@Description	Retrieve the top gainers with a specified limit, filtered by ending.
//	@Produce		json
//	@Tags			Client
//	@Param			limit			query	int		false	"Limit the number of results"
//	@Param			endingFilter	query	string	false	"Filter results by ending"
//	@Success		200				{array}	TickerData
//	@Router			/binance/ticker/24hr/gainers [get]
func (h *BinanceImpl) Get24HourGainersTickerData(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "500"))
	endingFilter := c.DefaultQuery("endingFilter", "")

	ticker, err := h.parser.Binance().Get24HourGainersTickerData(limit, endingFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticker)
}

// Get24HourGainersPairs
//
//	@Summary		Get the top gainers with a specified limit, filtered by ending and exclusion.
//	@Description	Retrieve the top gainers with a specified limit, filtered by ending and exclusion.
//	@Produce		json
//	@Tags			Client
//	@Param			limit			query	int		false	"Limit the number of results"
//	@Param			endingFilter	query	string	false	"Filter results by ending"
//	@Param			exclude			query	string	false	"Exclude results with specific ending"
//	@Success		200				{array}	PairListResponse
//	@Router			/binance/ticker/24hr/gainers/pairs [get]
func (h *BinanceImpl) Get24HourGainersPairs(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	endingFilter := c.DefaultQuery("endingFilter", "USDT")
	excludeFilter := c.DefaultQuery("exclude", "BNB")

	ticker, err := h.parser.Binance().GetTickersGainerForPairs(limit, endingFilter, excludeFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticker)
}

func NewBinance(parser2 parser.Parser) *BinanceImpl {
	return &BinanceImpl{parser: parser2}
}
