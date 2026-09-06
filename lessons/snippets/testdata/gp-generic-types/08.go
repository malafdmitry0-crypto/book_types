// Фрагмент 8: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func (s Set[K]) Add(v K) { s.m[v] = struct{}{} }
