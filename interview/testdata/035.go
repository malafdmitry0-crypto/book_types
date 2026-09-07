package main

import (
	"fmt"
)

func main() {
	var z complex64 = 3 + 4i
	fmt.Printf("%T %T\n", real(z), imag(z))
}
