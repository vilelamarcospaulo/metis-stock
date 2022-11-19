package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type Rule struct {
	title       string
	description string
	threshold   int                            // 0 - 100
	processor   func(stock.Stock) (int, error) // 0 - 100
}

func newRule(title string, desc string, processor func(stock.Stock) (int, error)) *Rule {
	return &Rule{
		title:       title,
		description: desc,
		threshold:   90,
		processor:   processor,
	}
}

func (r *Rule) Eval(stock stock.Stock) (float32, error) {
	score, err := r.processor(stock)
	if err != nil {
		return 0, err
	}

	if score > r.threshold {
		return 1, nil
	}

	return -1, nil
}
