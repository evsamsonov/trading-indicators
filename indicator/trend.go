package indicator

import "math"

const (
	FlatTrend float64 = 0
	UpTrend float64 = 1
	DownTrend float64 = -1
)

// Trend indicator allows to define trend direction. It bases on fast and slow EMA.
// flatMaxDiff determines max difference between fast (with shorter period)
// and slow EMA when Calculate returns flat
type Trend struct {
	fastEMAIndicator Indicator
	slowEMAIndicator Indicator
	flatMaxDiff float64
}

func NewTrend(
	fastEMAIndicator Indicator,
	slowEMAIndicator Indicator,
	flatMaxDiff float64,
) *Trend {
	return &Trend{
		fastEMAIndicator: fastEMAIndicator,
		slowEMAIndicator: slowEMAIndicator,
		flatMaxDiff: flatMaxDiff,
	}
}

func (t *Trend) Calculate(index int) float64 {
	fastVal := t.fastEMAIndicator.Calculate(index)
	slowVal := t.slowEMAIndicator.Calculate(index)

	if math.Abs(fastVal-slowVal) - t.flatMaxDiff <= 1e-6 {
		return FlatTrend
	}
	if fastVal > slowVal {
		return UpTrend
	}
	return DownTrend
}
