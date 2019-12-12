package indicator

import (
	"github.com/evsamsonov/trading-timeseries/timeseries"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestAtrIndicator_Calculate(t *testing.T) {
	series := timeseries.New()

	for _, item := range GetTestCandles() {
		candle := timeseries.NewCandle(time.Unix(item.time, 0))
		candle.High = item.high
		candle.Low = item.low
		candle.Open = item.open
		candle.Close = item.close

		err := series.AddCandle(candle)
		assert.Nil(t, err)
	}

	testCases := []struct {
		period   int
		index    int
		expected float64
	}{
		{16, 15, 0.720238},
		{16, 14, 0.0},
		{8, 15, 0.650576},
	}

	for _, test := range testCases {
		atrIndicator := NewAtrIndicator(series, test.period)
		atr := atrIndicator.Calculate(test.index)
		expectedAtr := test.expected
		if math.Abs(atr-expectedAtr) > epsilon {
			t.Errorf("Полученное значение %f не равно ожидаемому %f", atr, expectedAtr)
		}
	}
}
