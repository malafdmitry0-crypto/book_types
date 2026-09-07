package main

import (
	"fmt"
)

type N int

func main() {
	var x any = N(1)
	switch x.(type) {
	case int:
		fmt.Println("int")
	case N:
		fmt.Println("N")
	}
}
