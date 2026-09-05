// Фрагмент 1: book/chapters/11-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Real interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64
}

func Convert[To Real, From Real](x From) To {
    return To(x)
}

// result := Convert[int64](int32(42))
// To задан явно; From выводится из аргумента.
