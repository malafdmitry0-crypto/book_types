package main

import (
	"fmt"
)

func main() {
	var a any
	p, ok := a.(*int)
	fmt.Println(p == nil, ok)
}
