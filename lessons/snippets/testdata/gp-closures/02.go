// Фрагмент 2: book/chapters/gp-03-closures.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
half := fraction.Fraction{Numerator: 1, Denominator: 2}
index, err := FindIndex(len(values), func(i int) (bool, error) {
    if err := values[i].Validate(); err != nil {
        return false, err
    }
    return values[i].Equal(half), nil
})
