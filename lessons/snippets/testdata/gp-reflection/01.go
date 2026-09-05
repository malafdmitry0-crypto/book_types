// Фрагмент 1: book/chapters/gp-06-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
index, err := FindReflect(users, adult)
all, err := AllReflect(users, active)
raw, err := MapReflect(users, name)
// Каждый err нужно проверить до использования результата.
