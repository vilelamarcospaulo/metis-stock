package rule_test

import (
	"testing"

	"github.com/vilelamarcospaulo/metis/pkg/rule"
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func buildHistoricalData(roe []float64) []stock.YearResult {
	historicalData := []stock.YearResult{}
	for _, val := range roe {
		historicalData = append(historicalData,
			stock.YearResult{ROE: val})
	}

	return historicalData
}

func TestRoeGreaterThan_WhenBellowThan5_InAllYears(t *testing.T) {
	asset := stock.NewAsset(3, "FOO", "Foo Company")
	historicalData := buildHistoricalData([]float64{4, 3, 2, 1, 0})
	rule := rule.RoeGreaterThan()

	if rule.Evaluate(asset, historicalData) {
		t.Errorf("ROE [4,3,2,1,0] should be false")
	}
}

func TestRoeGreaterThan_WhenGreaterThan5_InHalfOfYears(t *testing.T) {
	asset := stock.NewAsset(3, "FOO", "Foo Company")
	historicalData := buildHistoricalData([]float64{10, 13, 11, 1, 0})
	rule := rule.RoeGreaterThan()

	if rule.Evaluate(asset, historicalData) {
		t.Errorf("ROE [10,13,11,1,0] should be false")
	}
}

func TestRoeGreaterThan_WhenGreaterThan5_ExceptOneYear(t *testing.T) {
	asset := stock.NewAsset(3, "FOO", "Foo Company")
	historicalData := buildHistoricalData([]float64{11, 12, 17, 11, 2})
	rule := rule.RoeGreaterThan()

	if !rule.Evaluate(asset, historicalData) {
		t.Errorf("ROE [11,12,17,11,2] should be true")
	}
}

func TestRoeGreaterThan_WhenGreaterThan5_InAllYears(t *testing.T) {
	asset := stock.NewAsset(3, "FOO", "Foo Company")
	historicalData := buildHistoricalData([]float64{22, 13, 10, 6, 8})
	rule := rule.RoeGreaterThan()

	if !rule.Evaluate(asset, historicalData) {
		t.Errorf("ROE [22,13,10,6,8] should be true")
	}
}

func TestRoeGreaterThan_WhenGreaterThan5_OnlyInFirst5Years(t *testing.T) {
	asset := stock.NewAsset(3, "FOO", "Foo Company")
	historicalData := buildHistoricalData([]float64{22, 13, 10, 6, 8, 1, 1, 1, 1, 1, 1})
	rule := rule.RoeGreaterThan()

	if !rule.Evaluate(asset, historicalData) {
		t.Errorf("ROE [22,13,10,6,8] should be true")
	}
}
