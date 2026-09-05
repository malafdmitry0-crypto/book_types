// Фрагмент 5: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var value interface{} = fraction.Fraction{Numerator: 1, Denominator: 2}
// value.String() // Не компилируется: interface{} не объявляет методов.
