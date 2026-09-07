package main

import (
	"fmt"
)

func all(s []int, p func(int) bool) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(all(nil, nil))
}
