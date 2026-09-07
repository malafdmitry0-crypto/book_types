package main

import (
	"fmt"
)

type IDs []int

func Copy[S ~[]E, E any](s S) S { return append(S(nil), s...) }
func main() {
	x := Copy(IDs{2, 5})
	fmt.Printf("%T %v", x, x)
}
