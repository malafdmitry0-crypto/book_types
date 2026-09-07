package main

import (
	"fmt"
)

type F func() int

func (f *F) Read() int { return (*f)() }
func main() {
	f := F(func() int { return 1 })
	var r interface{ Read() int } = &f
	f = func() int { return 2 }
	fmt.Println(r.Read())
}
