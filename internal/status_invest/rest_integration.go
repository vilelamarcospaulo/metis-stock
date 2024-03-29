package status_invest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type StatusInvestResponse struct {
	Success bool
	Data    map[string][]StatusInvestIndicatorData
}

type StatusInvestIndicatorData struct {
	Key   string // roe, p_vp, p_l, lucros_cagr5, dividaliquida_ebitda
	Ranks []HistoricalDataRank
}

type HistoricalDataRank struct {
	Rank  int // Year
	Value float64
}

func (p StatusInvestDataProvider) fetchAPI(ticker string) ([]StatusInvestIndicatorData, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	tickerCode := strings.ToLower(ticker)
	data := url.Values{}
	data.Set("codes", tickerCode)
	data.Set("time", "5")
	data.Add("byQuarter", "false")
	data.Add("futureData", "false")

	req, err := http.NewRequest("POST",
		"https://statusinvest.com.br/acao/indicatorhistoricallist",
		strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "PostmanRuntime/7.29.2")
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Unable to fetch StatusInvest url %s", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to parse response %s", err)
		return nil, err
	}

	var statusIResponse StatusInvestResponse
	json.Unmarshal(body, &statusIResponse)

	return statusIResponse.Data[tickerCode], nil
}
