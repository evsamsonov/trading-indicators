package indicator

// Представляет последовательность торговых свечей
type TimeSeries struct {
	candles []*candle
}

// Создает пустую структуру для списка торговых свечей
func NewTimeSeries() *TimeSeries {
	ts := new(TimeSeries)
	ts.candles = make([]*candle, 0)

	return ts
}

// Добавляет торговую свечу в список.
// Принимает свечи последовательно
func (ts *TimeSeries) AddCandle(c *candle) bool {
	if c == nil {
		return false
	}

	if ts.LastCandle() == nil || c.Time.After(ts.LastCandle().Time) {
		ts.candles = append(ts.candles, c)
		return true
	}

	return false
}

// Возвращает последнюю свечу в последовательности
func (ts *TimeSeries) LastCandle() *candle {
	if len(ts.candles) > 0 {
		return ts.candles[len(ts.candles)-1]
	}

	return nil
}

// Возвращает количество элементов в серии
func (ts *TimeSeries) Length() int {
	return len(ts.candles)
}

// Возвращает указанную свечу
func (ts *TimeSeries) Candle(index int) *candle {
	return ts.candles[index]
}
