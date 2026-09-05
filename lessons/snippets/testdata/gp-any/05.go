// Фрагмент 5: book/chapters/gp-05-any.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
FindAny([]interface{}{User{Name: "Аня"}}, func(x interface{}) bool {
    return len(x.(string)) > 0 // panic при обработке User
})
