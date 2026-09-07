package main

import (
	"fmt"
	"math"
)

func main() {
	n := math.NaN()
	m := map[float64]int{n: 1}
	delete(m, n)
	fmt.Print(len(m), " ")
	clear(m)
	fmt.Println(len(m))
}
