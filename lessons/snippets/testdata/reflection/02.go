// Фрагмент 2: book/chapters/08-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
u := User{Age: 30}
direct := reflect.ValueOf(u)
field := reflect.ValueOf(&u).Elem().FieldByName("Age")

fmt.Println(direct.CanSet(), field.CanSet()) // false true
field.SetInt(31)
fmt.Println(u.Age) // 31
