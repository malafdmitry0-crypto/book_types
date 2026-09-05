// Фрагмент 6: book/chapters/gp-09-constraints.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Equaler[E any] interface {
    Equal(E) bool
}

func ContainsEqual[E Equaler[E]](src []E, target E) bool {
    return Find(src, func(item E) bool { return item.Equal(target) }) >= 0
}
