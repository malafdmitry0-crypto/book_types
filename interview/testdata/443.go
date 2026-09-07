package main

import (
	"fmt"
)

func all(s []int, p func(int) bool) bool {
	if p == nil {
		panic("nil predicate")
	}
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	fmt.Println(all(nil, nil))
}
