package main

import (
	"fmt"
	"math"
)

func main() {
	n := math.NaN()
	m := map[float64]int{n: 1}
	v, ok := m[n]
	fmt.Println(len(m), v, ok)
}
