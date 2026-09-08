// Фрагмент 4: book/chapters/20-cpp-memory.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func makeNumber() *int {
    x := 42
    return &x // корректно: объект может пережить вызов функции
}
