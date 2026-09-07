package main

import (
	"fmt"
)

func main() {
	x := 1
	m := map[int]string{x: "a", x: "b"}
	fmt.Println(len(m), m[1])
}
