package main

import (
	"fmt"
)

func main() {
	b := []byte("go")
	m := map[string]int{string(b): 7}
	b[0] = 'n'
	fmt.Println(m["go"], m[string(b)], string(b))
}
