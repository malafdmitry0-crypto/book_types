package main

import (
	"fmt"
)

type T int

func (t T) Get() int { return int(t) }
func main() {
	var i interface{ Get() int } = T(1)
	f := i.Get
	i = T(2)
	fmt.Println(f(), i.Get())
}
