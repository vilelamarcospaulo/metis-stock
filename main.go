package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vilelamarcospaulo/metis/internal/status_invest"
	"github.com/vilelamarcospaulo/metis/pkg/evaluator"
	"github.com/vilelamarcospaulo/metis/pkg/rule"
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	logrus.Info("Starting Metis")

	dataProvider := status_invest.StatusInvestDataProvider{}

	// TODO :: Retrieve assets and rules from user wallet and preferences
	assets := []stock.Asset{
		stock.NewAsset(0, "SANB4", "Santander"),
		stock.NewAsset(0, "VIVT3", "Santander"),
	}
	rules := []rule.Rule{
		rule.RoeGreaterThan(),
		rule.RevenueCagrGreaterThan(),
		rule.DebtByEbitaLowerThan(),
	}

	evaluator := evaluator.NewEvaluator(dataProvider, rules)
	result := evaluator.Evaluate(assets)

	// TODO :: Extract to report module
	for _, assetEvaluation := range result {
		fmt.Printf("%s: %f\n", assetEvaluation.Asset.Name, assetEvaluation.Score)

		for _, ruleEvaluation := range assetEvaluation.Result {
			fmt.Printf("  %s: %t\n", ruleEvaluation.Rule.Description, ruleEvaluation.Result)
		}
	}
}
