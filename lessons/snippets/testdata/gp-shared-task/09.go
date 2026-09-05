// Фрагмент 9: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Predicate interface {
	Len() int
	Match(i int) (bool, error)
}

func Find(data Predicate) (int, error) {
	n := data.Len()
	for i := 0; i < n; i++ {
		matched, err := data.Match(i)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}

type Summable interface {
	Len() int
	Add(i int) error
}

func Sum(data Summable) error {
	n := data.Len()
	for i := 0; i < n; i++ {
		if err := data.Add(i); err != nil {
			return err
		}
	}
	return nil
}
