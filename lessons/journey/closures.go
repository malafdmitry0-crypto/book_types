package journey

// Для Filter замыкание само записывает выбранный элемент.
func ClosureFilter(n int, selectAt func(int) error) error {
	for i := 0; i < n; i++ {
		if err := selectAt(i); err != nil {
			return err
		}
	}
	return nil
}

// Для Reduce замыкание хранит типизированный аккумулятор.
func ClosureReduce(n int, stepAt func(int) error) error {
	for i := 0; i < n; i++ {
		if err := stepAt(i); err != nil {
			return err
		}
	}
	return nil
}
