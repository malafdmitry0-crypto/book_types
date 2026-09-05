// Фрагмент 1: book/chapters/gp-05-any.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func FindAny(src []interface{}, pred func(interface{}) bool) int {
	for i, item := range src {
		if pred(item) {
			return i
		}
	}
	return -1
}
