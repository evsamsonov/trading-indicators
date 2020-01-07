package indicator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverageVolume_Calculate(t *testing.T) {
	series := GetTestSeries()

	tests := []struct {
		name     string
		period   int
		index    int
		expected float64
	}{
		{name: "not enough data", period: 3, index: 1, expected: 0},
		{name: "period=3,index=3", period: 3, index: 2, expected: 3770033.33},
		{name: "period=14,index=13", period: 14, index: 13, expected: 1851464.2857},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			indicator := NewAverageVolume(series, test.period)
			result := indicator.Calculate(test.index)
			if test.expected == 0 {
				assert.Equal(t, test.expected, result)
			} else {
				assert.InEpsilon(t, test.expected, result, float64EqualityThreshold)
			}
		})
	}
}
