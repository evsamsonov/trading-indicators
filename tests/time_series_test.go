package tests

import (
	"github.com/evsamsonov/trading-timeseries/timeseries"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeSeries_AddCandle(t *testing.T) {
	series := timeseries.New()

	candleTime := time.Now()
	candle1 := timeseries.NewCandle(candleTime)
	if err := series.AddCandle(candle1); err != nil {
		t.Errorf("Не удалось добавить свечу 1")
	}

	candleTime = candleTime.Add(time.Minute)
	candle2 := timeseries.NewCandle(candleTime)
	if err := series.AddCandle(candle2); err != nil {
		t.Errorf("Не удалось добавить свечу 2")
	}

	if err := series.AddCandle(candle2); err == nil {
		t.Errorf("Повторное добавление свечи")
	}

	candleTime = candleTime.Add(-time.Hour)
	candle3 := timeseries.NewCandle(candleTime)
	if err := series.AddCandle(candle3); err == nil {
		t.Errorf("Добавлена свеча с временем раньше последней")
	}
}

func TestTimeSeries_Candle(t *testing.T) {
	series := timeseries.New()

	candleTime := time.Now()
	candle := timeseries.NewCandle(candleTime)
	err := series.AddCandle(candle)
	assert.Nil(t, err)

	if result := series.Candle(0); result != candle {
		t.Errorf("Не удалось получить свечу")
	}
}

func TestTimeSeries_Length(t *testing.T) {
	series := timeseries.New()

	if series.Length() != 0 {
		t.Errorf("Количество свечей не равно 0")
	}

	candleTime := time.Now()
	candle1 := timeseries.NewCandle(candleTime)
	err := series.AddCandle(candle1)
	assert.Nil(t, err)

	candleTime = candleTime.Add(time.Minute)
	candle2 := timeseries.NewCandle(candleTime)
	err = series.AddCandle(candle2)
	assert.Nil(t, err)

	if series.Length() != 2 {
		t.Errorf("Количество свечей не равно 2")
	}
}
