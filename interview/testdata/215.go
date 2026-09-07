package main

import (
	"fmt"
	"math"
)

func main() {
	n := math.NaN()
	m := map[float64]int{}
	m[n] = 1
	m[n] = 2
	fmt.Println(len(m))
}
