package binance

// Function to filter profitable pairs
func filterProfitablePairs(data []TickerData) []TickerData {
	// Define a slice to store profitable pairs
	var profitablePairs []TickerData

	// Iterate through the data and check profitability criteria (e.g., positive price change percent)
	for _, pair := range data {
		priceChangePercent := pair.PriceChangePercent
		if priceChangePercent[0] != '-' {
			// If priceChangePercent is not negative, consider it profitable
			profitablePairs = append(profitablePairs, pair)
		}
	}

	return profitablePairs
}
