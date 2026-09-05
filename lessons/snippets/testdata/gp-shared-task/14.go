// Фрагмент 14: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Example_directAny() {
	values := exampleFractions()
	boxed := Box(values)
	index, findErr := AnyFind(boxed, AnyPredicate(IsHalf))
	zero := fraction.Fraction{Denominator: 1}
	raw, sumErr := AnyReduce(boxed, zero, anyStep(Add))
	total, unboxErr := unbox(raw)
	if unboxErr != nil {
		fmt.Println(unboxErr)
		return
	}
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
