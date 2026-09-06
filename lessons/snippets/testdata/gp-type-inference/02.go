// Фрагмент 2: book/chapters/gp-type-inference.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Zero[T any]() T {
	var zero T
	return zero
}
