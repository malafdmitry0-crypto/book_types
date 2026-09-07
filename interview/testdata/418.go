package main

import (
	"fmt"
)

type A func() int
type B func() int

func (f A) Get() int  { return f() }
func (f B) Load() int { return f() }
func main() {
	n := 0
	a := A(func() int {
		n++
		return n
	})
	b := B(a)
	fmt.Println(a.Get(), b.Load())
}
