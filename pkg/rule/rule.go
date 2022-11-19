package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type Rule struct {
	title       string
	description string
	threshold   int                   // 0 - 100
	processor   func(stock.Stock) int // 0 - 100
}

func newRule(title string, desc string, processor func(stock.Stock) int) *Rule {
	return &Rule{
		title:       title,
		description: desc,
		threshold:   90,
		processor:   processor,
	}
}

func Eval(r *Rule, stock stock.Stock) float32 {
	if r.processor(stock) > r.threshold {
		return 1
	}

	return -1
}
