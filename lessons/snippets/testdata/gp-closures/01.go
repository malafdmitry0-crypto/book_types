// Фрагмент 1: book/chapters/gp-03-closures.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func FindIndex(n int, test func(int) (bool, error)) (int, error) {
    for i := 0; i < n; i++ {
        matched, err := test(i)
        if err != nil {
            return -1, err
        }
        if matched {
            return i, nil
        }
    }
    return -1, nil
}
