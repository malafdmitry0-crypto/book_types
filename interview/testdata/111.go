package main

import (
	"fmt"
)

func main() {
	var f func()
	var a any = f
	fmt.Println(f == nil, a == nil)
}
