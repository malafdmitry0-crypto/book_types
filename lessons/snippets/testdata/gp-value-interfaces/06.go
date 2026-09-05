// Фрагмент 6: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func DescribeUnknown(value interface{}) (string, error) {
    stringer, ok := value.(Stringer)
    if !ok {
        return "", fmt.Errorf("тип %T не предоставляет String() string", value)
    }
    return stringer.String(), nil
}

func DescribeStringer(value Stringer) string {
    return value.String()
}
