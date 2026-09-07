package main

import (
	"fmt"
)

type C interface {
	int
	string
}

func f[T C](v T) T { return v }
func main() {
	fmt.Println(f(1))
}
