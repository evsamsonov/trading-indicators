package indicator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExponentialMovingAverage_Calculate(t *testing.T) {
	series := GetTestSeries()

	tests := []struct {
		name           string
		smoothInterval int
		index          int
		expected       float64
	}{
		{name: "not enough data", smoothInterval: 3, index: 0, expected: 22.1044},
		{name: "smoothInterval=3,index=2", smoothInterval: 3, index: 2, expected: 22.433045},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			indicator, err := NewExponentialMovingAverage(series, test.smoothInterval)
			assert.Nil(t, err)

			result := indicator.Calculate(test.index)
			if test.expected == 0 {
				assert.Equal(t, test.expected, result)
			} else {
				assert.InEpsilon(t, test.expected, result, float64EqualityThreshold)
			}
		})
	}
}
