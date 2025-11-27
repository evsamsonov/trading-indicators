package indicator

import (
	"testing"

	"github.com/evsamsonov/trading-timeseries/timeseries"
	"github.com/stretchr/testify/assert"
)

func TestSmaAverageTrueRange_Calculate(t *testing.T) {
	series := GetTestSeries()
	smaAtr := NewSmaAverageTrueRange(series, 14)

	value := smaAtr.Calculate(13)
	expected := 0.7602071429

	assert.InEpsilon(t, expected, value, float64EqualityThreshold, "Expected %f, got %f", expected, value)
}

func TestSmaAverageTrueRange_WithFilter(t *testing.T) {
	series := GetTestSeries()
	filter := func(i int, candle *timeseries.Candle) bool {
		return candle.Volume > 2000000
	}

	t.Run("enough data", func(t *testing.T) {
		indicator := NewSmaAverageTrueRange(series, 5, WithSmaAverageTrueRangeFilter(filter))

		value := indicator.Calculate(20)
		expected := 1.247816
		assert.InEpsilon(t, expected, value, float64EqualityThreshold, "Expected %f, got %f", expected, value)
	})

	t.Run("not enough data", func(t *testing.T) {
		indicator := NewSmaAverageTrueRange(series, 6, WithSmaAverageTrueRangeFilter(filter))

		value := indicator.Calculate(20)
		assert.Equal(t, 0., value)
	})
}
