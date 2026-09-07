package main

import (
	"fmt"
)

func main() {
	a := struct {
		X int `json:"x"`
	}{1}
	b := struct {
		X int `json:"y"`
	}(a)
	fmt.Println(b)
}
