// Фрагмент 13: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func (s Stack[T]) String() string { return fmt.Sprint(s.items) }
