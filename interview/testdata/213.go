package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	m := map[any]int{}
	delete(m, []int{})
}
