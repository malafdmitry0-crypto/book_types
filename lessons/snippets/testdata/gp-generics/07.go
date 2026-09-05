// Фрагмент 7: book/chapters/gp-08-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func MapValues[E, R any](values []E, transform func(E) (R, error)) ([]R, error) {
    if values == nil {
        return nil, nil
    }
    out := make([]R, len(values))
    for i, value := range values {
        result, err := transform(value)
        if err != nil {
            return out, err
        }
        out[i] = result
    }
    return out, nil
}
