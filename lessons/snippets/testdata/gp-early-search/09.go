// Фрагмент 9: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Any(data Predicate) (bool, error) {
	index, err := Find(data)
	return index >= 0, err
}
