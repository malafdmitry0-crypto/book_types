package main

import (
	"fmt"
)

func helper() { fmt.Println(recover()) }
func main() {
	defer func() { fmt.Println(recover()) }()
	defer func() { helper() }()
	panic("boom")
}
