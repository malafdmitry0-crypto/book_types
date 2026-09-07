package main

import (
	"fmt"
)

func main() {
	var a any = (*int)(nil)
	var b any = (*int)(nil)
	fmt.Println(a == b)
}
