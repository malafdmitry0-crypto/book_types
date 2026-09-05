// Фрагмент 3: book/chapters/gp-filter-reduce.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ExampleGenericReduce() {
	values := []fraction.Fraction{{Numerator: 1, Denominator: 2}, {Numerator: 1, Denominator: 3}}
	text, err := journey.GenericReduce(values, "", func(total string, value fraction.Fraction) (string, error) {
		if err := value.Validate(); err != nil {
			return "", err
		}
		return strings.TrimPrefix(total+", "+value.String(), ", "), nil
	})
	fmt.Println(text, err)
	// Output: 1/2, 1/3 <nil>
}
