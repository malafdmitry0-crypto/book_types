// Фрагмент 11: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func MapStack[T, R any](s Stack[T], f func(T) R) Stack[R] {
	out := Stack[R]{items: make([]R, len(s.items))}
	for i, item := range s.items {
		out.items[i] = f(item)
	}
	return out
}
