package status_invest

import (
	"fmt"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type StatusInvestDataProvider struct {
}

type indexedIndicatorData map[string]StatusInvestIndicatorData

func (data indexedIndicatorData) buildHistoricalData(year int) stock.Historical {
	return stock.Historical{
		Year:             year,
		Return_on_equity: data["roe"].getByYear(year),
	}
}

func (h StatusInvestIndicatorData) getByYear(year int) float64 {
	for _, r := range h.Ranks {
		if r.Rank == year {
			return r.Value
		}
	}

	// TODO :: Return (nil, error), next layer fallback to another data provider
	fmt.Printf("results of year %d not found to %s", year, h.Key)
	panic(1)
}

func (s StatusInvestDataProvider) GetHistoricalData(ticker string) []stock.Historical {
	result, err := s.fetchAPI(ticker)
	if err != nil {
		panic(1)
	}

	indexedData := indexedIndicatorData{}
	for _, data := range result {
		indexedData[data.Key] = data
	}

	historicalData := []stock.Historical{}
	for year := 2022; year >= 2018; year-- {
		yearData := indexedData.buildHistoricalData(year)
		historicalData = append(historicalData, yearData)
	}

	return historicalData
}
