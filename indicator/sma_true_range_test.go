package indicator

import (
	"testing"

	"github.com/evsamsonov/trading-timeseries/timeseries"
	"github.com/stretchr/testify/assert"
)

func TestSmaTrueRange_Calculate(t *testing.T) {
	series := GetTestSeries()
	smaAtr := NewSmaTrueRange(series, 14)

	value := smaAtr.Calculate(13)
	expected := 0.7602071429

	assert.InEpsilon(t, expected, value, float64EqualityThreshold, "Expected %f, got %f", expected, value)
}

func TestSmaTrueRange_WithFilter(t *testing.T) {
	series := GetTestSeries()
	filter := func(i int, candle *timeseries.Candle) bool {
		return candle.Volume > 2000000
	}

	t.Run("enough data", func(t *testing.T) {
		indicator := NewSmaTrueRange(series, 5, WithSmaTrueRangeFilter(filter))

		value := indicator.Calculate(20)
		expected := 1.247816
		assert.InEpsilon(t, expected, value, float64EqualityThreshold, "Expected %f, got %f", expected, value)
	})

	t.Run("not enough data", func(t *testing.T) {
		indicator := NewSmaTrueRange(series, 6, WithSmaTrueRangeFilter(filter))

		value := indicator.Calculate(20)
		assert.Equal(t, 0., value)
	})
}
