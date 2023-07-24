package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func RevenueCagrGreaterThan() Rule {
	return NewRule(
		"revenue_cagr_greater_than",
		"Revenue CAGR greater than 5%",
		func(_ stock.Asset, yearResults YearResultList) float64 {
			return yearResults.OccurrenceRate(func(yr stock.YearResult) bool {
				return yr.RevenueCAGR >= 5
			})
		},
	)
}
