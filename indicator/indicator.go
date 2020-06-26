package indicator

// Indicator represents interface for indicators
type Indicator interface {
	Calculate(index int) float64 // index of trading candle
}
