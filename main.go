package main

import (
	"fmt"

	"github.com/vilelamarcospaulo/metis/internal/status_invest"
)

func main() {
	dataProvider := status_invest.StatusInvestDataProvider{}
	historicalData := dataProvider.GetHistoricalData("PETR4")

	fmt.Println(historicalData)
}
