// Фрагмент 6: book/chapters/gp-type-inference.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
fmt.Printf("%T %v\n", generic.Max(1, 2), generic.Max(1, 2))         // int 2
fmt.Printf("%T %v\n", generic.Max(1, 2.5), generic.Max(1, 2.5))     // float64 2.5
