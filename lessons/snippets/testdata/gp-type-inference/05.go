// Фрагмент 5: book/chapters/gp-type-inference.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
