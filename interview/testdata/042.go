package main

import (
	"fmt"
)

type A int

func (a A) Double() A { return a * 2 }

type B A

func main() {
	var x B = 4
	fmt.Println(x.Double())
}
