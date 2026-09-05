// Фрагмент 3: book/chapters/gp-03-closures.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "sort"
values := []int{2, 5, 9, 12}
index := sort.Search(len(values), func(i int) bool {
    return values[i] >= 8
}) // 2
