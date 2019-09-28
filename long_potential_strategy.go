package indicator

// Стратегия для расчета потенциала покупки
type LongPotentialStrategy struct {
	startOpen float64
	maxHigh   float64
	lowLimit  float64
}

func NewLongPotentialStrategy() *LongPotentialStrategy {
	return &LongPotentialStrategy{}
}

func (s *LongPotentialStrategy) Init(candle *Candle, atr float64) {
	s.startOpen = candle.Open
	s.maxHigh = candle.High
	s.lowLimit = candle.Open - atr
}

func (s *LongPotentialStrategy) IsFinish(candle *Candle, atr float64) bool {
	if candle.Low < s.lowLimit {
		return true
	}

	if candle.Close-candle.Open < 0 && candle.Low-s.maxHigh > atr {
		return true
	}

	return false
}

func (s *LongPotentialStrategy) Process(candle *Candle) {
	if s.maxHigh < candle.High {
		s.maxHigh = candle.High
	}
}

func (s *LongPotentialStrategy) Potential() float64 {
	return s.maxHigh - s.startOpen
}
