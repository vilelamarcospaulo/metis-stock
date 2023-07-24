package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func DebtByEbitaLowerThan() Rule {
	return NewRule(
		"debt_by_ebita_lt",
		"Debt by EBITA lower than 2%",
		func(_ stock.Asset, yearResults YearResultList) float64 {
			return yearResults.OccurrenceRate(func(yr stock.YearResult) bool {
				return yr.DebtByEBITDA <= 2
			})
		},
	)
}
