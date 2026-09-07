package main

import (
	"fmt"
)

func main() {
	var b []byte
	s := string(b)
	fmt.Println(s == "", len(s))
}
