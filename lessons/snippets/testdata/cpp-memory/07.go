// Фрагмент 7: book/chapters/20-cpp-memory.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
c := Counter{}
c.Inc() // сокращение для (&c).Inc(): c — адресуемая переменная

var i Incrementer = &c // подходит *Counter
// var j Incrementer = c // ошибка: Counter не реализует Incrementer
_ = i
