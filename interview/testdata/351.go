package main

import (
	"fmt"
	"reflect"
)

type S struct{ x int }

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	v := reflect.ValueOf(S{x: 7}).Field(0)
	_ = v.Interface()
}
