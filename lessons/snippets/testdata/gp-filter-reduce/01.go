// Фрагмент 1: book/chapters/gp-filter-reduce.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func GenericFilter[E any](src []E, test func(E) (bool, error)) ([]E, error) {
	if src == nil {
		return nil, nil
	}
	out := make([]E, 0, len(src))
	for _, value := range src {
		matched, err := test(value)
		if err != nil {
			return out, err
		}
		if matched {
			out = append(out, value)
		}
	}
	return out, nil
}
