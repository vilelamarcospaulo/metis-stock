package main

import (
	"github.com/vilelamarcospaulo/metis/executor"
	"github.com/vilelamarcospaulo/metis/external/statusi"
	"github.com/vilelamarcospaulo/metis/pkg/rule"
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func main() {
	stocksDataProvider := statusi.StatusInvestProvider{}

	stocks := []*stock.Stock{
		stock.NewStock("SANB4"),
	}

	rules := []rule.Rule{
		rule.ReturnOnEquity(),
	}

	portfolioFundamentalistEvaluator := executor.NewEvaluator(
		stocksDataProvider, stocks, rules,
	)

	portfolioFundamentalistEvaluator.Execute()
}
