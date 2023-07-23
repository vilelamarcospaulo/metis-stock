package status_invest

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type StatusInvestDataProvider struct {
}

type indexedIndicatorData map[string]StatusInvestIndicatorData

func (data indexedIndicatorData) buildHistoricalData(year int) stock.YearResult {
	return stock.YearResult{
		Year:         year,
		ROE:          data["roe"].getByYear(year),
		RevenueCAGR:  data["receitas_cagr5"].getByYear(year),
		DebtByEBITDA: data["dividaliquida_ebitda"].getByYear(year),
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

func (s StatusInvestDataProvider) GetHistoricalData(ticker string) []stock.YearResult {
	// TODO :: Cache results to avoid multiple requests to the same ticker
	// when evaluating multiple wallets

	logrus.Debugf("Fetching historical data for %s", ticker)
	result, err := s.fetchAPI(ticker)
	if err != nil {
		logrus.Errorf("Error fetching historical data for %s: %s", ticker, err)
		panic(1)
	}

	indexedData := indexedIndicatorData{}
	logrus.Debugf("Indexing indicators data for %s", ticker)
	for _, data := range result {
		indexedData[data.Key] = data
	}

	historicalData := []stock.YearResult{}

	// Only last 5 years???
	for year := 2022; year >= 2018; year-- {
		logrus.Debugf("Building historical data for %s in %d", ticker, year)
		yearData := indexedData.buildHistoricalData(year)
		historicalData = append(historicalData, yearData)
	}

	logrus.Debugf("Historical data for %s: %v", ticker, historicalData)
	return historicalData
}
