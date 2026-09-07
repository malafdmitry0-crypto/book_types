package main

import (
	"fmt"
)

type F func()

func (f F) Run() { f() }
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var r interface{ Run() } = F(func() {})
	m := map[any]int{}
	m[r] = 1
}
