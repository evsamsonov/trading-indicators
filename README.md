[![Build Status](https://travis-ci.org/evsamsonov/trading-indicators.svg?branch=master)](https://travis-ci.org/evsamsonov/trading-indicators)
[![Go Report Card](https://goreportcard.com/badge/github.com/evsamsonov/trading-indicators)](https://goreportcard.com/report/github.com/evsamsonov/trading-indicators)

# Trading indicators

The set of trading indicators

## Installation

```bash
$ go get github.com/evsamsonov/trading-indicators/indicator  
```

## Usage

All indicators requires trading data in [**timeseries**](https://github.com/evsamsonov/trading-timeseries) structure 

```go
dataset := []struct {
    Time   time.Time
    High   float64
    Low    float64
    Open   float64
    Close  float64
    Volume int64
}{
    {Time: time.Unix(1121979600, 0), High: 23, Low: 21.27, Open: 21.3125, Close: 22.1044, Volume: 4604900},
    {Time: time.Unix(1122238800, 0), High: 23.31999, Low: 22.15, Open: 22.15, Close: 23.21608, Volume: 4132600},
}

series := timeseries.New()
for _, item := range dataset {
    candle := timeseries.NewCandle(item.Time)
    candle.Open = item.Open
    candle.Close = item.Close
    candle.High = item.High
    candle.Low = item.Low
    candle.Volume = item.Volume

    err := series.AddCandle(candle)
    if err != nil {
        log.Printf("Failed to add candle: %v\n", err)
    }
}
```

### Average True Range

Indicator calculates [Average True Range](https://en.wikipedia.org/wiki/Average_true_range ) (ATR)

```go
period := 2
atrIndicator := indicator.NewAverageTrueRange(series, period)
fmt.Println(atrIndicator.Calculate(1))   // 1.4727950000000014
```

### Average Volume

Indicator calculates Average Volume

```go
period := 2
atrIndicator := indicator.NewAverageVolume(series, period)
fmt.Println(atrIndicator.Calculate(1))  // 4368750
```

### Exponential Moving Average

Indicator calculates Exponential Moving Average

```go
smoothInterval := 2
atrIndicator := indicator.NewExponentialMovingAverage(series, smoothInterval)
fmt.Println(atrIndicator.Calculate(1))  // 22.84552
```

### Trend

The indicator returns a trend direction. It bases on fast (with shorter period) and slow EMA. The third parameter (flatMaxDiff) of NewTrend allows setting max difference between fast and slow EMA when Calculate returns the flat. Option TrendWithFlatMaxDiffInPercent allows to pass flatMaxDiff in percent

```go
fastEMAIndicator, err := indicator.NewExponentialMovingAverage(series, 14)
if err != nil {
    log.Fatalln(err)
}
slowEMAIndicator, err := indicator.NewExponentialMovingAverage(series, 50)
if err != nil {
    log.Fatalln(err)
}

trendIndicator := indicator.NewTrend(fastEMAIndicator, slowEMAIndicator, 0.6, TrendWithFlatMaxDiffInPercent(false))
trend := trendIndicator.Calculate(1)
switch trend {
case indicator.UpTrend:
    fmt.Println("Up trend")
case indicator.DownTrend:
    fmt.Println("Down trend")
case indicator.FlatTrend:
    fmt.Println("Flat trend")
}
```
