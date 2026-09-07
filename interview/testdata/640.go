package main

import (
	"fmt"
)

func f() (n int) {
	if n := 8; n > 0 {
		return
	}
	return
}
func main() {
	fmt.Println(f())
}
