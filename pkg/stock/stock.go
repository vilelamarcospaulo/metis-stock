package stock

type StockClass = int64

const (
	ETF StockClass = iota
	STOCK
	FII
)

type Stock struct {
	class   StockClass
	Ticker  string
	Results []YearResult
}

func NewStock(ticker string) *Stock {
	return &Stock{
		class:  STOCK,
		Ticker: ticker,
	}
}
