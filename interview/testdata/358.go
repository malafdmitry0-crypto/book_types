package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	s := make([]int, 1, 3)
	reflect.ValueOf(s).SetLen(2)
}
