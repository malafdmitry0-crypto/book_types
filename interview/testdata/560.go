package main

import (
	"fmt"
	"reflect"
)

type N int

func types[T any](v T) (reflect.Kind, reflect.Type) {
	return reflect.TypeFor[T]().Kind(), reflect.TypeOf(v)
}
func main() {
	var x any = N(3)
	fmt.Println(types(x))
}
