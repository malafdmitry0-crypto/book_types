package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	delete(m, "x")
	clear(m)
	fmt.Println(m == nil, len(m))
}
