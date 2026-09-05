// Фрагмент 4: book/chapters/13-modern-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Converter struct{}

func (Converter) ToInt64[T ~int | ~int32 | ~int64](x T) int64 {
    return int64(x)
}

// n := (Converter{}).ToInt64(int32(42)) // Go 1.27+
