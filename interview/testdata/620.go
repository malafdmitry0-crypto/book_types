package main

import (
	"fmt"
)

func main() {
	defer func() {
		n := recover()
		fmt.Println(n + 1)
	}()
	panic(2)
}
