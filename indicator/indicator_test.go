package indicator

import (
	"log"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

const float64EqualityThreshold = 1e-6

type TestCandle struct {
	high   float64
	low    float64
	open   float64
	close  float64
	time   int64
	volume int64
}

func GetTestCandles() []TestCandle {
	return []TestCandle{
		{high: 23, low: 21.27, open: 21.3125, close: 22.1044, time: 1121979600, volume: 4604900},
		{high: 23.31999, low: 22.15, open: 22.15, close: 23.21608, time: 1122238800, volume: 4132600},
		{high: 23.2755, low: 22.111, open: 23.2755, close: 22.20585, time: 1122325200, volume: 2572600},
		{high: 22.58, low: 22.015, open: 22.166, close: 22.2, time: 1122411600, volume: 1441300},
		{high: 22.52, low: 22.13501, open: 22.29969, close: 22.15, time: 1122498000, volume: 1184000},
		{high: 22.3, low: 21.601, open: 22.20443, close: 21.861, time: 1122584400, volume: 1532000},
		{high: 22.1, low: 21.45, open: 21.8499, close: 22.09, time: 1122843600, volume: 1177800},
		{high: 22.38, low: 22.0511, open: 22.111, close: 22.25952, time: 1122930000, volume: 1419400},
		{high: 22.488, low: 21.61, open: 22.3005, close: 21.71516, time: 1123016400, volume: 1878000},
		{high: 22.789, low: 21.61001, open: 21.8, close: 22.59434, time: 1123102800, volume: 2547900},
		{high: 22.699, low: 22.351, open: 22.55, close: 22.64, time: 1123189200, volume: 649700},
		{high: 22.73994, low: 22.45001, open: 22.595, close: 22.56998, time: 1123448400, volume: 587800},
		{high: 22.69, low: 21.8, open: 22.605, close: 22.37001, time: 1123534800, volume: 1431000},
		{high: 22.58, low: 22.26, open: 22.493, close: 22.397, time: 1123621200, volume: 761500},
		{high: 22.442, low: 22.0011, open: 22.397, close: 22.099, time: 1123707600, volume: 1208000},
		{high: 22.16, low: 21.72, open: 22.09, close: 21.937, time: 1123794000, volume: 1130300},
		{high: 22.41705, low: 21.801, open: 21.95, close: 22.4, time: 1124053200, volume: 687100},
		{high: 23.16499, low: 22.305, open: 22.444, close: 22.928, time: 1124139600, volume: 1937800},
		{high: 23.705, low: 22.755, open: 22.9, close: 23.7, time: 1124226000, volume: 2417300},
		{high: 24, low: 23.301, open: 23.749, close: 23.83, time: 1124312400, volume: 1727400},
		{high: 24.42999, low: 23.27451, open: 23.27451, close: 24.18324, time: 1124398800, volume: 1234400},
		{high: 25.47, low: 23.435, open: 23.435, close: 25.18305, time: 1124658000, volume: 3260700},
		{high: 25.82, low: 25.16, open: 25.16, close: 25.60184, time: 1124744400, volume: 1661400},
		{high: 26.2835, low: 25.48, open: 26.2835, close: 25.54999, time: 1124830800, volume: 1173500},
		{high: 25.619, low: 25.2, open: 25.6, close: 25.31662, time: 1124917200, volume: 861400},
		{high: 25.55, low: 25.051, open: 25.436, close: 25.11544, time: 1125003600, volume: 1196700},
		{high: 25.549, low: 24.89005, open: 25.12, close: 25.32633, time: 1125262800, volume: 798000},
		{high: 25.64999, low: 25.237, open: 25.3, close: 25.385, time: 1125349200, volume: 1014100},
		{high: 25.548, low: 24.9, open: 25.385, close: 25.51, time: 1125435600, volume: 1326900},
		{high: 25.9, low: 25.401, open: 25.515, close: 25.50001, time: 1125522000, volume: 1121400},
		{high: 25.77, low: 25.5, open: 25.52501, close: 25.625, time: 1125608400, volume: 642300},
		{high: 25.75, low: 25.27001, open: 25.6275, close: 25.31, time: 1125867600, volume: 686500},
		{high: 25.298, low: 24.259, open: 25.27053, close: 24.39595, time: 1125954000, volume: 2541700},
		{high: 24.785, low: 23.667, open: 24.4, close: 24.26, time: 1126040400, volume: 2517600},
		{high: 25.5, low: 24.201, open: 24.3, close: 25.44803, time: 1126126800, volume: 2819900},
		{high: 26.1, low: 25.37, open: 25.50405, close: 25.68808, time: 1126213200, volume: 2785900},
		{high: 25.8994, low: 25.349, open: 25.685, close: 25.41001, time: 1126472400, volume: 975200},
		{high: 26.1, low: 25.1, open: 25.41, close: 25.899, time: 1126558800, volume: 1852800},
		{high: 25.7, low: 25.351, open: 25.6, close: 25.57033, time: 1126731600, volume: 774400},
		{high: 26.09, low: 25.4, open: 25.67, close: 25.75856, time: 1126818000, volume: 1327600},
		{high: 26.4972, low: 25.55501, open: 25.76, close: 26.079, time: 1127077200, volume: 1672400},
		{high: 27.09, low: 26, open: 26, close: 26.851, time: 1127163600, volume: 2977800},
		{high: 27.308, low: 26.45005, open: 26.8605, close: 27, time: 1127250000, volume: 1743500},
		{high: 27.47777, low: 26.61, open: 27.08, close: 26.69875, time: 1127336400, volume: 1415000},
		{high: 26.98899, low: 26.115, open: 26.7005, close: 26.927, time: 1127422800, volume: 1350800},
		{high: 26.9135, low: 26.425, open: 26.9135, close: 26.7, time: 1127682000, volume: 847200},
		{high: 27.0929, low: 26.70001, open: 26.707, close: 26.998, time: 1127768400, volume: 942700},
		{high: 27.67, low: 26.951, open: 26.9975, close: 27.49517, time: 1127854800, volume: 1883900},
		{high: 27.699, low: 26.875, open: 27.64, close: 27.05, time: 1127941200, volume: 1218600},
		{high: 27.4791, low: 27.1, open: 27.1, close: 27.37981, time: 1128027600, volume: 1026400},
		{high: 27.599, low: 27.275, open: 27.35, close: 27.56988, time: 1128286800, volume: 859500},
		{high: 27.78, low: 27.3, open: 27.575, close: 27.45011, time: 1128373200, volume: 1384900},
		{high: 27.5, low: 26.655, open: 27.5, close: 26.83195, time: 1128459600, volume: 1380600},
		{high: 26.6, low: 25.383, open: 26.6, close: 25.998, time: 1128546000, volume: 2342300},
		{high: 26.251, low: 25.151, open: 25.998, close: 25.55101, time: 1128632400, volume: 1367700},
		{high: 26.75, low: 25.515, open: 25.92, close: 26.64367, time: 1128891600, volume: 1381900},
		{high: 27.3, low: 26.4, open: 26.7495, close: 26.94842, time: 1128978000, volume: 1603200},
		{high: 27.29, low: 25.9, open: 27.075, close: 26.2, time: 1129064400, volume: 1044900},
		{high: 26.1, low: 24.95, open: 26.1, close: 25.22521, time: 1129150800, volume: 1355400},
		{high: 25.25, low: 24.3285, open: 25.2415, close: 24.51378, time: 1129237200, volume: 1332000},
		{high: 24.96, low: 24.37, open: 24.894, close: 24.70341, time: 1129496400, volume: 1149200},
		{high: 25.14877, low: 24.6501, open: 24.72775, close: 24.889, time: 1129582800, volume: 891900},
		{high: 24.7953, low: 23.78551, open: 24.7953, close: 23.825, time: 1129669200, volume: 1322100},
		{high: 24.35, low: 24, open: 24.1, close: 24.08998, time: 1129755600, volume: 913000},
		{high: 23.897, low: 23.347, open: 23.7725, close: 23.82137, time: 1129842000, volume: 1966400},
		{high: 24.49, low: 23.76, open: 23.875, close: 24.289, time: 1130101200, volume: 1229700},
		{high: 24.59877, low: 23.70501, open: 24.389, close: 23.905, time: 1130187600, volume: 1537600},
		{high: 24.575, low: 23.85, open: 24.0575, close: 24.5, time: 1130274000, volume: 1482200},
		{high: 24.8, low: 23.81, open: 24.575, close: 23.85, time: 1130360400, volume: 1584300},
		{high: 24.59, low: 23.401, open: 23.75, close: 24.55223, time: 1130446800, volume: 1347900},
		{high: 25.39997, low: 24.6, open: 24.8395, close: 25.31271, time: 1130706000, volume: 1804300},
		{high: 25.87, low: 25.13753, open: 25.1795, close: 25.375, time: 1130792400, volume: 1762600},
		{high: 26.48, low: 25.411, open: 25.411, close: 26.34148, time: 1130878800, volume: 2537000},
	}
}

func GetTestSeries() *timeseries.TimeSeries {
	series := timeseries.New()
	for _, item := range GetTestCandles() {
		candle := timeseries.NewCandle(time.Unix(item.time, 0))
		candle.High = item.high
		candle.Low = item.low
		candle.Open = item.open
		candle.Close = item.close
		candle.Volume = item.volume

		err := series.AddCandle(candle)
		if err != nil {
			log.Fatalf("Failed to add candle: %s", err)
		}
	}

	return series
}
