// Фрагмент 7: book/chapters/gp-11-iterators.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "slices"
names := slices.Collect(MapSeq(slices.Values(users), name))
