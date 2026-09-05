// Фрагмент 13: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Box(src []fraction.Fraction) []interface{} {
	if src == nil {
		return nil
	}
	out := make([]interface{}, len(src))
	for i, v := range src {
		out[i] = v
	}
	return out
}

func unbox(value interface{}) (fraction.Fraction, error) {
	f, ok := value.(fraction.Fraction)
	if !ok {
		return fraction.Fraction{}, fmt.Errorf("expected Fraction, got %T", value)
	}
	return f, nil
}

func AnyPredicate(test Predicate) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		f, err := unbox(value)
		if err != nil {
			return false, err
		}
		return test(f)
	}
}

func anyStep(step Step) func(interface{}, interface{}) (interface{}, error) {
	return func(a, b interface{}) (interface{}, error) {
		left, err := unbox(a)
		if err != nil {
			return nil, err
		}
		right, err := unbox(b)
		if err != nil {
			return nil, err
		}
		return step(left, right)
	}
}
