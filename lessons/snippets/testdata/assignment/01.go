// Фрагмент 1: book/chapters/04-assignment.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
const answer = 42
var a int8 = answer
var b float64 = answer
var c complex128 = answer

inferred := 42 // int: тип по умолчанию
const typed int = 42
// var d int64 = typed // typed уже имеет тип int

// _ = uint8(300) // ошибка компиляции: константа не помещается
// _ = int(3.9)   // ошибка компиляции: константа не представима как int
_ = int(3.0)      // допустимо: значение константы целое
