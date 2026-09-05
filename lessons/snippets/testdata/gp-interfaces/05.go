// Фрагмент 5: book/chapters/gp-04-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Ordered interface {
    Len() int
    Compare(i, j int) (int, error)
}

type Sortable interface {
    Ordered
    Swap(i, j int)
}
