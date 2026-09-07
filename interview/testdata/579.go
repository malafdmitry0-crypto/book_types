package main

import (
	"fmt"
)

type Getter interface{ Get(string) int }
type Store struct{ M map[string]int }

func (s Store) Get(k string) int { return s.M[k] }
func main() {
	a := Store{map[string]int{"x": 1}}
	var i Getter = a
	var j Getter = i
	a.M["x"] = 9
	fmt.Println(i.Get("x"), j.Get("x"))
}
