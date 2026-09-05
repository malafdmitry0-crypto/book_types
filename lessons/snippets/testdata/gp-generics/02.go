// Фрагмент 2: book/chapters/gp-08-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func All[E any](src []E, pred func(E) bool) bool {
	for _, item := range src {
		if !pred(item) {
			return false
		}
	}
	return true
}
