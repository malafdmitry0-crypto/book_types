package main

import (
	"fmt"
)

type Node []Node

func main() {
	var n Node
	n = append(n, nil)
	fmt.Println(len(n), n[0] == nil)
}
