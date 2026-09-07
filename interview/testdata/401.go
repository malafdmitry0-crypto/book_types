package main

import (
	"fmt"
)

type Handler interface{ Handle(int) int }
type HandlerFunc func(int) int

func (f HandlerFunc) Handle(n int) int { return f(n) }
func use(h Handler, n int) int         { return h.Handle(n) }
func main() {
	h := HandlerFunc(func(n int) int { return n * 2 })
	fmt.Println(use(h, 3))
}
