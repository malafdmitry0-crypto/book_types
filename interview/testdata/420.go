package main

import (
	"fmt"
)

type F func(...int) int

func (f F) Apply(x []int) int { return f(x...) }
func main() {
	f := F(func(x ...int) int { return len(x) })
	var r interface{ Apply([]int) int } = f
	fmt.Println(r.Apply([]int{1, 2, 3}))
}
