package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var f func()
	defer f()
	fmt.Println("body")
}
