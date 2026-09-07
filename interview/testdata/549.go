package main

import (
	"fmt"
	"reflect"
)

func main() {
	src := struct{ V any }{7}
	var n int
	reflect.ValueOf(&n).Elem().Set(reflect.ValueOf(src).Field(0).Elem())
	fmt.Println(n)
}
