package indicator

import (
	"github.com/evsamsonov/trading-timeseries/timeseries"
	"math"
)

// AverageTrueRange represents indicator to calculate Average True Range (ATR)
// https://en.wikipedia.org/wiki/Average_true_range
type AverageTrueRange struct {
	series       *timeseries.TimeSeries
	period       int
	cache        map[int]float64
	lastComputed int
}

func NewAverageTrueRange(series *timeseries.TimeSeries, period int) Indicator {
	return &AverageTrueRange{series, period, make(map[int]float64), period - 1}
}

func (ind *AverageTrueRange) Calculate(index int) float64 {
	if index < ind.period-1 {
		return 0
	}

	if val, ok := ind.cache[index]; ok {
		return val
	}

	for i := ind.lastComputed; i <= index; i++ {
		ind.cache[i] = ind.doCalculate(i)
	}

	ind.lastComputed = index
	return ind.cache[index]
}

func (ind *AverageTrueRange) doCalculate(i int) float64 {
	if i == ind.period-1 {
		// For first index ATR equals sum of previous values divided by period
		trueRangeSum := 0.0
		for j := 0; j < ind.period; j++ {
			trueRangeSum += ind.calculateTrueRange(j)
		}

		return float64(trueRangeSum) / float64(ind.period)
	}

	return (ind.cache[i-1]*float64(ind.period-1) + ind.calculateTrueRange(i)) / float64(ind.period)
}

func (ind *AverageTrueRange) calculateTrueRange(index int) float64 {
	candle := ind.series.Candle(index)
	highLowDiff := math.Abs(candle.High - candle.Low)
	if index == 0 {
		return highLowDiff
	}

	prevCandle := ind.series.Candle(index - 1)

	highCloseDiff := math.Abs(candle.High - prevCandle.Close)
	lowCloseDiff := math.Abs(candle.Low - prevCandle.Close)

	return math.Max(highLowDiff, math.Max(highCloseDiff, lowCloseDiff))
}
