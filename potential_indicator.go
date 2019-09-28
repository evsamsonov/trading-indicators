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
	Init(candle *Candle, atr float64)			// Инициализация
	IsFinish(candle *Candle, atr float64) bool 	// Это последняя свеча
	Process(candle *Candle)					    // Обрабатывает свечу
	Potential() float64							// Возвращает значение потенциала
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

	p.cache[index] = p.strategy.Potential()
	return p.cache[index]
}