// Фрагмент 5: book/chapters/gp-09-constraints.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Named interface {
    Name() string
}

func Names[E Named](src []E) []string {
    return Map(src, func(item E) string { return item.Name() })
}
