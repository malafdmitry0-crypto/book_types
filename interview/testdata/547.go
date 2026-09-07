package main

import (
	"fmt"
	"reflect"
)

type N int

func main() {
	var box struct{ V any }
	v := reflect.ValueOf(&box).Elem().Field(0)
	v.Set(reflect.ValueOf(N(3)))
	fmt.Printf("%T %v\n", box.V, box.V)
}
