package indicator

// Индикатор потенциала продажи
type ShortPotentialIndicator struct {
	series       *TimeSeries
	atrIndicator Indicator
	cache        map[int]float64
}

// Создает индикатор потенциала продажи
func NewShortPotentialIndicator(series *TimeSeries, atrIndicator Indicator) Indicator {
	return &ShortPotentialIndicator{series, atrIndicator, make(map[int]float64)}
}

// Вычисляет значение потенциала продажи
func (p *ShortPotentialIndicator) Calculate(index int) float64 {
	if val, ok := p.cache[index]; ok {
		return val
	}

	if p.series.Length() == index+1 {
		return 0
	}

	var candle *candle
	var startOpen float64
	var maxLow float64
	var highLimit float64
	var atr float64

	for i := index + 1; i < p.series.Length(); i++ {
		candle = p.series.Candle(i)
		atr = p.atrIndicator.Calculate(i)
		if atr == 0 {
			return 0
		}

		if startOpen == 0 {
			startOpen = candle.Open
			maxLow = candle.Low
			highLimit = startOpen + atr
		}

		if candle.High > highLimit {
			break
		}

		if candle.Close-candle.Open > 0 && candle.High-maxLow > atr {
			break
		}

		if maxLow > candle.Low {
			maxLow = candle.Low
		}
	}

	return startOpen - maxLow
}
