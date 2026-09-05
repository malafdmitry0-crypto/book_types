// Фрагмент 1: book/chapters/gp-04-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Predicate interface {
    Len() int
    Match(i int) (bool, error)
}
