// Фрагмент 2: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Find(data Predicate) (int, error) {
	n := data.Len()
	for i := 0; i < n; i++ {
		matched, err := data.Match(i)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}
