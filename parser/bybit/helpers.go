package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// parseResponse handles the parsing of the HTTP response.
func parseResponse(res *http.Response) (*Response, error) {
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response status: %s", res.Status)
	}

	var bybitResponse Response
	if err := json.NewDecoder(res.Body).Decode(&bybitResponse); err != nil {
		return nil, fmt.Errorf("decoding response failed: %v", err)
	}

	return &bybitResponse, nil
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
