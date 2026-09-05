// Фрагмент 3: book/chapters/04-assignment.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var p *int = nil
typedNil := (*int)(nil)
var s []int = nil
var m map[string]int = nil
var fn func() = nil
var ch chan int = nil
var x any = nil

n := 42
ptr := &n
value := *ptr
// _ = int(nil) // недопустимо
