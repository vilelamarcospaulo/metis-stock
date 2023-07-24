package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func RoeGreaterThan() Rule {
	return NewRule(
		"roe_gt_five",
		"Check company's ROE is greater than 5% for the past years",
		func(_ stock.Asset, yearResults YearResultList) float64 {
			return yearResults.OccurrenceRate(func(yr stock.YearResult) bool {
				return yr.ROE >= 5
			})
		},
	)
}
