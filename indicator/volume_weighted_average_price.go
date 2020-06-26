package indicator

import (
	"sync"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

const oneDay = time.Hour * 24

// VolumeWeightedAveragePrice represents indicator to calculate volume-weighted average price (VWAP).
// More details https://en.wikipedia.org/wiki/Volume-weighted_average_price
type VolumeWeightedAveragePrice struct {
	series *timeseries.TimeSeries
	cache  *vwapCache
}

// NewVolumeWeightedAveragePrice creates VolumeWeightedAveragePrice
func NewVolumeWeightedAveragePrice(series *timeseries.TimeSeries) *VolumeWeightedAveragePrice {
	return &VolumeWeightedAveragePrice{
		series: series,
		cache:  newVwapCache(),
	}
}

// Calculate returns VWAP value for candle with given index
func (v *VolumeWeightedAveragePrice) Calculate(index int) float64 {
	unit, ok := v.cache.Get(index)
	if ok {
		return unit.vwap
	}

	day := v.series.Candle(index).Time.Truncate(oneDay)

	startIndex, unit := v.findLastCalculated(index, day)
	volumeTotal := unit.volumeTotal
	priceVolumeTotal := unit.priceVolumeTotal

	for i := startIndex; i <= index; i++ {
		candle := v.series.Candle(i)
		if !candle.Time.Truncate(oneDay).Equal(day) {
			break
		}

		typicalPrice := calcTypicalPrice(candle)
		priceVolumeTotal = priceVolumeTotal + float64(candle.Volume)*typicalPrice
		volumeTotal = volumeTotal + candle.Volume

		v.cache.Add(i, vwapUnit{
			vwap:             priceVolumeTotal / float64(volumeTotal),
			priceVolumeTotal: priceVolumeTotal,
			volumeTotal:      volumeTotal,
		})
	}

	unit, _ = v.cache.Get(index)
	return unit.vwap
}

func (v *VolumeWeightedAveragePrice) findLastCalculated(index int, day time.Time) (startIndex int, item vwapUnit) {
	for i := index - 1; i >= 0; i-- {
		candle := v.series.Candle(i)
		if !candle.Time.Truncate(oneDay).Equal(day) {
			return i + 1, vwapUnit{}
		}

		item, ok := v.cache.Get(i)
		if ok {
			return i + 1, item
		}
	}

	return 0, vwapUnit{}
}

func calcTypicalPrice(candle *timeseries.Candle) float64 {
	return (candle.High + candle.Low + candle.Close) / 3
}

type vwapCache struct {
	mu    sync.RWMutex
	items map[int]vwapUnit
}

func newVwapCache() *vwapCache {
	return &vwapCache{
		items: make(map[int]vwapUnit),
	}
}

type vwapUnit struct {
	vwap             float64
	priceVolumeTotal float64
	volumeTotal      int64
}

func (v *vwapCache) Get(index int) (vwapUnit, bool) {
	v.mu.RLock()
	defer v.mu.RUnlock()
	val, ok := v.items[index]
	return val, ok
}

func (v *vwapCache) Add(index int, item vwapUnit) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.items[index] = item
}
