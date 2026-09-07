package main

import (
	"fmt"
	"sort"
)

type Ints []int

func (s Ints) Len() int           { return len(s) }
func (s Ints) Less(i, j int) bool { return s[i] < s[j] }
func (s Ints) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func main() {
	s := []int{3, 1, 2}
	sort.Sort(Ints(s))
	fmt.Println(s)
}
