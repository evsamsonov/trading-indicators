package tests

import (
	ind "github.com/evsamsonov/trading-indicators"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestShortPotentialStrategy_IsFinish(t *testing.T) {
	strategy := ind.NewShortPotentialStrategy()

	testes := []struct {
		high           float64
		close          float64
		open           float64
		atr            float64
		expectedResult bool
	}{
		{high: 202.01, close: 0, open: 0, atr: 0, expectedResult: true},
		{high: 202, close: 203, open: 202.99, atr: 1.9, expectedResult: true},
		{high: 202, close: 203, open: 203, atr: 1.9, expectedResult: false},
		{high: 202, close: 203, open: 202.99, atr: 2, expectedResult: false},
	}

	initCandle := ind.NewCandle(time.Now())
	initCandle.Open = 201
	initCandle.Low = 200
	strategy.Init(initCandle, 1)

	for _, test := range testes {
		candle := ind.NewCandle(time.Now())
		candle.High = test.high
		candle.Close = test.close
		candle.Open = test.open

		result := strategy.IsFinish(candle, test.atr)

		assert.Equal(t, test.expectedResult, result)
	}
}

func TestShortPotentialStrategy_Process(t *testing.T) {
	strategy := ind.NewShortPotentialStrategy()

	testes := []struct {
		low         float64
		expectedLow float64
	}{
		{low: 199.99, expectedLow: 199.99},
		{low: 200.01, expectedLow: 200},
	}

	initCandle := ind.NewCandle(time.Now())

	for _, test := range testes {
		initCandle.Open = 201
		initCandle.Low = 200
		strategy.Init(initCandle, 1)

		candle := ind.NewCandle(time.Now())
		candle.Low = test.low
		strategy.Process(candle)

		expectedStrategy := ind.NewShortPotentialStrategy()
		initCandle.Low = test.expectedLow
		expectedStrategy.Init(initCandle, 1)

		assert.Equal(t, strategy, expectedStrategy)
	}
}

func TestShortPotentialStrategy_Potential(t *testing.T) {
	strategy := ind.NewShortPotentialStrategy()

	initCandle := ind.NewCandle(time.Now())
	initCandle.Open = 202
	initCandle.Low = 200
	strategy.Init(initCandle, 1)

	assert.Equal(t, strategy.Potential(), 2.)
}
