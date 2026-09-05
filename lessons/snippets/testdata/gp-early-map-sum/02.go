// Фрагмент 2: book/chapters/gp-early-map-sum.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Map(data Mapping) error {
	n := data.Len()
	for i := 0; i < n; i++ {
		if err := data.Apply(i); err != nil {
			return err
		}
	}
	return nil
}
