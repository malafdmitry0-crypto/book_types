// Фрагмент 6: book/chapters/gp-08-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Addable[T any] interface {
    Add(T) (T, error)
}

func SumValues[T Addable[T]](values []T, initial T) (T, error) {
    total := initial
    for _, value := range values {
        next, err := total.Add(value)
        if err != nil {
            return total, err
        }
        total = next
    }
    return total, nil
}
