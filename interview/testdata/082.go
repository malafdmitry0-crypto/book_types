package main

import (
	"fmt"
)

func main() {
	var e error
	var a any = e
	fmt.Println(a == nil)
}
