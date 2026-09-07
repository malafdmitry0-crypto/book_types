package main

import (
	"fmt"
)

func Reduce[E, A any](s []E, a A, f func(A, E) A) A { return a }
func main() {
	fmt.Println(Reduce([]int{1}, 0, func(a int64, v int) int32 { return int32(a) + int32(v) }))
}
