// Фрагмент 1: book/chapters/gp-early-contracts.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Min(data Ordered) (int, error) {
	n := data.Len()
	if n == 0 {
		return -1, nil
	}
	best := 0
	for i := 1; i < n; i++ {
		comparison, err := data.Compare(i, best)
		if err != nil {
			return -1, err
		}
		if comparison < 0 {
			best = i
		}
	}
	return best, nil
}
