// Фрагмент 1: book/chapters/gp-practicum.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Example_mapComposition() {
	var events []string
	f := func(n int) int { events = append(events, fmt.Sprintf("f%d", n)); return n + 1 }
	g := func(n int) int { events = append(events, fmt.Sprintf("g%d", n)); return n * 2 }
	first := generic.Map(generic.Map([]int{1, 2}, f), g)
	fmt.Println(first, events)
	events = nil
	second := generic.Map([]int{1, 2}, func(n int) int { return g(f(n)) })
	fmt.Println(second, events)
	// Output:
	// [4 6] [f1 f2 g2 g3]
	// [4 6] [f1 g2 f2 g3]
}
