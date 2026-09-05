// Фрагмент 18: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Example_directReflection() {
	values := exampleFractions()
	index, findErr := ReflectFind(values, IsHalf)
	zero := fraction.Fraction{Denominator: 1}
	raw, sumErr := ReflectReduce(values, zero, Add)
	total, unboxErr := unbox(raw)
	if unboxErr != nil {
		fmt.Println(unboxErr)
		return
	}
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
