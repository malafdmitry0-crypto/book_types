// Фрагмент 5: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var values containers.Stack[fraction.Fraction]
var pointers containers.Stack[*fraction.Fraction]
half := fraction.Fraction{Numerator: 1, Denominator: 2}
values.Push(half)
pointers.Push(&half)
// values.Push(&half) // ошибка компиляции: *Fraction не является Fraction
// values = pointers  // ошибка компиляции: Stack[*Fraction] — другой тип
fmt.Printf("%T %T\n", values, pointers)
