// Фрагмент 2: book/chapters/gp-09-constraints.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// imports: "cmp", "slices"
slices.Sort([]int{3, 1, 2})
slices.SortFunc(users, func(a, b User) int {
    return cmp.Compare(a.Age, b.Age)
})
