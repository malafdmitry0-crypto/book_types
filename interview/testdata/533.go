package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.MakeFunc(reflect.TypeFor[func() any](), func([]reflect.Value) []reflect.Value { return []reflect.Value{reflect.ValueOf(7)} })
	f := v.Interface().(func() any)
	fmt.Printf("%T %v\n", f(), f())
}
