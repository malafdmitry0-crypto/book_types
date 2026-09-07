package main

import (
	"fmt"
	"math"
)

func main() {
	m := map[float64]int{0: 1}
	m[math.Copysign(0, -1)] = 2
	fmt.Println(len(m), m[0])
}
