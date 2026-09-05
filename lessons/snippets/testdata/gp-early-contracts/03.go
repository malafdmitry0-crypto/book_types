// Фрагмент 3: book/chapters/gp-early-contracts.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Sort(data Sortable) error {
	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			comparison, err := data.Compare(j, j-1)
			if err != nil {
				return err
			}
			if comparison >= 0 {
				break
			}
			data.Swap(j, j-1)
		}
	}
	return nil
}
