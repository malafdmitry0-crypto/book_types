// Фрагмент 10: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func SortedKeys[K cmp.Ordered](s Set[K]) []K {
	keys := make([]K, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
