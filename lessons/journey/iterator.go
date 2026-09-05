package journey

import "iter"

// Позиция относится к порядку обхода Seq. Для slices.Values она совпадает с индексом.
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

// Единственная ошибка выдаётся последней. yield(false) прекращает источник.
func IteratorFilter[E any](src iter.Seq[E], test func(E) (bool, error)) iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		for value := range src {
			matched, err := test(value)
			if err != nil {
				var zero E
				yield(zero, err)
				return
			}
			if matched && !yield(value, nil) {
				return
			}
		}
	}
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
