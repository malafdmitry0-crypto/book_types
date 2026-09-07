package main

import (
	"fmt"
)

func pair() (int, int) { return 3, 4 }
func f(a, b int) int   { return a + b }
func main() {
	fmt.Println(f(pair()))
}
