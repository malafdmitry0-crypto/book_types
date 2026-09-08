// Фрагмент 2: book/chapters/21-static-dynamic.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var x any = 42
fmt.Printf("%T\n", x) // int

x = "hello"
fmt.Printf("%T\n", x) // string

// _ = x + 1 // ошибка компиляции даже сразу после x = 42
