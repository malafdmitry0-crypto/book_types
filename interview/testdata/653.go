package main

import (
	"fmt"
)

type I interface{ Sum([]int) int }
type N struct{}

func (N) Sum(...int) int { return 1 }
func main() {
	var x I = N{}
	fmt.Println(x.Sum([]int{1}))
}
