package main

import (
	"fmt"
	"math"
)

func main() {
	a := struct{ F float64 }{math.NaN()}
	fmt.Println(a == a)
}
