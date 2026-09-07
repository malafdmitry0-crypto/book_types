package main

import (
	"fmt"
	"sort"
)

type Row struct {
	K    int
	Name string
}
type Rows []Row

func (s Rows) Len() int           { return len(s) }
func (s Rows) Less(i, j int) bool { return s[i].K < s[j].K }
func (s Rows) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func main() {
	s := Rows{{1, "a"}, {0, "x"}, {1, "b"}}
	sort.Stable(s)
	fmt.Println(s)
}
