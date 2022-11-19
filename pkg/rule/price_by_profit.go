package rule

import (
	"fmt"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func PriceByProfit() Rule {
	return newRule("P/L", "P/L belows 30", checkPriceByProfit)
}

func checkPriceByProfit(stock stock.Stock) (int, error) {
	if len(stock.Results) < 1 {
		return 0, fmt.Errorf("not enough data, %d years", len(stock.Results))
	}

	result := 100
	if stock.Results[0].PriceByProfit > 30.0 {
		result = 0
	}

	return result, nil
}
