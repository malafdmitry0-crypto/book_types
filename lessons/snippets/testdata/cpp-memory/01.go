// Фрагмент 1: book/chapters/20-cpp-memory.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var a int = 10
// var b int64 = a // ошибка компиляции
var b int64 = int64(a)
_ = b
