// Фрагмент 2: book/chapters/11-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// Не компилируется: any не гарантирует допустимость конверсии.
// func Cast[To, From any](x From) To { return To(x) }

func Assert[T any](x any) (T, bool) {
    value, ok := x.(T)
    return value, ok
}

func IsString[T any](x T) bool {
    _, ok := any(x).(string)
    return ok
}
