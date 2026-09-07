package main

import (
	"fmt"
)

type F func()

func (f F) Ready() bool { return f != nil }

type Checker interface{ Ready() bool }

func main() {
	var f F
	var c Checker = f
	fmt.Println(c == nil, c.Ready())
}
