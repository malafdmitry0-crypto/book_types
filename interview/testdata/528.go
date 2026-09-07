package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := reflect.ValueOf(func(v any) bool { return v == nil })
	out := f.Call([]reflect.Value{reflect.Zero(reflect.TypeFor[any]())})
	fmt.Println(out[0].Bool())
}
