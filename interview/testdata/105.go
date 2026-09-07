package main

import (
	"fmt"
)

type S struct{ X []int }

func main() {
	a := S{}
	fmt.Println(a == a)
}
