// Фрагмент 14: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Container[T any] interface {
	Len() int
	Push(T)
}
