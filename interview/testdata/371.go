package main

import (
	"fmt"
)

func main() {
	b := []byte("")
	fmt.Println(b == nil, len(b))
}
