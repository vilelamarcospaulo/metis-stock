package evaluator

import (
	"github.com/vilelamarcospaulo/metis/pkg/rule"
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type DataProvider interface {
	GetHistoricalData(ticker string) []stock.Historical
}

type RuleEvaluation struct {
	Rule   rule.Rule
	Result bool
}

type AssetEvaluation struct {
	Asset          stock.Asset
	HistoricalData []stock.Historical
	Score          float64
	Result         []RuleEvaluation
}

func (a *AssetEvaluation) CalculateScore() {
	// TODO - This is a very simple implementation. We should consider
	//        the rule's weight in the final score.

	for _, ruleEvaluation := range a.Result {
		if ruleEvaluation.Result {
			a.Score++
		}
	}
}

type Evaluator struct {
	dataProvider DataProvider
	rules        []rule.Rule
}

func NewEvaluator(dataProvider DataProvider, rules []rule.Rule) Evaluator {
	return Evaluator{
		dataProvider: dataProvider,
		rules:        rules,
	}
}

func (e *Evaluator) Evaluate(assets []stock.Asset) []AssetEvaluation {
	evaluations := []AssetEvaluation{}
	for _, asset := range assets {
		historicalData := e.dataProvider.GetHistoricalData(asset.Ticker)
		assetEvaluation := AssetEvaluation{Asset: asset}

		for _, rule := range e.rules {
			ruleResult := rule.Evaluate(asset, historicalData)
			assetEvaluation.Result = append(assetEvaluation.Result,
				RuleEvaluation{Rule: rule, Result: ruleResult})
		}

		assetEvaluation.CalculateScore()
		evaluations = append(evaluations, assetEvaluation)
	}

	return evaluations
}
