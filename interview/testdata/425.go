package main

import (
	"fmt"
	"sort"
)

func main() {
	s := sort.IntSlice{2, 1, 3}
	sort.Sort(sort.Reverse(s))
	fmt.Println(s)
}
