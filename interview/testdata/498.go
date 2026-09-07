package main

import (
	"fmt"
)

type Pred func(int) bool

func Keep[F ~func(E) bool, E any](f F) F { return f }
func main() {
	var f Pred = func(n int) bool { return n > 0 }
	g := Keep(f)
	fmt.Printf("%T %v\n", g, g(1))
}
