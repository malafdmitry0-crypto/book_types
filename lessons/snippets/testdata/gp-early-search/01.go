// Фрагмент 1: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Predicate interface {
    Len() int
    Match(i int) (bool, error)
}
