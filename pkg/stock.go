package models

type AssetType int

const (
	Stock AssetType = iota
	ETF
	FII
)

type Asset struct {
	ID     int64  `json:"id"`
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
}

func NewAsset(id int64, ticker string, name string) Asset {
	return Asset{
		ID:     id,
		Ticker: ticker,
		Name:   name,
	}
}
