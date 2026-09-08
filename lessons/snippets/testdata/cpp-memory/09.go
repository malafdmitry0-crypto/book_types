// Фрагмент 9: book/chapters/20-cpp-memory.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var p *Counter = nil
var i Incrementer = p
var empty Incrementer
// p == nil: true
// i == nil: false
// empty == nil: true
