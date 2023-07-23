package main

import (
	"github.com/sirupsen/logrus"
	"github.com/vilelamarcospaulo/metis/internal/status_invest"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	logrus.Info("Starting Metis")

	dataProvider := status_invest.StatusInvestDataProvider{}
	dataProvider.GetHistoricalData("PETR4")
}
