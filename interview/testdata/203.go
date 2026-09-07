package main

import (
	"fmt"
)

func main() {
	var m map[int]int
	delete(m, 1)
	clear(m)
	fmt.Println(m == nil)
}
