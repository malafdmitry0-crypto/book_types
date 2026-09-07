package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := reflect.ValueOf(func(v any) string { return fmt.Sprintf("%T", v) })
	out := f.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(out[0].String())
}
