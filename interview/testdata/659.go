package main

import (
	"fmt"
)

type H interface{ Do(int) int }
type F func(int) int

func (f F) Do(x int) int { return f(x) }
func main() {
	var h H = F(func(x int) int { return x * 3 })
	fmt.Println(h.Do(4))
}
