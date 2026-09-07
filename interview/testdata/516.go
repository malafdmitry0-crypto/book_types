package main

import (
	"fmt"
)

type Sink[E any] interface{ Put(E) }
type SinkFunc[E any] func(E)

func (f SinkFunc[E]) Put(e E) { f(e) }

type Logged[E any] struct{ Next Sink[E] }

func (l Logged[E]) Put(e E) {
	fmt.Println("before")
	l.Next.Put(e)
	fmt.Println("after")
}
func main() {
	s := Logged[int]{Next: SinkFunc[int](func(n int) { fmt.Println("sink", n) })}
	s.Put(7)
}
