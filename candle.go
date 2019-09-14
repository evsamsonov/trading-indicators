package indicator

import (
	"time"
)

// Описывает торговую свечу
type candle struct {
	Time   time.Time
	High   float64
	Low    float64
	Open   float64
	Close  float64
	Volume int64
}

// Создает новую торговую свечу
func NewCandle(time time.Time) *candle {
	return &candle{Time: time}
}
