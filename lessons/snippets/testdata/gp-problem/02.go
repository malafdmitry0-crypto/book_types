// Фрагмент 2: book/chapters/gp-01-problem.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// Такое объявление допустимо, но наши типы его НЕ реализуют.
type Adder interface {
    Add(interface{}) (interface{}, error)
}
