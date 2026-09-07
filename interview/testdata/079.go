package main

import (
	"fmt"
)

func main() {
	x := 1
	{
		x := "one"
		fmt.Printf("%T ", x)
	}
	fmt.Printf("%T\n", x)
}
