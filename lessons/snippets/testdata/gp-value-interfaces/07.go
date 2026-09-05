// Фрагмент 7: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
f := fraction.Fraction{Numerator: 1, Denominator: 2}

text, err := DescribeUnknown(f) // Работает; соответствие проверится внутри.
fmt.Println(text, err)         // 1/2 <nil>

text = DescribeStringer(f)     // Работает; Fraction уже удовлетворяет Stringer.
fmt.Println(text)              // 1/2
