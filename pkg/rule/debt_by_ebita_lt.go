package rule

import (
	"math"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func debtByEbitaLowerThan(_ stock.Asset, HistoricalData []stock.YearResult) float64 {
	lower := 0
	analised := int(math.Min(float64(len(HistoricalData)), 5))

	for _, val := range HistoricalData[:analised] {
		if val.DebtByEBITDA <= 2 {
			lower++
		}
	}

	return float64(lower) / float64(analised)
}

func DebtByEbitaLowerThan() Rule {
	return NewRule(
		"debt_by_ebita_lt",
		"Debt by EBITA lower than 2%",
		debtByEbitaLowerThan,
	)
}
