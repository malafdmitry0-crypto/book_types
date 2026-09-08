// Фрагмент 3: book/chapters/21-static-dynamic.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var x any = 42
if n, ok := x.(int); ok {
    fmt.Println(n + 1) // 43: статический тип n — int
}

x = "hello"
n, ok := x.(int)
fmt.Println(n, ok) // 0 false: внутри уже string
