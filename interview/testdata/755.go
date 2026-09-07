package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 0
	f := sync.OnceValue(func() int {
		n++
		return 0
	})
	fmt.Println(f(), f(), n)
}
