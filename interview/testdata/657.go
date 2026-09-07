package main

import (
	"fmt"
)

type A struct{}

func (A) M() int { return 1 }

type B struct{}

func (B) M() int { return 2 }

type C struct {
	A
	B
}
type I interface{ M() int }

func main() {
	var i I = C{}
	fmt.Println(i.M())
}
