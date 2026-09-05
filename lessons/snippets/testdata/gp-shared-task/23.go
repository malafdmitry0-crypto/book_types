// Фрагмент 23: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func IteratorFind[E any](src iter.Seq[E], test func(E) (bool, error)) (int, error) {
	index := 0
	for value := range src {
		matched, err := test(value)
		if err != nil {
			return -1, err
		}
		if matched {
			return index, nil
		}
		index++
	}
	return -1, nil
}

func IteratorReduce[E, A any](src iter.Seq[E], initial A, step func(A, E) (A, error)) (A, error) {
	total := initial
	for value := range src {
		next, err := step(total, value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}
