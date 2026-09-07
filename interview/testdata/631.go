package main

import (
	"fmt"
)

func main() {
	n := 0
	f := func() bool {
		n++
		return true
	}
	switch {
	case f():
		fmt.Print("a ")
	case f():
		fmt.Print("b ")
	}
	fmt.Println(n)
}
