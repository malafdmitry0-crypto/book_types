package main

import (
	"fmt"
)

type N int

func main() {
	var x any = N(4)
	switch v := x.(type) {
	case N, int:
		_, ok := v.(N)
		fmt.Println(ok)
	}
}
