package indicator

import "github.com/evsamsonov/trading-timeseries/timeseries"

// AverageVolume is a indicator to calculate average volume
// for given smoothInterval using simple moving average
type AverageVolume struct {
	series *timeseries.TimeSeries
	period int
	cache  map[int]float64
}

func NewAverageVolume(series *timeseries.TimeSeries, period int) *AverageVolume {
	return &AverageVolume{
		series: series,
		period: period,
		cache:  make(map[int]float64),
	}
}

func (av *AverageVolume) Calculate(index int) float64 {
	if index < av.period-1 {
		return 0
	}

	if val, ok := av.cache[index]; ok {
		return val
	}

	volumeSum := 0.0
	for i := index - av.period + 1; i <= index; i++ {
		volumeSum += float64(av.series.Candle(i).Volume)
	}

	return volumeSum / float64(av.period)
}
