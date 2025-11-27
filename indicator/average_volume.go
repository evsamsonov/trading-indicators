package indicator

import (
	"sync"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

type AverageVolumeOption func(a *AverageVolume)

func WithAverageVolumeFilter(filter FilterFunc) func(*AverageVolume) {
	return func(t *AverageVolume) {
		t.filter = filter
	}
}

// AverageVolume calculates a simple moving average of volume
// over a given period, optionally filtering candles.
type AverageVolume struct {
	series *timeseries.TimeSeries
	period int
	mu     sync.RWMutex
	cache  map[int]float64
	filter FilterFunc
}

func NewAverageVolume(
	series *timeseries.TimeSeries,
	period int,
	opts ...AverageVolumeOption,
) *AverageVolume {
	averageVolume := &AverageVolume{
		series: series,
		period: period,
		cache:  make(map[int]float64),
	}
	for _, opt := range opts {
		opt(averageVolume)
	}
	return averageVolume
}

func (a *AverageVolume) Calculate(index int) float64 {
	if index < a.period-1 {
		return 0
	}

	a.mu.RLock()
	val, ok := a.cache[index]
	a.mu.RUnlock()
	if ok {
		return val
	}

	volumeSum := 0.0
	count := 0
	for i := index; count < a.period; i-- {
		if i < 0 {
			return 0
		}

		candle := a.series.Candle(i)
		if a.filter != nil && !a.filter(i, candle) {
			continue
		}
		volumeSum += float64(candle.Volume)
		count++
	}

	avg := volumeSum / float64(a.period)

	a.mu.Lock()
	a.cache[index] = avg
	a.mu.Unlock()

	return avg
}
