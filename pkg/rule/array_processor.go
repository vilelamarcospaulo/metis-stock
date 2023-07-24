package rule

import (
	"math"

	"github.com/vilelamarcospaulo/metis/pkg/stock"
)

func (yearResults YearResultList) OccurrenceRate(
	predicate func(stock.YearResult) bool,
) float64 {
	occurences := 0
	sampleSize := int(math.Min(float64(len(yearResults)), 5))

	for _, val := range yearResults[:sampleSize] {
		if predicate(val) {
			occurences++
		}
	}

	return float64(occurences) / float64(sampleSize)
}
