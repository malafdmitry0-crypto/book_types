// Фрагмент 3: book/chapters/gp-interface-limits.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func PrintAll(values []Stringer) {
    for _, value := range values {
        fmt.Println(value.String())
    }
}
