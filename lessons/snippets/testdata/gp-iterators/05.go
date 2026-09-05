// Фрагмент 5: book/chapters/gp-11-iterators.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func AllSeq[E any](src iter.Seq[E], pred func(E) bool) bool {
	for item := range src {
		if !pred(item) {
			return false
		}
	}
	return true
}
