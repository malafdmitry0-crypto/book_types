package main

import (
	"fmt"
)

func main() {
	b := append([]byte{65}, "я"...)
	fmt.Println(b)
}
