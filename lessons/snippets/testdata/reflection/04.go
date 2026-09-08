// Фрагмент 4: book/chapters/08-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
v := reflect.ValueOf(User{Age: 30})
back := v.Interface() // статический тип any, динамический User
u, ok := back.(User)
fmt.Println(u.Age, ok) // 30 true
