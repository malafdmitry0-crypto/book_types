// Фрагмент 6: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ExampleFind() {
	// 1. Данные и условие поиска.
	values := []fraction.Fraction{
		{Numerator: 1, Denominator: 3},
		{Numerator: 2, Denominator: 4},
		{Numerator: 3, Denominator: 4},
	}
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	isHalf := func(value fraction.Fraction) (bool, error) {
		if err := value.Validate(); err != nil {
			return false, err
		}
		return value.Equal(half), nil
	}

	// 2. Адаптер связывает срез и условие с методами Len и Match.
	query := fractionQuery{Values: values, Test: isHalf}

	// 3. Алгоритм видит только интерфейс адаптера.
	index, err := algorithms.Find(query)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if index == -1 {
		fmt.Println("Совпадений нет")
		return
	}

	// 4. Найденное значение берём из исходного типизированного среза.
	fmt.Println("Индекс:", index)
	fmt.Println("Дробь:", values[index])
	// Output:
	// Индекс: 1
	// Дробь: 1/2
}
