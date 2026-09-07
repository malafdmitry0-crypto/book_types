package main

import (
	"fmt"
)

func main() {
	for i, r := range "aяb" {
		fmt.Printf("%d:%U ", i, r)
	}
	fmt.Println()
}
