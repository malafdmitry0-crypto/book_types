package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0)
	fmt.Println((*[0]int)(s) == nil)
}
