package main

import (
	"fmt"
)

func main() {
	var f func()
	defer func() { fmt.Println(recover() != nil) }()
	f()
}
