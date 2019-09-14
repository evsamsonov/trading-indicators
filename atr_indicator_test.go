package indicator

import (
	"math"
	"testing"
	"time"
)

const epsilon = 1e-6

func TestAtrIndicator_Compute(t *testing.T) {
	series := NewTimeSeries()

	candleData := []struct {
		high  float64
		low   float64
		open  float64
		close float64
	}{
		{high: 23, low: 21.27, open: 21.3125, close: 22.1044},
		{high: 23.31999, low: 22.15, open: 22.15, close: 23.21608},
		{high: 23.2755, low: 22.111, open: 23.2755, close: 22.20585},
		{high: 22.58, low: 22.015, open: 22.166, close: 22.2},
		{high: 22.52, low: 22.13501, open: 22.29969, close: 22.15},
		{high: 22.3, low: 21.601, open: 22.20443, close: 21.861},
		{high: 22.1, low: 21.45, open: 21.8499, close: 22.09},
		{high: 22.38, low: 22.0511, open: 22.111, close: 22.25952},
		{high: 22.488, low: 21.61, open: 22.3005, close: 21.71516},
		{high: 22.789, low: 21.61001, open: 21.8, close: 22.59434},
		{high: 22.699, low: 22.351, open: 22.55, close: 22.64},
		{high: 22.73994, low: 22.45001, open: 22.595, close: 22.56998},
		{high: 22.69, low: 21.8, open: 22.605, close: 22.37001},
		{high: 22.58, low: 22.26, open: 22.493, close: 22.397},
		{high: 22.442, low: 22.0011, open: 22.397, close: 22.099},
		{high: 22.16, low: 21.72, open: 22.09, close: 21.937},
	}

	candleTime := time.Now()
	for _, item := range candleData {
		candle := NewCandle(candleTime)
		candle.High = item.high
		candle.Low = item.low
		candle.Open = item.open
		candle.Close = item.close

		series.AddCandle(candle)
		candleTime = candleTime.Add(time.Minute)
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
