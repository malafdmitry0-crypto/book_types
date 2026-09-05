// Фрагмент 15: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var errorType = reflect.TypeOf((*error)(nil)).Elem()

func checkPredicate(src, test interface{}) (reflect.Value, reflect.Value, error) {
	s, f := reflect.ValueOf(src), reflect.ValueOf(test)
	if !s.IsValid() || s.Kind() != reflect.Slice {
		return s, f, fmt.Errorf("source must be a slice")
	}
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return s, f, fmt.Errorf("predicate must be a non-nil function")
	}
	t := f.Type()
	if t.IsVariadic() || t.NumIn() != 1 || t.NumOut() != 2 || t.In(0) != s.Type().Elem() || t.Out(0) != reflect.TypeOf(false) || t.Out(1) != errorType {
		return s, f, fmt.Errorf("predicate must have signature func(E) (bool, error)")
	}
	return s, f, nil
}

func reflectedError(v reflect.Value) error {
	if v.IsNil() {
		return nil
	}
	return v.Interface().(error)
}
