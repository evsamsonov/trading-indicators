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
	var indicator ind.Indicator

	atrIndicator := &mocks.Indicator{}
	potentialStrategy := &mocks.PotentialStrategy{}

	series := ind.NewTimeSeries()

	// Пустая серия
	indicator = ind.NewPotentialIndicator(series, atrIndicator, potentialStrategy)
	assert.Equal(t, .0, indicator.Calculate(0))

	testCandles := []TestCandle{
		{time: 1121979600},
		{time: 1122238800},
		{time: 1122325200},
	}

	for _, item := range testCandles {
		candle := ind.NewCandle(time.Unix(item.time, 0))
		series.AddCandle(candle)
	}

	// На ATR 0
	indicator = ind.NewPotentialIndicator(series, atrIndicator, potentialStrategy)
	atrIndicator.On("Calculate", 1).Return(.0).Once()
	assert.Equal(t, .0, indicator.Calculate(0))

	// На IsFinish
	indicator = ind.NewPotentialIndicator(series, atrIndicator, potentialStrategy)
	atrIndicator.On("Calculate", 1).Return(0.1).Once()
	potentialStrategy.On("Init", series.Candle(1), 0.1).Once()
	potentialStrategy.On("IsFinish", series.Candle(1), 0.1).Return(true).Once()
	potentialStrategy.On("Potential").Return(2.).Once()
	assert.Equal(t, indicator.Calculate(0), 2.)

	// С вызовом Process
	indicator = ind.NewPotentialIndicator(series, atrIndicator, potentialStrategy)
	atrIndicator.On("Calculate", 1).Return(0.1).Once()
	potentialStrategy.On("Init", series.Candle(1), 0.1).Once()
	potentialStrategy.On("IsFinish", series.Candle(1), 0.1).Return(false).Once()
	potentialStrategy.On("Process", series.Candle(1)).Once()

	atrIndicator.On("Calculate", 2).Return(0.1).Once()
	potentialStrategy.On("Init", series.Candle(2), 0.1).Once()
	potentialStrategy.On("IsFinish", series.Candle(2), 0.1).Return(true).Once()

	potentialStrategy.On("Potential").Return(2.).Once()
	assert.Equal(t, indicator.Calculate(0), 2.)
	assert.Equal(t, indicator.Calculate(0), 2.)   // Проверка кэша
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
