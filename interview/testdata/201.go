package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	v, ok := m["x"]
	fmt.Println(v, ok, len(m))
}
