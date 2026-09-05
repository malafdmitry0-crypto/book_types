// Фрагмент 5: book/chapters/gp-03-closures.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var texts []string
if values != nil {
    texts = make([]string, len(values))
}
err := MapIndex(len(values), func(i int) error {
    if err := values[i].Validate(); err != nil {
        return err
    }
    texts[i] = values[i].String()
    return nil
})
