// Фрагмент 1: book/chapters/gp-11-iterators.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// Концептуальная форма iter.Seq[E]:
// type Seq[E any] func(yield func(E) bool)

// import "slices"
for user := range slices.Values(users) {
    // user имеет тип User
    _ = user
}
