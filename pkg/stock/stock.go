package stock

type StockClass = int64

const (
	ETF StockClass = iota
	STOCK
	FII
)

type Stock struct {
	class  StockClass
	ticker string
}

func NewStock(ticker string) *Stock {
	return &Stock{
		class:  STOCK,
		ticker: ticker,
	}
}
