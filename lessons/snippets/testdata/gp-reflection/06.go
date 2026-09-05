// Фрагмент 6: book/chapters/gp-06-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
raw, err := MapReflect(users, name)
if err != nil {
    // Обработать несовместимый вызов.
    return
}
names, ok := raw.([]string)
if !ok {
    // Нарушено ожидание вызывающего кода о типе результата.
    return
}
