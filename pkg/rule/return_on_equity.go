package rule

import (
	"fmt"
	"math"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func ReturnOnEquity() Rule {
	return newRule("ROE", "ROE above 5% last 5 years", calcRoe)
}

func calcRoe(stock stock.Stock) (int, error) {
	yearsToConsider := 5 // TODO :: Create a config file

	if len(stock.Results) < yearsToConsider {
		return 0, fmt.Errorf("not enough data, %d years", len(stock.Results))
	}

	yearsAboveExpected := 0
	for year := 0; year < yearsToConsider; year++ {
		if stock.Results[year].ROE == math.MinInt {
			return 0, fmt.Errorf("not valid data for year %d", year)
		}

		if stock.Results[year].ROE > 0.5 {
			yearsAboveExpected += 1
		}
	}

	return ((yearsAboveExpected / yearsToConsider) * 100), nil
}
