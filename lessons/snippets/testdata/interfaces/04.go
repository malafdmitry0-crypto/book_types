// Фрагмент 4: book/chapters/07-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "fmt"
func describe(x any) string {
    switch v := x.(type) {
    case nil:
        return "nil interface"
    case int:
        return fmt.Sprintf("int: %d", v)
    case string:
        return "string: " + v
    case float32, float64:
        return fmt.Sprintf("float: %v", v) // здесь v имеет тип any
    default:
        return fmt.Sprintf("other: %T", v)
    }
}
