package main

import (
	"fmt"
)

type R interface{ Read() }
type W interface{ Write() }
type T struct{}

func (T) Read()  {}
func (T) Write() {}
func main() {
	var r R = T{}
	_, ok := r.(W)
	fmt.Println(ok)
}
