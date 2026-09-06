// Фрагмент 4: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var s containers.Stack[fraction.Fraction] // нулевое значение пригодно
s.Push(fraction.Fraction{Numerator: 1, Denominator: 2})
s.Push(fraction.Fraction{Numerator: 1, Denominator: 3})
fmt.Println("Len:", s.Len(), "Stack:", s)
top, ok := s.Pop()
fmt.Println("Pop:", top, ok)
