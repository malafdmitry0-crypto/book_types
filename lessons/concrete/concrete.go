package concrete

import (
	"gotypes/fraction"
	"gotypes/money"
)

func FindFraction(values []fraction.Fraction, target fraction.Fraction) (int, error) {
	if err := target.Validate(); err != nil {
		return -1, err
	}
	for i, value := range values {
		if err := value.Validate(); err != nil {
			return -1, err
		}
		if value.Equal(target) {
			return i, nil
		}
	}
	return -1, nil
}
func FindFractionFunc(values []fraction.Fraction,
	test func(fraction.Fraction) (bool, error)) (int, error) {
	for i, value := range values {
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
func SumMoney(values []money.Money, initial money.Money) (money.Money, error) {
	total := initial
	for _, value := range values {
		next, err := total.Add(value)
		if err != nil {
			return total, err // в этом варианте возвращаем уже накопленный префикс
		}
		total = next
	}
	return total, nil
}
func MapFractionText(values []fraction.Fraction, transform func(fraction.Fraction) (string, error)) ([]string, error) {
	if values == nil {
		return nil, nil
	}
	out := make([]string, len(values))
	for i, v := range values {
		text, err := transform(v)
		if err != nil {
			return out, err
		}
		out[i] = text
	}
	return out, nil
}
