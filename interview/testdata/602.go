package main

import (
	"fmt"
)

func f() (n int) {
	defer func() { n += 4 }()
	return 3
}
func main() {
	fmt.Println(f())
}
