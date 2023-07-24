package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type YearResultList []stock.YearResult
type evaluator func(stock.Asset, YearResultList) float64

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

func (r Rule) Evaluate(asset stock.Asset, h YearResultList) bool {
	return r.evaluator(asset, h) >= r.threshold
}
