package main

import (
	"fmt"
)

func factory() func() {
	fmt.Println("factory")
	return func() { fmt.Println("deferred") }
}
func main() {
	defer factory()()
	fmt.Println("body")
}
