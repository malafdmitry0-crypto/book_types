// Фрагмент 24: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Example_directIterator() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := IteratorFind(slices.Values(values), IsHalf)
	total, sumErr := IteratorReduce(slices.Values(values), zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
