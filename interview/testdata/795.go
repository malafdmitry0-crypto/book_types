package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	seq := func(yield func(int) bool) {
		yield(1)
		yield(2)
	}
	for range seq {
		break
	}
}
