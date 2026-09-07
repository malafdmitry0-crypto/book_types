package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.MakeFunc(reflect.TypeFor[func(...int) int](), func(in []reflect.Value) []reflect.Value {
		fmt.Println(len(in), in[0].Kind())
		return []reflect.Value{reflect.ValueOf(in[0].Len())}
	})
	fmt.Println(v.Interface().(func(...int) int)(1, 2, 3))
}
