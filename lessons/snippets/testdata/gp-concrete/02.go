// Фрагмент 2: book/chapters/gp-02-concrete.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func FindFractionFunc(values []fraction.Fraction,
    test func(fraction.Fraction) (bool, error)) (int, error) {
    for i, value := range values {
        matched, err := test(value)
        if err != nil {
            return -1, err
        }
        if matched {
            return i, nil
        }
    }
    return -1, nil
}
