package indicator

import "github.com/evsamsonov/trading-timeseries/timeseries"

type FilterFunc func(i int, candle *timeseries.Candle) bool

// AverageVolumeFilterFunc is deprecated, use FilterFunc instead.
type AverageVolumeFilterFunc = FilterFunc
