package main

import (
	"fmt"
)

type I interface{ M() int }
type N int

func (n N) M() int { return int(n) }
func main() {
	var x I = N(2)
	f := x.M
	x = N(9)
	fmt.Println(f(), x.M())
}
