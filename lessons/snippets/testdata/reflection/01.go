// Фрагмент 1: book/chapters/08-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type User struct {
    Age int
}

u := User{Age: 30}
var x any = u
t := reflect.TypeOf(x)
v := reflect.ValueOf(x)
