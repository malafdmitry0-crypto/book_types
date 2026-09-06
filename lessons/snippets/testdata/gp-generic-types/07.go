// Фрагмент 7: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func NewSet[K comparable]() Set[K] {
	return Set[K]{m: make(map[K]struct{})}
}
