package main

import (
	"fmt"
)

func main() {
	var p *[3]int
	sum := 0
	for i := range p {
		sum += i
	}
	fmt.Println(sum)
}
