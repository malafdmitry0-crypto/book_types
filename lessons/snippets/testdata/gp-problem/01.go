// Фрагмент 1: book/chapters/gp-01-problem.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
half := fraction.Fraction{Numerator: 1, Denominator: 2}
other := fraction.Fraction{Numerator: 2, Denominator: 4}

sameFields := half == other // false
sameValue := half.Equal(other) // true
