// Фрагмент 3: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func IsHalf(value fraction.Fraction) (bool, error) {
	if err := value.Validate(); err != nil {
		return false, err
	}
	return value.Equal(fraction.Fraction{Numerator: 1, Denominator: 2}), nil
}

func Add(a, b fraction.Fraction) (fraction.Fraction, error) { return a.Add(b) }
