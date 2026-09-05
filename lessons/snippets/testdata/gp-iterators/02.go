// Фрагмент 2: book/chapters/gp-11-iterators.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func MapSeq[E, R any](src iter.Seq[E], transform func(E) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range src {
			if !yield(transform(item)) {
				return
			}
		}
	}
}
