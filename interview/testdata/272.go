package main

import (
	"fmt"
)

func kind[T any](v T) int {
	switch v.(type) {
	case int:
		return 1
	}
	return 0
}
func main() {
	fmt.Println(kind(1))
}
