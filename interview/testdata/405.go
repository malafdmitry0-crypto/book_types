package main

import (
	"fmt"
)

type F func() int

func (f F) Read() int { return f() }
func main() {
	n := 1
	f := F(func() int { return n })
	var r interface{ Read() int } = f
	n = 9
	fmt.Println(r.Read())
}
