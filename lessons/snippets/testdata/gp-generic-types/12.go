// Фрагмент 12: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
texts := containers.MapStack(s, fraction.Fraction.String)
fmt.Printf("%T %v\n", texts, texts) // containers.Stack[string] [1/2 1/2]
