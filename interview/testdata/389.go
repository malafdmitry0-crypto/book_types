package main

import (
	"fmt"
)

func main() {
	s := []int{1}
	var a any = (*[1]int)(s)
	s = append(s, 2)
	s[0] = 9
	fmt.Println(a.(*[1]int)[0])
}
