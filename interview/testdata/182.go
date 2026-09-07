package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2}
	p := (*[2]int)(s)
	p[0] = 9
	fmt.Println(s)
}
