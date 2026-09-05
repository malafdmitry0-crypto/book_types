// Фрагмент 4: book/chapters/gp-06-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func AllReflect(src, pred interface{}) (bool, error) {
	s, f, err := checkPredicate(src, pred)
	if err != nil {
		return false, err
	}
	for i := 0; i < s.Len(); i++ {
		if !f.Call([]reflect.Value{s.Index(i)})[0].Bool() {
			return false, nil
		}
	}
	return true, nil
}
