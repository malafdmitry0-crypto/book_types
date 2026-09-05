// Фрагмент 2: book/chapters/gp-07-generation.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func MapUsersGenerated(src []User, fn func(User) string) []string {
	if src == nil {
		return nil
	}
	out := make([]string, len(src))
	for i, item := range src {
		out[i] = fn(item)
	}
	return out
}
