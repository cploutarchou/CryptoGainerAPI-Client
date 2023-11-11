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

// FilterPairsEndingWith returns pairs that end with the specified ending.
func (b *Binance) FilterPairsEndingWith(ending string) ([]string, error) {
	tickerData, err := b.Get24HourTickerData()
	if err != nil {
		return nil, err
	}

	var filteredPairs []string

	for _, data := range tickerData {
		if len(data.Symbol) >= len(ending) && data.Symbol[len(data.Symbol)-len(ending):] == ending {
			filteredPairs = append(filteredPairs, data.Symbol)
		}
	}

	return filteredPairs, nil
}
