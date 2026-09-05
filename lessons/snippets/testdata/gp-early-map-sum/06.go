// Фрагмент 6: book/chapters/gp-early-map-sum.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Sum(data Summable) error {
	n := data.Len()
	for i := 0; i < n; i++ {
		if err := data.Add(i); err != nil {
			return err
		}
	}
	return nil
}
