// Фрагмент 2: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
f := fraction.Fraction{Numerator: 1, Denominator: 2}
m := money.Money{Amount: 250, Currency: "USD"}
PrintValue(f) // 1/2
PrintValue(m) // 250 USD (minor units)
