package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func parseResponse(res *http.Response) (*[]TickerData, error) {
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response status: %s", res.Status)
	}

	var bybitResponse Response
	if err := json.NewDecoder(res.Body).Decode(&bybitResponse); err != nil {
		return nil, fmt.Errorf("decoding response failed: %v", err)
	}

	var data []TickerData
	for _, i := range bybitResponse.Result.List {
		if pct, err := strconv.ParseFloat(i.Price24hPcnt, 64); err == nil {
			i.Price24hPcntFloat = pct
		}
		data = append(data, i)
	}
	return &data, nil
}

// IsValidMarket checks if the provided market is valid.
func IsValidMarket(m Market) bool {
	switch m {
	case Spot, Linear, Option, Inverse:
		return true
	default:
		return false
	}
}
