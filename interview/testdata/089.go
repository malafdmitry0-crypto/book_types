package main

import (
	"fmt"
)

func main() {
	var a any = (*int)(nil)
	switch a.(type) {
	case nil:
		fmt.Println("nil")
	case *int:
		fmt.Println("pointer")
	}
}
