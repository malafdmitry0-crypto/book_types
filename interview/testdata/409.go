package main

import (
	"fmt"
)

type F func() int

func (f F) Read() any { return f() }
func main() {
	var r interface{ Read() any } = F(func() int { return 3 })
	fmt.Printf("%T %v\n", r.Read(), r.Read())
}
