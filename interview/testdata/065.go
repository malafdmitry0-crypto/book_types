package main

import (
	"fmt"
)

func main() {
	a := struct {
		X int `json:"x"`
	}{1}
	b := (*struct {
		X int `json:"y"`
	})(&a)
	b.X = 9
	fmt.Println(a.X)
}
