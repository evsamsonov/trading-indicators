package indicator

import "github.com/evsamsonov/trading-timeseries/timeseries"

type FilterFunc func(i int, candle *timeseries.Candle) bool
