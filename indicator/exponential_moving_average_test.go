package indicator

import (
	"sync"
	"testing"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
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

func TestExponentialMovingAverage_Calculate_WithFilter(t *testing.T) {
	tests := []struct {
		name           string
		closes         []float64
		smoothInterval int
		filter         ExponentialMovingAverageFilterFunc
		expected       []float64
	}{
		{
			name:           "filter middle candle",
			closes:         []float64{10, 20, 30, 40},
			smoothInterval: 3,
			filter: func(i int, candle *timeseries.Candle) bool {
				return i != 1
			},
			expected: []float64{10, 10, 20, 30},
		},
		{
			name:           "filter first candle",
			closes:         []float64{10, 20, 30},
			smoothInterval: 1,
			filter: func(i int, candle *timeseries.Candle) bool {
				return i != 0
			},
			expected: []float64{0, 20, 30},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			series := timeseries.New()
			baseTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
			for i, closePrice := range test.closes {
				candle := timeseries.NewCandle(baseTime.Add(time.Duration(i) * time.Hour))
				candle.Close = closePrice
				series.AddCandle(candle)
			}

			ema, err := NewExponentialMovingAverage(
				series,
				test.smoothInterval,
				WithExponentialMovingAverageFilter(test.filter),
			)
			assert.Nil(t, err)

			for i := 0; i < series.Length(); i++ {
				val := ema.Calculate(i)
				if test.expected[i] == 0 {
					assert.Equal(t, test.expected[i], val, "Index %d", i)
				} else {
					assert.InEpsilon(t, test.expected[i], val, float64EqualityThreshold, "Index %d", i)
				}
			}
		})
	}
}

func TestExponentialMovingAverage_CalculateAfterAddCandle(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("%v", r)
		}
	}()
	series := GetTestSeries()
	indicator, err := NewExponentialMovingAverage(series, 3)
	assert.Nil(t, err)

	err = series.AddCandle(&timeseries.Candle{Time: time.Now()})
	assert.Nil(t, err)
	indicator.Calculate(series.Length() - 1)
}

func BenchmarkExponentialMovingAverage_Calculate(b *testing.B) {
	series := GetTestSeries()
	indicator, err := NewExponentialMovingAverage(series, 3)
	assert.Nil(b, err)

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				indicator.Calculate(60 + j)
			}(j)
		}
		wg.Wait()
	}
}
