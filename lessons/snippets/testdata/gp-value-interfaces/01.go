// Фрагмент 1: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Stringer interface {
    String() string
}

func PrintValue(value Stringer) {
    fmt.Println(value.String())
}
