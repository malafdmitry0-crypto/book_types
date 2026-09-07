package main

import (
	"fmt"
)

func main() {
	var f func() any = func() int { return 1 }
	fmt.Println(f())
}
