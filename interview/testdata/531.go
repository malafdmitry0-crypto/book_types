package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	v := reflect.MakeFunc(reflect.TypeFor[func() int](), func([]reflect.Value) []reflect.Value { return nil })
	fmt.Println("created")
	v.Call(nil)
}
