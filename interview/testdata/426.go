package main

import (
	"fmt"
	"sort"
)

func main() {
	s := sort.IntSlice{3, 1, 2}
	sort.Sort(sort.Reverse(sort.Reverse(s)))
	fmt.Println(s)
}
