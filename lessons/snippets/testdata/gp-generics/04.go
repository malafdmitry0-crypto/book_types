// Фрагмент 4: book/chapters/gp-08-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Map[E, R any](src []E, transform func(E) R) []R {
	if src == nil {
		return nil
	}
	out := make([]R, len(src))
	for i, item := range src {
		out[i] = transform(item)
	}
	return out
}
