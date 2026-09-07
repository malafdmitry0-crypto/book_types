package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	_ = reflect.MapOf(reflect.TypeFor[[]int](), reflect.TypeFor[int]())
}
