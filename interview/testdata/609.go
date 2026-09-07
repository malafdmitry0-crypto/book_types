package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover()) }()
	defer func() { panic("second") }()
	panic("first")
}
