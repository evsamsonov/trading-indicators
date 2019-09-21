package tests

import (
	indicator "github.com/evsamsonov/trading-indicators"
	"testing"
	"time"
)

func TestTimeSeries_AddCandle(t *testing.T) {
	series := indicator.NewTimeSeries()

	candleTime := time.Now()
	candle1 := indicator.NewCandle(candleTime)
	if ok := series.AddCandle(candle1); !ok {
		t.Errorf("Не удалось добавить свечу 1")
	}

	candleTime = candleTime.Add(time.Minute)
	candle2 := indicator.NewCandle(candleTime)
	if ok := series.AddCandle(candle2); !ok {
		t.Errorf("Не удалось добавить свечу 2")
	}

	if ok := series.AddCandle(candle2); ok {
		t.Errorf("Повторное добавление свечи")
	}

	candleTime = candleTime.Add(-time.Hour)
	candle3 := indicator.NewCandle(candleTime)
	if ok := series.AddCandle(candle3); ok {
		t.Errorf("Добавлена свеча с временем раньше последней")
	}
}

func TestTimeSeries_Candle(t *testing.T) {
	series := indicator.NewTimeSeries()

	candleTime := time.Now()
	candle := indicator.NewCandle(candleTime)
	series.AddCandle(candle)
	if result := series.Candle(0); result != candle {
		t.Errorf("Не удалось получить свечу")
	}
}

func TestTimeSeries_Length(t *testing.T) {
	series := indicator.NewTimeSeries()

	if series.Length() != 0 {
		t.Errorf("Количество свечей не равно 0")
	}

	candleTime := time.Now()
	candle1 := indicator.NewCandle(candleTime)
	series.AddCandle(candle1)

	candleTime = candleTime.Add(time.Minute)
	candle2 := indicator.NewCandle(candleTime)
	series.AddCandle(candle2)

	if series.Length() != 2 {
		t.Errorf("Количество свечей не равно 2")
	}
}
