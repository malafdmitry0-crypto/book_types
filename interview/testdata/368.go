package main

import (
	"fmt"
)

func main() {
	b := []byte("go")
	s := string(b)
	b[0] = 'n'
	fmt.Println(s, string(b))
}
