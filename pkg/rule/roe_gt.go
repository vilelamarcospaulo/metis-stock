package rule

import (
	"math"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func checkRoeGreaterThan(_ stock.Asset, HistoricalData []stock.Historical) float64 {
	greater := 0
	analised := int(math.Min(float64(len(HistoricalData)), 5))

	for _, val := range HistoricalData[:analised] {
		if val.Return_on_equity >= 5 {
			greater++
		}
	}

	return float64(greater) / float64(analised)
}

func RoeGreaterThan() Rule {
	return NewRule(
		"roe_gt_five",
		"Check company's ROE is greater than 5% for the past years",
		checkRoeGreaterThan,
	)
}
