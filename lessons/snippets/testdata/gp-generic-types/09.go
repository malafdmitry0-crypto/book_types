// Фрагмент 9: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func (s Set[K]) Has(v K) bool {
	_, ok := s.m[v]
	return ok
}
