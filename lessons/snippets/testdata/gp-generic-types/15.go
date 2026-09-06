// Фрагмент 15: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Fill[T any](c Container[T], values ...T) {
	for _, v := range values {
		c.Push(v)
	}
}
