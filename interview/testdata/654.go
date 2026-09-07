package main

import (
	"fmt"
)

type I = int
type F interface{ Get() int }
type N struct{}

func (N) Get() I { return 6 }
func main() {
	var f F = N{}
	fmt.Println(f.Get())
}
