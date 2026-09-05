// Фрагмент 1: book/chapters/gp-00-origins.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Stringer interface {
    String() string
}

func Describe(value Stringer) string {
    return value.String()
}
