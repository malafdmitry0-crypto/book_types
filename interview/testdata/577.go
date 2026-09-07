package main

import (
	"fmt"
)

type counter struct{ N int }

func (c *counter) Get() int       { return c.N }
func (c *counter) Reset()         { c.N = 0 }
func New() interface{ Get() int } { return &counter{7} }
func main() {
	x := New()
	fmt.Printf("%T %v\n", x, x.Get())
	_, ok := x.(interface{ Reset() })
	fmt.Println(ok)
}
