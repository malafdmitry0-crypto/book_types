package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var m map[string]int
	m["x"] = 1
}
