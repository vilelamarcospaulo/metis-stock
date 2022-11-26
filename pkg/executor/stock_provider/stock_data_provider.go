package stock_provider

import "github.com/vilelamarcospaulo/metis/pkg/stock"

type StockProvider interface {
	LoadFundamentals(s *stock.Stock, currentYear int, windowSize int) error
}
