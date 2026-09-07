package main

import (
	"fmt"
	"reflect"
)

func main() {
	n := 1
	v := reflect.MakeFunc(reflect.TypeFor[func() int](), func([]reflect.Value) []reflect.Value { return []reflect.Value{reflect.ValueOf(n)} })
	n = 9
	fmt.Println(v.Interface().(func() int)())
}
