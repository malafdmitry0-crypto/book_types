package main

import (
	"fmt"
)

type K string

func main() {
	a := map[K]int{}
	fmt.Println(map[string]int(a))
}
