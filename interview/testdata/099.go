package main

import (
	"fmt"
)

func main() {
	var s []int
	var a any = s
	fmt.Println(a == nil)
}
