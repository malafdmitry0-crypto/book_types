package main

import (
	"fmt"
)

type F func()

func (f F) Run() { f() }
func main() {
	n := 0
	h := F(func() { n++ })
	fmt.Print(n, " ")
	h.Run()
	fmt.Println(n)
}
