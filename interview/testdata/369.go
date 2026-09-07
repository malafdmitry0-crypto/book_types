package main

import (
	"fmt"
)

func main() {
	s := "go"
	b := []byte(s)
	b[0] = 'n'
	fmt.Println(s, string(b))
}
