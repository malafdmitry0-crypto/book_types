// Фрагмент 1: book/chapters/gp-10-standard-library.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// Старый интерфейсный API: три метода у адаптера UsersByAge.
sort.Sort(UsersByAge(users))

// Go 1.8+: срез как interface{}, сравнение через индексы.
sort.Slice(users, func(i, j int) bool {
    return users[i].Age < users[j].Age
})

// Go 1.21+: типизированные элементы в callback.
slices.SortFunc(users, func(a, b User) int {
    return cmp.Compare(a.Age, b.Age)
})

// Для типов со встроенным порядком callback вообще не нужен.
slices.Sort([]int{3, 1, 2})
