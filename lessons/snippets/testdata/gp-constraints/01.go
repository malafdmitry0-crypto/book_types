// Фрагмент 1: book/chapters/gp-09-constraints.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Contains[E comparable](src []E, target E) bool {
	return Find(src, func(item E) bool { return item == target }) >= 0
}
