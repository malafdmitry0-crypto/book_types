package main

import (
	"fmt"
)

func main() {
	var a any = int32(7)
	switch v := a.(type) {
	case int32:
		fmt.Printf("%T\n", v)
	}
}
