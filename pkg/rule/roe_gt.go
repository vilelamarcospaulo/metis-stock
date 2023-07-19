package rule

import "github.com/vilelamarcospaulo/metis/pkg/stock"

func checkRoeGreaterThan(_ stock.Asset, HistoricalData []stock.Historical) float64 {
	greater := 0
	checked := 0
	for _, val := range HistoricalData {
		checked++
		if val.Return_on_equity >= 5 {
			greater++
		}
	}

	return float64(greater) / float64(checked)
}

func RoeGreaterThan() Rule {
	return newRule(
		"roe_gt_five",
		"Check company's ROE is greater than 5% for the past years",
		checkRoeGreaterThan,
	)
}
