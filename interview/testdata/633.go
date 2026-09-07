package main

import (
	"fmt"
)

func main() {
	n := 0
	for i := 0; i < 3; i++ {
		switch i {
		case 1:
			break
		}
		n++
	}
	fmt.Println(n)
}
