package rule

import (
	"fmt"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func PriceByEquity() Rule {
	return newRule("P/VP", "P/VP belows 5", CheckPriceByEquity)
}

func CheckPriceByEquity(stock stock.Stock) (int, error) {
	if len(stock.Results) < 1 {
		return 0, fmt.Errorf("not enough data, %d years", len(stock.Results))
	}

	result := 100
	if stock.Results[0].PriceByEquity > 5.0 {
		result = 0
	}

	return result, nil
}
