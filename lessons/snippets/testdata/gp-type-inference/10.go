// Фрагмент 10: book/chapters/gp-type-inference.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
copied := generic.Clone(ages{3, 9, 5}) // S = ages из аргумента, E = int из ограничения ~[]E
fmt.Printf("%T %v %d\n", copied, copied, copied.Oldest()) // generic_test.ages [3 9 5] 9
