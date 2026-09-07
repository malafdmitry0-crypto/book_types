package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	src := struct{ V any }{7}
	var n int
	reflect.ValueOf(&n).Elem().Set(reflect.ValueOf(src).Field(0))
}
