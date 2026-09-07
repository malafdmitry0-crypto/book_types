package main

import (
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.TypeFor[func(int) string]()
	v := reflect.MakeFunc(typ, func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(fmt.Sprint(in[0].Int() * 2))}
	})
	f := v.Interface().(func(int) string)
	fmt.Println(f(4))
}
