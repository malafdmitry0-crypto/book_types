// Фрагмент 5: book/chapters/gp-08-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
names := Map(users, name) // []string
ages := Map(users, func(u User) int { return u.Age }) // []int

// Map(users, func(s string) int { return len(s) })
// Ошибка компиляции: []User и параметр string не согласованы.
