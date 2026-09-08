// Фрагмент 2: book/chapters/20-cpp-memory.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
const n = 42
var a int32 = n
var b int64 = n
x := n // тип по умолчанию для этой целочисленной константы: int
_, _, _ = a, b, x
