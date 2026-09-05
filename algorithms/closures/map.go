package closures

// MapIndex выполняет apply слева направо; при ошибке сохраняет уже выполненные действия.
// n должен соответствовать доступным callback индексам. Отката нет.
func MapIndex(n int, apply func(int) error) error {
	for i := 0; i < n; i++ {
		if err := apply(i); err != nil {
			return err
		}
	}
	return nil
}
