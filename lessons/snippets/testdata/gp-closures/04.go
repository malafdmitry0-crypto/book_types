// Фрагмент 4: book/chapters/gp-03-closures.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func MapIndex(n int, apply func(int) error) error {
    for i := 0; i < n; i++ {
        if err := apply(i); err != nil {
            return err
        }
    }
    return nil
}
