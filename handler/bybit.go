package handler

import (
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser/binance"
	"github.com/cploutarchou/CryptoGainerAPI-Client/parser/bybit"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Bybit interface {
	Get24HourTickerData(c *gin.Context)
	GetTickerForPair(c *gin.Context)
	Get24HourGainersTickerData(c *gin.Context)
	Get24HourGainersPairs(c *gin.Context)
}

type BybitImpl struct {
	parser parser.Parser
}

type BybitTickerData []bybit.TickerData
type BybitTicker binance.TickerData

// Get24HourTickerData retrieves 24-hour ticker data for a specified market type.
//
//	@Summary		Retrieve 24-hour ticker data for the specified market type in Bybit.
//	@Description	This function fetches the 24-hour ticker data for all trading pairs in a given market (Spot, Linear, Option, or Inverse).
//
//	The market type can be specified as a query parameter; if not provided, the default market type is 'spot'.
//
//	@Produce		json
//	@Tags			Bybit
//	@Param			market	query		string			false	"Market type (spot, linear, option, inverse)"	Enums(spot, linear, option, inverse)	default(spot)
//	@Success		200		{object}	[]TickerData	"List of TickerData representing 24-hour ticker information for each trading pair"
//	@Failure		400		"Invalid market type provided"
//	@Failure		500		"Internal Server Error"
//	@Router			/bybit/ticker/24hr [get]
func (h *BybitImpl) Get24HourTickerData(c *gin.Context) {
	market := c.DefaultQuery("market", "spot")
	var validMarket bybit.Market
	switch market {
	case "spot":
		validMarket = bybit.Spot
	case "option":
		validMarket = bybit.Option
	case "linear":
		validMarket = bybit.Linear
	case "inverse":
		validMarket = bybit.Inverse
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid market type"})
		return
	}

	tickerData, err := h.parser.Bybit().Get24HourTickerData(validMarket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickerData)
}

// Get24HourGainersTickerData retrieves a list of top gaining trading pairs over the last 24 hours.
//
//	@Summary		Retrieve top gainers in a specified market with an optional limit and ending filter.
//	@Description	This function fetches trading pairs that have shown significant price gains over the last 24 hours in a specified market.
//
//	It allows filtering by a specific market type and a limit on the number of results. An optional ending filter can also be applied to refine the results.
//
//	@Produce		json
//	@Tags			Bybit
//	@Param			limit			query		int				false	"Limit the number of results; default is 500"	default(500)
//	@Param			endingFilter	query		string			false	"Filter results by a specific ending symbol"
//	@Param			market			query		string			false	"Market type (spot, linear, option, inverse); default is 'spot'"	Enums(spot, linear, option, inverse)	default(spot)
//	@Success		200				{object}	[]TickerData	"List of TickerData representing top gainers"
//	@Failure		400				"Invalid market type or query parameters"
//	@Failure		500				"Internal Server Error"
//	@Router			/bybit/ticker/24hr/gainers [get]
func (h *BybitImpl) Get24HourGainersTickerData(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "500"))
	endingFilter := c.DefaultQuery("endingFilter", "")
	market := c.DefaultQuery("market", "spot")

	var validMarket bybit.Market
	switch market {
	case "spot":
		validMarket = bybit.Spot
	case "option":
		validMarket = bybit.Option
	case "linear":
		validMarket = bybit.Linear
	case "inverse":
		validMarket = bybit.Inverse
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid market type"})
		return
	}

	tickerData, err := h.parser.Bybit().Get24HourGainersTickerData(validMarket, limit, endingFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickerData)
}

// Get24HourGainersPairs retrieves a list of top gaining trading pairs over the last 24 hours.
//
//	@Summary		Retrieve top gainers in a specified market with filtering options.
//	@Description	This function fetches trading pairs that have shown significant price gains over the last 24 hours.
//
//	It allows filtering by a specific market type, a limit on the number of results, and options to include
//	or exclude pairs based on their ending symbol.
//
//	@Produce		json
//	@Tags			Bybit
//	@Param			limit			query		int					false	"Limit the number of results; default is 100"						default(100)
//	@Param			endingFilter	query		string				false	"Filter results by a specific ending symbol; default is 'USDT'"		default(USDT)
//	@Param			exclude			query		string				false	"Exclude results with a specific ending symbol; default is 'BNB'"	default(BNB)
//	@Param			market			query		string				false	"Market type (spot, linear, option, inverse); default is 'spot'"	Enums(spot, linear, option, inverse)	default(spot)
//	@Success		200				{object}	[]PairListResponse	"List of PairListResponse representing top gainers"
//	@Failure		400				"Invalid market type or query parameters"
//	@Failure		500				"Internal Server Error"
//	@Router			/bybit/ticker/24hr/gainers/pairs [get]
func (h *BybitImpl) Get24HourGainersPairs(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	endingFilter := c.DefaultQuery("endingFilter", "USDT")
	excludeFilter := c.DefaultQuery("exclude", "BNB")
	market := c.DefaultQuery("market", "spot")

	var validMarket bybit.Market
	switch market {
	case "spot":
		validMarket = bybit.Spot
	case "option":
		validMarket = bybit.Option
	case "linear":
		validMarket = bybit.Linear
	case "inverse":
		validMarket = bybit.Inverse
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid market type"})
		return
	}

	ticker, err := h.parser.Bybit().GetTickersGainerForPairs(validMarket, limit, endingFilter, excludeFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticker)
}

// GetTickerForPair retrieves ticker data for a specified trading pair in a given market.
//
//	@Summary		Retrieve ticker data for a specific trading pair.
//	@Description	This function fetches the latest ticker information for a specified trading pair in a chosen market.
//
//	It supports different market types like spot, option, linear, and inverse.
//
//	@Produce		json
//	@Tags			Bybit
//	@Param			pair	path		string		true	"Trading pair symbol (e.g., BTCUSDT)"
//	@Param			market	query		string		false	"Market type (spot, linear, option, inverse); default is 'spot'"	Enums(spot, linear, option, inverse)	default(spot)
//	@Success		200		{object}	TickerData	"TickerData object containing detailed ticker information for the specified pair"
//	@Failure		400		"Invalid market type or trading pair symbol"
//	@Failure		500		"Internal Server Error"
//	@Router			/bybit/ticker/24hr/{pair} [get]
func (h *BybitImpl) GetTickerForPair(c *gin.Context) {
	pair := c.Param("pair")
	market := c.DefaultQuery("market", "spot")

	var validMarket bybit.Market
	switch market {
	case "spot":
		validMarket = bybit.Spot
	case "option":
		validMarket = bybit.Option
	case "linear":
		validMarket = bybit.Linear
	case "inverse":
		validMarket = bybit.Inverse
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid market type"})
		return
	}

	tickerData, err := h.parser.Bybit().Get24HourTickerDataSymbol(validMarket, pair)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickerData)
}

func NewBybit(parser2 parser.Parser) *BybitImpl {
	return &BybitImpl{parser: parser2}
}
