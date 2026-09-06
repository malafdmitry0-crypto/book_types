// Фрагмент 2: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}
