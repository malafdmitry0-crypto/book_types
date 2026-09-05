// Фрагмент 4: book/chapters/06-defined-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
src := []int{1, 2, 3}
// dst := []int64(src) // запрещено
dst := make([]int64, len(src))
for i, n := range src {
    dst[i] = int64(n)
}

// Не работают также []string → []any и map[string]int → map[string]int64.
