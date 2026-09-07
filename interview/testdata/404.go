package main

import (
	"fmt"
)

type F func() string

func (f F) Run() string {
	if f == nil {
		return "skipped"
	}
	return f()
}
func main() {
	var f F
	var r interface{ Run() string } = f
	fmt.Println(r == nil, r.Run())
}
