// Фрагмент 17: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ReflectReduce(src, initial, step interface{}) (interface{}, error) {
	s, a, f := reflect.ValueOf(src), reflect.ValueOf(initial), reflect.ValueOf(step)
	if !s.IsValid() || s.Kind() != reflect.Slice || !a.IsValid() {
		return initial, fmt.Errorf("source must be a slice and initial must have a concrete type")
	}
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return initial, fmt.Errorf("step must be a non-nil function")
	}
	t := f.Type()
	if t.IsVariadic() || t.NumIn() != 2 || t.NumOut() != 2 || t.In(0) != a.Type() || t.In(1) != s.Type().Elem() || t.Out(0) != a.Type() || t.Out(1) != errorType {
		return initial, fmt.Errorf("step must have signature func(A,E) (A,error)")
	}
	for i := 0; i < s.Len(); i++ {
		r := f.Call([]reflect.Value{a, s.Index(i)})
		if err := reflectedError(r[1]); err != nil {
			return a.Interface(), err
		}
		a = r[0]
	}
	return a.Interface(), nil
}
