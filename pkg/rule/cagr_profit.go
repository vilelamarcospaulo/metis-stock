package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func CagrProfit() Rule {
	return newRule("CARGR", "CARGR profit above 5% last 5 years", checkCargrProfit)
}

func checkCargrProfit(s stock.Stock) (int, error) {
	return yearsAbove(s,
		5,
		5.0,
		func(yr stock.YearResult) float32 {
			return yr.CagrProfit
		})
}
