package main

import (
	"fmt"
	"math"
)

func main() {
	x := math.NaN()
	a, b := &x, &x
	fmt.Println(a == b, *a == *b)
}
