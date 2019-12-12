package indicator

// Описывает индерфейс индикатора. Единственный метод получает
// идентификатор торговой свечи и возвращает подсчитанное значение индикатора
type Indicator interface {
	Calculate(index int) float64
}
