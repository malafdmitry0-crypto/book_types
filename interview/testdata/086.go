package main

import (
	"fmt"
)

func main() {
	var a any = int32(7)
	n, ok := a.(int64)
	fmt.Println(n, ok)
}
