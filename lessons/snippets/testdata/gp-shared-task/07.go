// Фрагмент 7: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ClosureReduce(n int, stepAt func(int) error) error {
	for i := 0; i < n; i++ {
		if err := stepAt(i); err != nil {
			return err
		}
	}
	return nil
}
