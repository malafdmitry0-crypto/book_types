package main

import (
	"fmt"
)

func main() {
	var a any = []int{1}
	var b any = []string{"1"}
	fmt.Println(a == b)
}
