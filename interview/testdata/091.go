package main

import (
	"fmt"
)

func main() {
	var a any = int32(7)
	switch v := a.(type) {
	case int, int32:
		_, ok := v.(int32)
		fmt.Println(ok)
	}
}
