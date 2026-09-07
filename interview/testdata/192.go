package main

import (
	"fmt"
)

func main() {
	var p *[7]int
	fmt.Println(len(p), cap(p))
}
