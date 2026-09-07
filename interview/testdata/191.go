package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var p *[3]int
	for _, v := range p {
		fmt.Println(v)
	}
}
