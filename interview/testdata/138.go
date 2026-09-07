package main

import (
	"fmt"
)

type T int

func (T) F(int)    {}
func (T) F(string) {}
func main() {
	fmt.Println(T(0))
}
