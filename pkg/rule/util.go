package rule

import (
	"fmt"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func yearsAbove(stock stock.Stock, yearsToConsider int, threshold float32, getValue func(stock.YearResult) float32) (int, error) {
	if len(stock.Results) < yearsToConsider {
		return 0, fmt.Errorf("not enough data, %d years needed [%d]", len(stock.Results), yearsToConsider)
	}

	yearsAboveExpected := 0
	for year := 0; year < yearsToConsider; year++ {
		itemValue := getValue(stock.Results[year])

		if itemValue > threshold {
			yearsAboveExpected += 1
		}
	}

	return ((yearsAboveExpected / yearsToConsider) * 100), nil
}

func yearsBelow(stock stock.Stock, yearsToConsider int, threshold float32, getValue func(stock.YearResult) float32) (int, error) {
	if len(stock.Results) < yearsToConsider {
		return 0, fmt.Errorf("not enough data, %d years needed [%d]", len(stock.Results), yearsToConsider)
	}

	yearsBelowExpected := 0
	for year := 0; year < yearsToConsider; year++ {
		itemValue := getValue(stock.Results[year])

		if itemValue < threshold {
			yearsBelowExpected += 1
		}

		// fmt.Printf("%s|%f|%d\n", stock.Ticker, (float32(yearsBelowExpected)/float32(yearsToConsider))*100, yearsBelowExpected)
	}

	return int((float32(yearsBelowExpected) / float32(yearsToConsider)) * 100), nil
}
