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
		stock.NewStock("MOVI3"),
		stock.NewStock("NEOE3"),
		stock.NewStock("SULA11"),
	}

	rules := []rule.Rule{
		rule.PriceByEquity(),
		rule.PriceByProfit(),
		rule.ReturnOnEquity(),
		rule.CagrProfit(),
		rule.DebtByEbitda(),
	}

	portfolioFundamentalistEvaluator := executor.NewEvaluator(
		stocksDataProvider, stocks, rules,
	)

	portfolioFundamentalistEvaluator.Execute()
}
