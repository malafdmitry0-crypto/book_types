// Фрагмент 8: book/chapters/20-cpp-memory.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
c := Counter{N: 10}
var a any = c
var b any = &c
c.N = 99
// a.(Counter).N == 10
// b.(*Counter).N == 99
