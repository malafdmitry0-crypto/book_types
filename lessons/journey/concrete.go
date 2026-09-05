package journey

import "gotypes/fraction"

func ConcreteFind(src []fraction.Fraction, test func(fraction.Fraction) (bool, error)) (int, error) {
	for i, value := range src {
		matched, err := test(value)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}
func ConcreteFilter(src []fraction.Fraction, test func(fraction.Fraction) (bool, error)) ([]fraction.Fraction, error) {
	if src == nil {
		return nil, nil
	}
	out := make([]fraction.Fraction, 0, len(src))
	for _, value := range src {
		matched, err := test(value)
		if err != nil {
			return out, err
		}
		if matched {
			out = append(out, value)
		}
	}
	return out, nil
}
func ConcreteReduce(src []fraction.Fraction, initial fraction.Fraction, step func(fraction.Fraction, fraction.Fraction) (fraction.Fraction, error)) (fraction.Fraction, error) {
	total := initial
	for _, value := range src {
		next, err := step(total, value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}
