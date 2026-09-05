package journey

func AnyFind(src []interface{}, test func(interface{}) (bool, error)) (int, error) {
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
func AnyFilter(src []interface{}, test func(interface{}) (bool, error)) ([]interface{}, error) {
	if src == nil {
		return nil, nil
	}
	out := make([]interface{}, 0, len(src))
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
func AnyReduce(src []interface{}, initial interface{}, step func(interface{}, interface{}) (interface{}, error)) (interface{}, error) {
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
