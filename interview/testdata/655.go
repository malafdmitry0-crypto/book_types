package main

import (
	"fmt"
)

type N struct{}

func (n *N) OK() bool { return n == nil }

type Wrapper struct{ *N }
type I interface{ OK() bool }

func main() {
	var w Wrapper
	var i I = w
	fmt.Println(i.OK())
}
