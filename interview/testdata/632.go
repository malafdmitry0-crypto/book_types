package main

import (
	"fmt"
)

func main() {
	n := 1
	switch n {
	case 1:
		fmt.Print("a ")
		fallthrough
	case 99:
		fmt.Println("b")
	}
}
