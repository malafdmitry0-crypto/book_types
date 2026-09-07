package main

import (
	"fmt"
	"math"
)

func main() {
	z := math.Copysign(0, -1)
	fmt.Println(z == 0, math.Signbit(z), math.Signbit(float64(float32(z))))
}
