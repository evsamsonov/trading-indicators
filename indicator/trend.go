package indicator

import "math"

const (
	FlatTrend float64 = 0
	UpTrend   float64 = 1
	DownTrend float64 = -1
)

type TrendOption func(*Trend)

// TrendWithFlatMaxDiffInPercent allows to pass flatMaxDiff in percent.
// The default value is false
func TrendWithFlatMaxDiffInPercent(val bool) func(*Trend) {
	return func(t *Trend) {
		t.flatMaxDiffInPercent = val
	}
}

// Trend returns a trend direction.
// It bases on fast (with shorter period) and slow EMA.
// flatMaxDiff allows setting max difference between
// fast and slow EMA when Calculate returns the flat.
type Trend struct {
	fastEMAIndicator     Indicator
	slowEMAIndicator     Indicator
	flatMaxDiff          float64
	flatMaxDiffInPercent bool
}

func NewTrend(
	fastEMAIndicator Indicator,
	slowEMAIndicator Indicator,
	flatMaxDiff float64,
	opts ...TrendOption,
) *Trend {
	trend := &Trend{
		fastEMAIndicator: fastEMAIndicator,
		slowEMAIndicator: slowEMAIndicator,
		flatMaxDiff:      flatMaxDiff,
	}
	for _, opt := range opts {
		opt(trend)
	}
	return trend
}

func (t *Trend) Calculate(index int) float64 {
	fastVal := t.fastEMAIndicator.Calculate(index)
	slowVal := t.slowEMAIndicator.Calculate(index)

	flatMaxDiff := t.flatMaxDiff
	if t.flatMaxDiffInPercent {
		flatMaxDiff = slowVal * t.flatMaxDiff / 100
	}

	if math.Abs(fastVal-slowVal)-flatMaxDiff <= 1e-6 {
		return FlatTrend
	}
	if fastVal > slowVal {
		return UpTrend
	}
	return DownTrend
}
