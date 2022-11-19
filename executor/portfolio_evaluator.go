package executor

import (
	"fmt"

	"github.com/vilelamarcospaulo/metis/executor/stock_provider"
	"github.com/vilelamarcospaulo/metis/pkg/rule"
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type PortfolioEvaluator struct {
	provider stock_provider.StockProvider
	Assets   []*stock.Stock
	Rules    []rule.Rule
}

func NewEvaluator(provider stock_provider.StockProvider, assets []*stock.Stock, rules []rule.Rule) *PortfolioEvaluator {
	return &PortfolioEvaluator{
		Assets:   assets,
		Rules:    rules,
		provider: provider,
	}
}

func (p *PortfolioEvaluator) Execute() {
	p.loadFundamentals()
	p.eval()
}

func (p *PortfolioEvaluator) loadFundamentals() {
	fmt.Print("------------------------------------\n")
	fmt.Print("-------- START LOAD REPORTS --------\n")

	// TODO :: Use goroutines for load multiple assets
	for _, asset := range p.Assets {
		fmt.Printf("-- LOADING REPORT TO [%s]\n", asset.Ticker)
		p.provider.LoadFundamentals(asset, 2022, 5)
	}
	fmt.Print("------------------------------------\n\n")
}

func (p *PortfolioEvaluator) eval() {
	fmt.Print("-------- START RUNNING RULES -------\n")

	for _, asset := range p.Assets {
		fmt.Print("------------------------------------\n\n")
		fmt.Printf("-- [%s] \n", asset.Ticker)

		for _, rule := range p.Rules {
			// fmt.Printf("\tRUNNING RULE [%s] \n", rule.Title)

			result, err := rule.Eval(*asset)
			if err != nil {
				fmt.Printf("\t ERROR :: [%s] :: %s\n", rule.Title, err.Error())
				continue
			}

			fmt.Printf("\t [%s] :: [%f] \n", rule.Description, result)
		}
	}
}
