package main

import (
	"fmt"
)

type S struct {
	_ []int
	X int
}

func main() {
	a := S{}
	fmt.Println(a == a)
}
