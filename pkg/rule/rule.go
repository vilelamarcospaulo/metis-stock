package rule

import (
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type Rule struct {
	Title       string
	Description string
	threshold   int                            // 0 - 100
	processor   func(stock.Stock) (int, error) // 0 - 100
}

func newRule(title string, desc string, processor func(stock.Stock) (int, error)) Rule {
	return Rule{
		Title:       title,
		Description: desc,
		threshold:   70,
		processor:   processor,
	}
}

func (r *Rule) Eval(stock stock.Stock) (float32, error) {
	score, err := r.processor(stock)
	if err != nil {
		return 0, err
	}

	if score > r.threshold {
		return float32(score), nil
	}

	return -1, nil
}
