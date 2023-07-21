package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type evaluator func(AssetType stock.Asset, HistoricalData []stock.Historical) float64

type Rule struct {
	Title       string
	Description string
	evaluator   evaluator
	threshold   float64
}

func NewRule(title string, description string, evaluator evaluator) Rule {
	return Rule{
		Title:       title,
		Description: description,
		evaluator:   evaluator,
		threshold:   0.80,
	}
}

func (r Rule) Evaluate(asset stock.Asset, historicalData []stock.Historical) bool {
	return r.evaluator(asset, historicalData) >= r.threshold
}
