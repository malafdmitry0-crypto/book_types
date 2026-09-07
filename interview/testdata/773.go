package main

import (
	"fmt"
	"maps"
)

func main() {
	x := 1
	a := map[string]*int{"x": &x}
	b := maps.Clone(a)
	*b["x"] = 7
	delete(b, "x")
	fmt.Println(*a["x"], len(a), len(b))
}
