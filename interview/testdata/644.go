package main

import (
	"fmt"
)

type N struct{}

func (*N) M() {}

type I interface{ M() }

func main() {
	n := N{}
	n.M()
	var x I = n
	fmt.Println(x)
}
