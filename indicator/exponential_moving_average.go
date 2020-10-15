package indicator

import (
	"errors"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

// ExponentialMovingAverage represents indicator to calculate Exponential Moving Average (EMA)
type ExponentialMovingAverage struct {
	series         *timeseries.TimeSeries
	smoothInterval int
	cache          map[int]float64
	maxCacheIndex  int
}

func NewExponentialMovingAverage(series *timeseries.TimeSeries, smoothInterval int) (*ExponentialMovingAverage, error) {
	if series.Length() == 0 {
		return nil, errors.New("series is empty")
	}

	if smoothInterval < 0 {
		return nil, errors.New("smoothInterval cannot be negative")
	}

	return &ExponentialMovingAverage{
		series:         series,
		smoothInterval: smoothInterval,
		cache:          make(map[int]float64),
	}, nil
}

func (ma *ExponentialMovingAverage) Calculate(index int) float64 {
	if ema, ok := ma.cache[index]; ok {
		return ema
	}

	smooth := 2 / (float64(ma.smoothInterval) + 1)
	if len(ma.cache) == 0 {
		ma.cache[0] = ma.series.Candle(0).Close
	}

	for i := ma.maxCacheIndex + 1; i <= index; i++ {
		ma.cache[i] = smooth*ma.series.Candle(i).Close + (1-smooth)*ma.cache[i-1]
	}

	ma.maxCacheIndex = index

	return ma.cache[index]
}
