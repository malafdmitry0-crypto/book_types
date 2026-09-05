// Фрагмент 12: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
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
