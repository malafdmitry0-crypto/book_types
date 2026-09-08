// Фрагмент 6: book/chapters/20-cpp-memory.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
c := Counter{N: 10}
var i Incrementer = &c
i.Inc()
// c.N == 11
