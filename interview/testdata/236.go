package main

import (
	"fmt"
)

type F func() int
type G func() int

func main() {
	n := 0
	f := F(func() int {
		n++
		return n
	})
	g := G(f)
	fmt.Println(f(), g())
}
