package indicator

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
	"github.com/stretchr/testify/assert"
)

func TestVolumeWeightedAveragePrice_Calculate(t *testing.T) {
	series := timeseries.New()

	assert.Nil(t, series.AddCandle(createCandle("2020-06-24T00:00:00+00:00", 1, 2, 3, 100)))
	assert.Nil(t, series.AddCandle(createCandle("2020-06-25T00:00:00+00:00", 4, 5, 6, 200)))
	assert.Nil(t, series.AddCandle(createCandle("2020-06-25T01:00:00+00:00", 7, 8, 9, 300)))
	assert.Nil(t, series.AddCandle(createCandle("2020-06-26T00:00:00+00:00", 10, 11, 12, 400)))

	tests := []struct {
		index int
		want  float64
	}{
		{index: 0, want: 2},
		{index: 1, want: 5},
		{index: 2, want: 6.8},
	}

	vwapIndicator := NewVolumeWeightedAveragePrice(series)

	for _, tt := range tests {
		t.Run(fmt.Sprintf("index=%d", tt.index), func(t *testing.T) {
			vwap := vwapIndicator.Calculate(tt.index)
			assert.InEpsilon(t, tt.want, vwap, float64EqualityThreshold)
		})
	}
}

func createCandle(date string, high, low, close float64, volume int64) *timeseries.Candle {
	candle := timeseries.NewCandle(parseDate(date))
	candle.High = high
	candle.Low = low
	candle.Close = close
	candle.Volume = volume
	return candle
}

func parseDate(t string) time.Time {
	date, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Fatal("failed to parse time", err)
	}
	return date
}
