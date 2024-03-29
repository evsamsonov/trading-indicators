package indicator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrend_Calculate(t *testing.T) {
	tests := []struct {
		name                 string
		fastVal              float64
		slowVal              float64
		flatMaxDiffInPercent bool
		want                 float64
	}{
		{
			name:    "up trend",
			fastVal: 1.61,
			slowVal: 1.0,
			want:    UpTrend,
		},
		{
			name:    "down trend",
			fastVal: 1.0,
			slowVal: 1.61,
			want:    DownTrend,
		},
		{
			name:    "flat trend",
			fastVal: 1.0,
			slowVal: 1.6,
			want:    FlatTrend,
		},
		{
			name:    "flat trend 2",
			fastVal: 1.6,
			slowVal: 1.0,
			want:    FlatTrend,
		},
		{
			name:                 "flat trend, diff in percent",
			fastVal:              100.5,
			slowVal:              100,
			flatMaxDiffInPercent: true,
			want:                 FlatTrend,
		},
		{
			name:                 "up trend, diff in percent",
			fastVal:              100.7,
			slowVal:              100,
			flatMaxDiffInPercent: true,
			want:                 UpTrend,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slowEMAIndicator := &MockIndicator{}
			fastEMAIndicator := &MockIndicator{}

			ind := NewTrend(
				fastEMAIndicator,
				slowEMAIndicator,
				0.6,
				TrendWithFlatMaxDiffInPercent(tt.flatMaxDiffInPercent),
			)
			fastEMAIndicator.On("Calculate", 1).Return(tt.fastVal)
			slowEMAIndicator.On("Calculate", 1).Return(tt.slowVal)

			res := ind.Calculate(1)
			assert.Equal(t, tt.want, res)
		})
	}
}
