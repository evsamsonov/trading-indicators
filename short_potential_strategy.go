package indicator

// Стратегия для расчета потенциала продажи
type ShortPotentialStrategy struct {
	startOpen float64
	maxLow float64
	highLimit float64
}

func NewShortPotentialStrategy() *ShortPotentialStrategy {
	return &ShortPotentialStrategy{}
}

func (s *ShortPotentialStrategy) Init(candle *Candle, atr float64) {
	s.startOpen = candle.Open
	s.maxLow = candle.Low
	s.highLimit = candle.Open + atr
}

func (s *ShortPotentialStrategy) IsFinish(candle *Candle, atr float64) bool {
	if candle.High > s.highLimit {
		return true
	}

	if candle.Close-candle.Open > 0 && candle.High-s.maxLow > atr {
		return true
	}

	return false
}

func (s *ShortPotentialStrategy) Process(candle *Candle) {
	if s.maxLow > candle.Low {
		s.maxLow = candle.Low
	}
}

func (s *ShortPotentialStrategy) Potential() float64 {
	return s.startOpen - s.maxLow
}