package main

import (
	"fmt"
	"math"
)

func main() {
	x := math.NaN()
	fmt.Println(x == x, float32(x) == float32(x))
}
