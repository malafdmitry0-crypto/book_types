// Фрагмент 1: book/chapters/gp-08-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Find[E any](src []E, pred func(E) bool) int {
	for i, item := range src {
		if pred(item) {
			return i
		}
	}
	return -1
}
