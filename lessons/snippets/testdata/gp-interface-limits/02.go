// Фрагмент 2: book/chapters/gp-interface-limits.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type FractionAdder interface {
    Add(fraction.Fraction) (fraction.Fraction, error)
}

func AddFraction(left FractionAdder, right fraction.Fraction) (fraction.Fraction, error) {
    return left.Add(right)
}

var _ FractionAdder = fraction.Fraction{}
