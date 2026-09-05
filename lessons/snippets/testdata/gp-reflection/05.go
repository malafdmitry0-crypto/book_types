// Фрагмент 5: book/chapters/gp-06-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func MapReflect(src, transform interface{}) (interface{}, error) {
	s, f, err := checkTransform(src, transform)
	if err != nil {
		return nil, err
	}
	target := reflect.SliceOf(f.Type().Out(0))
	if s.IsNil() {
		return reflect.Zero(target).Interface(), nil
	}
	out := reflect.MakeSlice(target, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		out.Index(i).Set(f.Call([]reflect.Value{s.Index(i)})[0])
	}
	return out.Interface(), nil
}
