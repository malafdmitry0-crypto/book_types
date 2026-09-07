package main

import (
	"fmt"
)

type N int

func main() {
	var a any = int(7)
	n, ok := a.(N)
	fmt.Println(n, ok)
}
