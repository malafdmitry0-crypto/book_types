// Фрагмент 5: book/chapters/06-defined-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
ch := make(chan int)
var recv <-chan int = ch
var send chan<- int = ch
explicit := (<-chan int)(ch)
// both := (chan int)(recv) // восстановить оба направления нельзя
