package main

import (
	"fmt"
)

type M map[string]int

func main() {
	a := M{"x": 1}
	var b map[string]int = a
	b["x"] = 2
	fmt.Println(a["x"])
}
