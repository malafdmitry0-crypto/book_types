// Фрагмент 5: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
index, err := FindIndex(len(values), func(i int) (bool, error) {
    return isHalf(values[i])
})
