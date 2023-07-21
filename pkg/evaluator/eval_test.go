package evaluator_test

import (
	"testing"

	"github.com/vilelamarcospaulo/metis/pkg/evaluator"
	"github.com/vilelamarcospaulo/metis/pkg/rule"
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type MockDataProvider struct {
}

func (m MockDataProvider) GetHistoricalData(ticker string) []stock.Historical {
	if ticker == "PETR4" {
		return []stock.Historical{
			{Year: 2022, Return_on_equity: 10},
			{Year: 2021, Return_on_equity: 3},
			{Year: 2020, Return_on_equity: 2},
			{Year: 2019, Return_on_equity: 3},
			{Year: 2018, Return_on_equity: 3},
		}
	}
	return []stock.Historical{
		{Year: 2022, Return_on_equity: 10},
		{Year: 2021, Return_on_equity: 10},
		{Year: 2020, Return_on_equity: 10},
		{Year: 2019, Return_on_equity: 10},
		{Year: 2018, Return_on_equity: 3},
	}
}

func buildEvaluator() evaluator.Evaluator {
	return evaluator.NewEvaluator(
		MockDataProvider{},
		[]rule.Rule{
			rule.RoeGreaterThan(),
			rule.NewRule(
				"always true",
				"always true",
				func(asset stock.Asset, historicalData []stock.Historical) float64 {
					return 1.0
				},
			),
		})
}

func TestEvaluator(t *testing.T) {
	evaluator := buildEvaluator()
	assets := []stock.Asset{
		stock.NewAsset(1, "PETR4", "Petrobras"),
		stock.NewAsset(2, "VALE3", "Vale"),
	}

	assetEvaluations := evaluator.Evaluate(assets)

	if len(assetEvaluations) != 2 {
		t.Errorf("Expected 2 evaluations, got %d", len(assetEvaluations))
	}

	if assetEvaluations[0].Asset.Ticker != "PETR4" {
		t.Errorf("Expected PETR4, got %s", assetEvaluations[0].Asset.Name)
	}

	if assetEvaluations[0].Score != 1 {
		t.Errorf("Expected PETR4 have score 1, got %f", assetEvaluations[0].Score)
	}

	if assetEvaluations[0].Result[0].Result != false {
		t.Errorf("Expected PETR4 have false result to ROE, got %t", assetEvaluations[0].Result[0].Result)
	}

	if assetEvaluations[1].Asset.Ticker != "VALE3" {
		t.Errorf("Expected VALE3, got %s", assetEvaluations[1].Asset.Name)
	}

	if assetEvaluations[1].Score != 2 {
		t.Errorf("Expected VALE3 have score 2, got %f", assetEvaluations[1].Score)
	}

	if assetEvaluations[1].Result[0].Result != true {
		t.Errorf("Expected VALE3 have true result to ROE, got %t", assetEvaluations[1].Result[0].Result)
	}
}
