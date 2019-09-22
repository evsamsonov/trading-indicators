package indicator

// Индикатор потенциала продажи
type PotentialIndicator struct {
	series       *TimeSeries
	atrIndicator Indicator
	strategy 	 PotentialStrategy
	cache        map[int]float64
}

// Описывает стратегию расчета потенциала продажа
type PotentialStrategy interface {
	// Инициализация
	Init(candle *Candle, atr float64)

	// Это последняя свеча
	IsFinish(candle *Candle, atr float64) bool

	// Обрабатывает свечу
	Process(candle *Candle)

	// Возвращает значение потенциала
	Potential() float64
}

// Создает индикатор потенциала продажи
func NewPotentialIndicator(series *TimeSeries, atrIndicator Indicator, strategy PotentialStrategy) Indicator {
	return &PotentialIndicator{series, atrIndicator, strategy, make(map[int]float64)}
}

// Вычисляет значение потенциала продажи
func (p *PotentialIndicator) Calculate(index int) float64 {
	if val, ok := p.cache[index]; ok {
		return val
	}

	if p.series.Length() == index+1 {
		return 0
	}

	var candle *Candle
	var atr float64
	for i := index + 1; i < p.series.Length(); i++ {
		candle = p.series.Candle(i)
		atr = p.atrIndicator.Calculate(i)
		if atr == 0 {
			return 0
		}

		if i == index + 1 {
			p.strategy.Init(candle, atr)
		}

		if p.strategy.IsFinish(candle, atr) {
			break
		}

		p.strategy.Process(candle)
	}

	return p.strategy.Potential()
}

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
