// Фрагмент 4: book/chapters/gp-interface-limits.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
fractions := []fraction.Fraction{
    {Numerator: 1, Denominator: 2},
}
// PrintAll(fractions) // Не компилируется: []Fraction не является []Stringer.
