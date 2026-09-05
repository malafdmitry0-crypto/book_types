// Фрагмент 1: book/chapters/gp-02-concrete.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func FindFraction(values []fraction.Fraction, target fraction.Fraction) (int, error) {
    if err := target.Validate(); err != nil {
        return -1, err
    }
    for i, value := range values {
        if err := value.Validate(); err != nil {
            return -1, err
        }
        if value.Equal(target) {
            return i, nil
        }
    }
    return -1, nil
}
