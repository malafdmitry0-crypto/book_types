// Фрагмент 2: book/chapters/gp-06-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func checkTransform(src, fn interface{}) (reflect.Value, reflect.Value, error) {
	s, f := reflect.ValueOf(src), reflect.ValueOf(fn)
	if !s.IsValid() || s.Kind() != reflect.Slice {
		return s, f, fmt.Errorf("source must be a slice")
	}
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return s, f, fmt.Errorf("callback must be a non-nil function")
	}
	t := f.Type()
	if t.IsVariadic() || t.NumIn() != 1 || t.NumOut() != 1 || t.In(0) != s.Type().Elem() {
		return s, f, fmt.Errorf("callback must have signature func(%v) R", s.Type().Elem())
	}
	return s, f, nil
}
