// Фрагмент 2: book/chapters/07-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var x any = int(42)
n := x.(int)              // 42
wide, ok := x.(int64)      // 0, false
// panicValue := x.(int64) // panic

if n, ok := x.(int); ok {
    converted := int64(n)  // assertion, затем числовая конверсия
    _ = converted
}
