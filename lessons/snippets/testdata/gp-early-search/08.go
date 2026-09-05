// Фрагмент 8: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ExampleAll() {
	values := []fraction.Fraction{
		{Numerator: 1, Denominator: 3},
		{Numerator: 2, Denominator: 4},
	}
	zero := fraction.Fraction{Denominator: 1}
	isPositive := func(value fraction.Fraction) (bool, error) {
		comparison, err := value.Compare(zero)
		return comparison > 0, err
	}
	query := fractionQuery{Values: values, Test: isPositive}

	positive, err := algorithms.All(query)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Все положительные:", positive)
	// Output: Все положительные: true
}
