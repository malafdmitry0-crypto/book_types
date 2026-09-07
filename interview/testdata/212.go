package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	m := map[any]int{}
	fmt.Println(m[[]int{1}])
}
