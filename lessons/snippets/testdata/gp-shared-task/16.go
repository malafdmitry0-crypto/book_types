// Фрагмент 16: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ReflectFind(src, test interface{}) (int, error) {
	s, f, err := checkPredicate(src, test)
	if err != nil {
		return -1, err
	}
	for i := 0; i < s.Len(); i++ {
		r := f.Call([]reflect.Value{s.Index(i)})
		if err := reflectedError(r[1]); err != nil {
			return -1, err
		}
		if r[0].Bool() {
			return i, nil
		}
	}
	return -1, nil
}
