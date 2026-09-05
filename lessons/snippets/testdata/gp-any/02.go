// Фрагмент 2: book/chapters/gp-05-any.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func AllAny(src []interface{}, pred func(interface{}) bool) bool {
	for _, item := range src {
		if !pred(item) {
			return false
		}
	}
	return true
}
