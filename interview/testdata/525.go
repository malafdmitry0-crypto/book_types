package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	f := reflect.ValueOf(func([]int) {})
	f.CallSlice([]reflect.Value{reflect.ValueOf([]int{1})})
}
