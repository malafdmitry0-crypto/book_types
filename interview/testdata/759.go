package main

import (
	"context"
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	context.WithValue(context.Background(), []int{1}, 7)
}
