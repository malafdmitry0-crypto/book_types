// Фрагмент 3: book/chapters/gp-06-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func FindReflect(src, pred interface{}) (int, error) {
	s, f, err := checkPredicate(src, pred)
	if err != nil {
		return -1, err
	}
	for i := 0; i < s.Len(); i++ {
		if f.Call([]reflect.Value{s.Index(i)})[0].Bool() {
			return i, nil
		}
	}
	return -1, nil
}
