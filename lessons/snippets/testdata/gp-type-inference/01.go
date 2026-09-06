// Фрагмент 1: book/chapters/gp-type-inference.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
users := []user{{"Аня", 30}, {"Борис", 17}}
names := generic.Map(users, userName)              // E = user, R = string выведены
same := generic.Map[user, string](users, userName) // то же самое явно
