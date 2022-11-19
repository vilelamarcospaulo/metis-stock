package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func DebtByEbitda() Rule {
	return newRule("Div. Liquida / EBITDA", "Div. Liquida / EBITDA below 2% last 5 years", checkDebtByEbitda)
}

func checkDebtByEbitda(s stock.Stock) (int, error) {
	return yearsBelow(s,
		5,
		2.0,
		func(yr stock.YearResult) float32 {
			return yr.DebtByEbitda
		})
}
