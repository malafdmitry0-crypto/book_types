// Фрагмент 3: book/chapters/gp-05-any.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func MapAny(src []interface{}, transform func(interface{}) interface{}) []interface{} {
	if src == nil {
		return nil
	}
	out := make([]interface{}, len(src))
	for i, item := range src {
		out[i] = transform(item)
	}
	return out
}
