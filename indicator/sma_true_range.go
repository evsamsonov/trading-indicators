package indicator

import (
	"math"
	"sync"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

type SmaTrueRangeOption func(*SmaTrueRange)

// WithSmaTrueRangeFilter allows filtering candles when calculating SMA ATR.
func WithSmaTrueRangeFilter(filter FilterFunc) SmaTrueRangeOption {
	return func(a *SmaTrueRange) {
		a.filter = filter
	}
}

// SmaTrueRange calculates a simple moving average of the true range (TR)
// over a given period, optionally filtering candles.
type SmaTrueRange struct {
	series *timeseries.TimeSeries
	period int
	mu     sync.RWMutex
	cache  map[int]float64
	filter FilterFunc
}

func NewSmaTrueRange(
	series *timeseries.TimeSeries,
	period int,
	opts ...SmaTrueRangeOption,
) *SmaTrueRange {
	smaAtr := &SmaTrueRange{
		series: series,
		period: period,
		cache:  make(map[int]float64),
	}
	for _, opt := range opts {
		opt(smaAtr)
	}
	return smaAtr
}

func (a *SmaTrueRange) Calculate(index int) float64 {
	if index < 0 || index >= a.series.Length() {
		return 0
	}

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
	for i := index; count < a.period && i >= 0; i-- {
		candle := a.series.Candle(i)
		if a.filter != nil && !a.filter(i, candle) {
			continue
		}

		trueRangeSum += a.calculateTrueRange(i)
		count++
	}

	if count == 0 {
		return 0
	}

	avg := trueRangeSum / float64(count)

	a.mu.Lock()
	a.cache[index] = avg
	a.mu.Unlock()

	return avg
}

func (a *SmaTrueRange) calculateTrueRange(index int) float64 {
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
