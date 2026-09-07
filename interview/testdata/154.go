package main

import (
	"fmt"
)

type I interface{ F() int }
type T struct{}

func (T) F() int { return 7 }

type O struct{ I }

func main() {
	o := O{T{}}
	fmt.Println(o.F())
}
