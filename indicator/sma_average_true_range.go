package indicator

import (
	"math"
	"sync"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

type SmaAverageTrueRangeOption func(*SmaAverageTrueRange)

// WithSmaAverageTrueRangeFilter allows filtering candles when calculating SMA ATR.
func WithSmaAverageTrueRangeFilter(filter FilterFunc) SmaAverageTrueRangeOption {
	return func(a *SmaAverageTrueRange) {
		a.filter = filter
	}
}

// SmaAverageTrueRange calculates a simple moving average of the true range (TR)
// over a given period, optionally filtering candles.
type SmaAverageTrueRange struct {
	series *timeseries.TimeSeries
	period int
	mu     sync.RWMutex
	cache  map[int]float64
	filter FilterFunc
}

func NewSmaAverageTrueRange(
	series *timeseries.TimeSeries,
	period int,
	opts ...SmaAverageTrueRangeOption,
) *SmaAverageTrueRange {
	smaAtr := &SmaAverageTrueRange{
		series: series,
		period: period,
		cache:  make(map[int]float64),
	}
	for _, opt := range opts {
		opt(smaAtr)
	}
	return smaAtr
}

func (a *SmaAverageTrueRange) Calculate(index int) float64 {
	if index < a.period-1 {
		return 0
	}

	a.mu.RLock()
	if val, ok := a.cache[index]; ok {
		a.mu.RUnlock()
		return val
	}
	a.mu.RUnlock()

	trueRangeSum := 0.0
	count := 0
	for i := index; count < a.period; i-- {
		if i < 0 {
			return 0
		}

		candle := a.series.Candle(i)
		if a.filter != nil && !a.filter(i, candle) {
			continue
		}

		trueRangeSum += a.calculateTrueRange(i)
		count++
	}

	avg := trueRangeSum / float64(a.period)

	a.mu.Lock()
	a.cache[index] = avg
	a.mu.Unlock()

	return avg
}

func (a *SmaAverageTrueRange) calculateTrueRange(index int) float64 {
	candle := a.series.Candle(index)
	highLowDiff := math.Abs(candle.High - candle.Low)
	if index == 0 {
		return highLowDiff
	}

	prevCandle := a.series.Candle(index - 1)

	highCloseDiff := math.Abs(candle.High - prevCandle.Close)
	lowCloseDiff := math.Abs(candle.Low - prevCandle.Close)

	return math.Max(highLowDiff, math.Max(highCloseDiff, lowCloseDiff))
}
