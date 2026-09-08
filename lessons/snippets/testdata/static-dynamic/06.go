// Фрагмент 6: book/chapters/21-static-dynamic.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var x any = Robot{}
s, ok := x.(Speaker)
if ok {
    fmt.Println(s.Speak()) // beep
}

x = 42
_, ok = x.(Speaker)
fmt.Println(ok) // false
