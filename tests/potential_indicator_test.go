package tests

import (
	ind "github.com/evsamsonov/trading-indicators"
	"github.com/evsamsonov/trading-indicators/mocks"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestShortPotentialIndicator_Calculate(t *testing.T) {
	atrIndicator := &mocks.Indicator{}

	series := ind.NewTimeSeries()

	testCandles := []TestCandle{
		{high: 23, low: 21.27, open: 21.3125, close: 22.1044, time: 1121979600},
		{high: 23.31999, low: 20.15, open: 22.16, close: 23.21608, time: 1122238800},
		{high: 23.2755, low: 19.0, open: 23.2755, close: 22.20585, time: 1122325200},
	}

	for _, item := range testCandles {
		candle := ind.NewCandle(time.Unix(item.time, 0))
		candle.High = item.high
		candle.Low = item.low
		candle.Open = item.open
		candle.Close = item.close

		series.AddCandle(candle)
	}

	indicator := ind.NewPotentialIndicator(series, atrIndicator, ind.NewShortPotentialStrategy())

	// На ATR 0
	atrIndicator.On("Calculate", 1).Return(.0).Once()
	assert.Equal(t, .0, indicator.Calculate(0))

	// На break по high limit
	atrIndicator.On("Calculate", 1).Return(0.1).Once()
	assert.Less(t, math.Abs(indicator.Calculate(0) - 2.01), epsilon)

	// На break по противополной свече
	atrIndicator.On("Calculate", 1).Return(3.).Once()
	assert.Less(t, math.Abs(indicator.Calculate(0) - 2.01), epsilon)

	// На обновление максимального low
	atrIndicator.On("Calculate", 1).Return(3.5).Once()
	atrIndicator.On("Calculate", 2).Return(3.5).Once()
	assert.Less(t, math.Abs(indicator.Calculate(0) - 3.16), epsilon)
}

func TestIntegrationShortPotentialIndicator_Calculate(t *testing.T) {
	series := ind.NewTimeSeries()

	for _, item := range GetTestCandles() {
		candle := ind.NewCandle(time.Unix(item.time, 0))
		candle.High = item.high
		candle.Low = item.low
		candle.Open = item.open
		candle.Close = item.close

		series.AddCandle(candle)
	}

	atrIndicator := ind.NewAtrIndicator(series, 14)

	indicator := ind.NewPotentialIndicator(series, atrIndicator, ind.NewShortPotentialStrategy())

	assert.Equal(t, .0, indicator.Calculate(0))
	assert.Equal(t, .0, indicator.Calculate(11))
	assert.Less(t, math.Abs(indicator.Calculate(12) - 0.773), epsilon)
	assert.Less(t, math.Abs(indicator.Calculate(14) - 0.37), epsilon)
	assert.Equal(t, .0, indicator.Calculate(20))
	assert.Less(t, math.Abs(indicator.Calculate(50) - 2.424), epsilon)
}

