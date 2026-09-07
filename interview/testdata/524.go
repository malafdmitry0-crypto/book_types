package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	f := reflect.ValueOf(func(xs ...int) {})
	f.Call([]reflect.Value{reflect.ValueOf([]int{1, 2})})
}
