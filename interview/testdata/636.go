package main

import (
	"fmt"
)

type N int

func main() {
	var x any = N(4)
	switch v := x.(type) {
	case N:
		fmt.Printf("%T %v", v, v+1)
	default:
		fmt.Print("other")
	}
}
