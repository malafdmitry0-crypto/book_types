package main

import (
	"fmt"
)

type S struct {
	_ int
	X int
}

func main() {
	a := S{X: 1}
	b := S{X: 1}
	fmt.Println(a == b)
}
