// Фрагмент 11: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Example_directInterface() {
	values := exampleFractions()
	query := Search{Values: values, Test: IsHalf}
	index, findErr := algorithms.Find(query)
	sum := Fold{Values: values, Total: fraction.Fraction{Denominator: 1}, Step: Add}
	sumErr := algorithms.Sum(&sum)
	fmt.Println(index, findErr, sum.Total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
