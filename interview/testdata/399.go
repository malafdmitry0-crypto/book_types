package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	s := func(yield func(int) bool) {
		yield(1)
		yield(2)
	}
	for range s {
		break
	}
}
