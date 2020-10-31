package indicator

import (
	"errors"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

// ExponentialMovingAverage represents indicator to calculate Exponential Moving Average (EMA)
type ExponentialMovingAverage struct {
	series         *timeseries.TimeSeries
	smoothInterval int
	maxIndex       int
	cache          []float64
	smooth         float64
}

func NewExponentialMovingAverage(series *timeseries.TimeSeries, smoothInterval int) (*ExponentialMovingAverage, error) {
	if series.Length() == 0 {
		return nil, errors.New("series is empty")
	}
	if smoothInterval < 0 {
		return nil, errors.New("smoothInterval cannot be negative")
	}
	smooth := 2 / (float64(smoothInterval) + 1)

	return &ExponentialMovingAverage{
		series:         series,
		smoothInterval: smoothInterval,
		cache:          make([]float64, series.Length()),
		smooth:         smooth,
	}, nil
}

func (a *ExponentialMovingAverage) Calculate(index int) float64 {
	if index >= len(a.cache) {
		a.cache = append(a.cache, 0)
		a.cache = a.cache[:cap(a.cache)]
	}
	if a.cache[index] != 0 {
		return a.cache[index]
	}
	if a.maxIndex == 0 {
		a.cache[0] = a.series.Candle(0).Close
	}

	for i := a.maxIndex + 1; i <= index; i++ {
		a.cache[i] = a.smooth*a.series.Candle(i).Close + (1-a.smooth)*a.cache[i-1]
	}

	a.maxIndex = index
	return a.cache[index]
}
