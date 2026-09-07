package main

import (
	"fmt"
)

func main() {
	var a any = (*int)(nil)
	p, ok := a.(*int)
	fmt.Println(p == nil, ok)
}
