package main

import (
	"fmt"
)

type R interface{ Read() int }
type RW interface {
	Read() int
	Write() int
}
type N struct{}

func (N) Read() int  { return 3 }
func (N) Write() int { return 7 }
func main() {
	var r R = N{}
	rw, ok := r.(RW)
	fmt.Println(ok, rw.Write())
}
