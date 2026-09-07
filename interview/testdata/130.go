package main

import (
	"fmt"
)

type T int

func (t T) Get() T { return t }
func main() {
	f := (*T).Get
	t := T(7)
	fmt.Println(f(&t))
}
