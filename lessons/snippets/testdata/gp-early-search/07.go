// Фрагмент 7: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func All(data Predicate) (bool, error) {
	n := data.Len()
	for i := 0; i < n; i++ {
		matched, err := data.Match(i)
		if err != nil {
			return false, err
		}
		if !matched {
			return false, nil
		}
	}
	return true, nil
}
