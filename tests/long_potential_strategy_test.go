package tests

import (
	ind "github.com/evsamsonov/trading-indicators"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLongPotentialStrategy_IsFinish(t *testing.T) {
	strategy := ind.NewLongPotentialStrategy()

	testes := []struct{
		low float64
		close float64
		open float64
		atr float64
		expectedResult bool
	} {
		{low: 199.99, close: 0, open: 0, atr: 0, expectedResult: true},
		{low: 202, close: 203, open: 203.01, atr: 1.99, expectedResult: true},
		{low: 201, close: 203, open: 203, atr: 0.9, expectedResult: false},
		{low: 201, close: 203, open: 203.1, atr: 1.9, expectedResult: false},
	}

	initCandle := ind.NewCandle(time.Now())
	initCandle.Open = 201
	initCandle.High = 200
	strategy.Init(initCandle, 1)

	for _, test := range testes {
		candle := ind.NewCandle(time.Now())
		candle.Low = test.low
		candle.Close = test.close
		candle.Open = test.open

		result := strategy.IsFinish(candle, test.atr)

		assert.Equal(t, test.expectedResult, result)
	}
}

func TestLongPotentialStrategy_Process(t *testing.T) {
	strategy := ind.NewLongPotentialStrategy()

	testes := []struct{
		high float64
		expectedHigh float64
	} {
		{high: 201.01, expectedHigh: 201.01},
		{high: 200.99, expectedHigh: 201},
	}

	initCandle := ind.NewCandle(time.Now())

	for _, test := range testes {
		initCandle.Open = 200
		initCandle.High = 201
		strategy.Init(initCandle, 1)

		candle := ind.NewCandle(time.Now())
		candle.High = test.high
		strategy.Process(candle)

		expectedStrategy := ind.NewLongPotentialStrategy()
		initCandle.High = test.expectedHigh
		expectedStrategy.Init(initCandle, 1)

		assert.Equal(t, strategy, expectedStrategy)
	}
}

func TestLongPotentialStrategy_Potential(t *testing.T) {
	strategy := ind.NewLongPotentialStrategy()

	initCandle := ind.NewCandle(time.Now())
	initCandle.Open = 200
	initCandle.High = 202
	strategy.Init(initCandle, 1)

	assert.Equal(t, strategy.Potential(), 2.)
}
