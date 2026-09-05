// Фрагмент 6: book/chapters/gp-11-iterators.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "slices"
lengths := MapSeq(
    MapSeq(slices.Values(users), name),
    func(s string) int { return len(s) },
)
allLong := AllSeq(lengths, func(n int) bool { return n >= 4 })
