// Фрагмент 4: book/chapters/gp-11-iterators.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func FindSeq[E any](src iter.Seq[E], pred func(E) bool) (E, bool) {
	for item := range src {
		if pred(item) {
			return item, true
		}
	}
	var zero E
	return zero, false
}
