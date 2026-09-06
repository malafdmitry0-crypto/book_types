// Фрагмент 9: book/chapters/gp-type-inference.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Clone[S ~[]E, E any](src S) S {
	if src == nil {
		return nil
	}
	out := make(S, len(src))
	copy(out, src)
	return out
}
