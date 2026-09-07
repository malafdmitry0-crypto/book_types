package main

import (
	"fmt"
	"math"
)

func main() {
	n := math.NaN()
	m := map[float64]int{n: 7}
	v, ok := m[n]
	delete(m, n)
	fmt.Println(v, ok, len(m))
	clear(m)
	fmt.Println(len(m))
}
