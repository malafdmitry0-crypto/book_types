package main

import (
	"fmt"
)

type N int

func (n N) M() int         { return int(n) }
func Call[T ~int](x T) int { return x.M() }
func main() {
	fmt.Println(Call(N(1)))
}
