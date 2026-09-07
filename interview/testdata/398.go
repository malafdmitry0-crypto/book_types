package main

import (
	"fmt"
)

func main() {
	s := func(yield func(int) bool) { fmt.Println(yield(1)) }
	for v := range s {
		fmt.Println(v)
		break
	}
}
