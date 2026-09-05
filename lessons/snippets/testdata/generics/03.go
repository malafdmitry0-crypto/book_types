// Фрагмент 3: book/chapters/11-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Map[S ~[]E, E, R any](src S, convert func(E) R) []R {
    if src == nil {
        return nil
    }
    dst := make([]R, len(src))
    for i, item := range src {
        dst[i] = convert(item)
    }
    return dst
}

// result := Map([]int{1, 2}, func(n int) int64 { return int64(n) })
