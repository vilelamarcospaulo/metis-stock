package statusi

import (
	"fmt"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

type StatusInvestProvider struct {
}

func (p StatusInvestProvider) LoadFundamentals(s *stock.Stock, currentYear int, windowSize int) error {
	fmt.Printf("fetching StatusInvest report for [%s]\n", s.Ticker)
	indexedHistorical, err := p.fetchData(s.Ticker)
	if err != nil {
		return err
	}

	for year := currentYear; year > currentYear-windowSize; year-- {
		fmt.Printf("building result for [%s][%d]\n", s.Ticker, currentYear)
		fundamentals := p.buildYearResult(indexedHistorical, year)
		s.Results = append(s.Results, fundamentals)
	}

	return nil
}

func (p StatusInvestProvider) fetchData(ticker string) (map[string]HistoricalData, error) {
	historicalDataList, err := p.fetchAPI(ticker)
	if err != nil {
		return nil, err
	}

	indexedData := make(map[string]HistoricalData)
	for _, data := range historicalDataList {
		indexedData[data.Key] = data
	}

	return indexedData, nil
}

func (p StatusInvestProvider) buildYearResult(historicalData map[string]HistoricalData, year int) stock.YearResult {
	result := stock.YearResult{}

	roe, err := historicalData["roe"].getByRank(year)
	if err != nil {
		panic(1)
	}

	result.ROE = roe

	return result
}
