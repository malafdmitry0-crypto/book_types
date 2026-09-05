package journey

func GenericFind[E any](src []E, test func(E) (bool, error)) (int, error) {
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
func GenericFilter[E any](src []E, test func(E) (bool, error)) ([]E, error) {
	if src == nil {
		return nil, nil
	}
	out := make([]E, 0, len(src))
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
func GenericReduce[E, A any](src []E, initial A, step func(A, E) (A, error)) (A, error) {
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
