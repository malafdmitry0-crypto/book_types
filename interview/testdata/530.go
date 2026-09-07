package main

import (
	"fmt"
	"reflect"
)

type F func(int) int

func main() {
	v := reflect.MakeFunc(reflect.TypeFor[F](), func(in []reflect.Value) []reflect.Value { return []reflect.Value{in[0]} })
	fmt.Printf("%T %v\n", v.Interface(), v.Interface().(F)(7))
}
