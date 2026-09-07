package main

import (
	"fmt"
)

func main() {
	a := struct {
		X int
		Y string
	}{1, "a"}
	fmt.Println(struct {
		Y string
		X int
	}(a))
}
