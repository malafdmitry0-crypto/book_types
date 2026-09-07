package main

import (
	"fmt"
)

func main() {
	a, b := 7, 7
	fmt.Println(&a == &b, *(&a) == *(&b))
}
