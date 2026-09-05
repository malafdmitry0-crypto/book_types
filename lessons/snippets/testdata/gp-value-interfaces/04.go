// Фрагмент 4: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
f, ok := value.(fraction.Fraction)
if ok {
    fmt.Println(f.Numerator, f.Denominator)
}
