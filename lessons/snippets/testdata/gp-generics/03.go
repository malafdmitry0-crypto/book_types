// Фрагмент 3: book/chapters/gp-08-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
index := Find(users, adult)
all := All(users, active)
textIndex := Find([]string{"go", "generic"}, func(s string) bool {
    return len(s) > 2
})
