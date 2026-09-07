package main

import (
	"fmt"
)

func main() {
	seq := func(yield func(int) bool) {
		defer fmt.Print("end ")
		if !yield(1) {
			return
		}
		yield(2)
	}
	for x := range seq {
		fmt.Print(x, " ")
		break
	}
	fmt.Println("done")
}
