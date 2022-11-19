package stock

type StockClass = int64

const (
	ETF StockClass = iota
	STOCK
	FII
)

type Stock struct {
	class        StockClass
	ticker       string
	Fundamentals []Fundamentals
}

func NewStock(ticker string) *Stock {
	return &Stock{
		class:  STOCK,
		ticker: ticker,
	}
}

func AppendF(s *Stock, f Fundamentals) {
	s.Fundamentals = append(s.Fundamentals, f)
}
