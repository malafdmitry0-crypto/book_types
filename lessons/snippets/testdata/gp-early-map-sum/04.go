// Фрагмент 4: book/chapters/gp-early-map-sum.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ExampleMap() {
	values := []fraction.Fraction{
		{Numerator: 2, Denominator: 4},
		{Numerator: 6, Denominator: 3},
	}
	texts := make([]string, len(values))
	toText := func(value fraction.Fraction) (string, error) {
		if err := value.Validate(); err != nil {
			return "", err
		}
		return value.String(), nil
	}
	job := fractionText{Source: values, Target: texts, Transform: toText}

	err := algorithms.Map(job)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	// Результат записан в texts через адаптер.
	fmt.Println("Строки:", texts)
	// Output: Строки: [1/2 2]
}
