// Фрагмент 8: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Example_directClosure() {
	values := exampleFractions()
	index, findErr := closures.FindIndex(len(values), func(i int) (bool, error) { return IsHalf(values[i]) })
	total := fraction.Fraction{Denominator: 1}
	sumErr := ClosureReduce(len(values), func(i int) error {
		next, err := total.Add(values[i])
		if err != nil {
			return err
		}
		total = next
		return nil
	})
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
