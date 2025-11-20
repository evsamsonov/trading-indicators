package indicator

import (
	"testing"

	"github.com/evsamsonov/trading-timeseries/timeseries"
	"github.com/stretchr/testify/assert"
)

func TestAverageVolume_Calculate(t *testing.T) {
	series := GetTestSeries()

	tests := []struct {
		name     string
		period   int
		index    int
		filter   AverageVolumeFilterFunc
		expected float64
	}{
		{
			name:     "not enough data",
			period:   3,
			index:    1,
			expected: 0,
		},
		{
			name:     "smoothInterval=3,index=5",
			period:   3,
			index:    5,
			expected: 1385766.66,
		},
		{
			name:     "smoothInterval=14,index=13",
			period:   14,
			index:    13,
			expected: 1851464.2857,
		},
		{
			name:     "with filter",
			period:   3,
			index:    5,
			filter:   func(i int, candle *timeseries.Candle) bool { return candle.High >= 23 },
			expected: 3770033.33,
		},
		{
			name:     "with filter, not enough data",
			period:   4,
			index:    5,
			filter:   func(i int, candle *timeseries.Candle) bool { return candle.High >= 23 },
			expected: 0.,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indicator := NewAverageVolume(series, tt.period, WithAverageVolumeFilter(tt.filter))
			result := indicator.Calculate(tt.index)
			if tt.expected == 0 {
				assert.Equal(t, tt.expected, result)
			} else {
				assert.InEpsilon(t, tt.expected, result, float64EqualityThreshold)
			}
		})
	}
}
