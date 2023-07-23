package rule

import (
	"math"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func revenueCagrGreaterThan(_ stock.Asset, HistoricalData []stock.YearResult) float64 {
	greater := 0
	analised := int(math.Min(float64(len(HistoricalData)), 5))

	for _, val := range HistoricalData[:analised] {
		if val.RevenueCAGR >= 5 {
			greater++
		}
	}

	return float64(greater) / float64(analised)
}

func RevenueCagrGreaterThan() Rule {
	return NewRule(
		"revenue_cagr_greater_than",
		"Revenue CAGR greater than 5%",
		revenueCagrGreaterThan,
	)
}
