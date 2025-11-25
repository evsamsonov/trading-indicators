package indicator

import (
	"errors"
	"sync"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

type ExponentialMovingAverageFilterFunc func(i int, candle *timeseries.Candle) bool

type ExponentialMovingAverageOption func(e *ExponentialMovingAverage)

func WithExponentialMovingAverageFilter(filter ExponentialMovingAverageFilterFunc) ExponentialMovingAverageOption {
	return func(e *ExponentialMovingAverage) {
		e.filter = filter
	}
}

// ExponentialMovingAverage represents indicator to calculate Exponential Moving Average (EMA)
type ExponentialMovingAverage struct {
	series         *timeseries.TimeSeries
	smoothInterval int
	filter         ExponentialMovingAverageFilterFunc

	smooth   float64
	mu       sync.Mutex
	maxIndex int
	cache    []float64
}

func NewExponentialMovingAverage(
	series *timeseries.TimeSeries,
	smoothInterval int,
	opts ...ExponentialMovingAverageOption,
) (*ExponentialMovingAverage, error) {
	if series.Length() == 0 {
		return nil, errors.New("series is empty")
	}
	if smoothInterval < 0 {
		return nil, errors.New("smoothInterval cannot be negative")
	}
	smooth := 2 / (float64(smoothInterval) + 1)

	ema := &ExponentialMovingAverage{
		series:         series,
		smoothInterval: smoothInterval,
		cache:          make([]float64, series.Length()),
		smooth:         smooth,
		maxIndex:       -1,
	}
	for _, opt := range opts {
		opt(ema)
	}
	return ema, nil
}

func (e *ExponentialMovingAverage) Calculate(index int) float64 {
	e.mu.Lock()
	defer e.mu.Unlock()

	if index >= len(e.cache) {
		for range len(e.cache) - index + 1 {
			e.cache = append(e.cache, 0)
		}
	}
	if index <= e.maxIndex && e.cache[index] != 0 {
		return e.cache[index]
	}

	for i := e.maxIndex + 1; i <= index; i++ {
		candle := e.series.Candle(i)

		prevEMA := float64(0)
		if i > 0 {
			prevEMA = e.cache[i-1]
		}

		// If filtered, use previous value
		if e.filter != nil && !e.filter(i, candle) {
			e.cache[i] = prevEMA
			continue
		}

		if prevEMA == 0 {
			e.cache[i] = candle.Close
		} else {
			e.cache[i] = e.smooth*candle.Close + (1-e.smooth)*prevEMA
		}
	}

	e.maxIndex = index
	return e.cache[index]
}
