package main

import (
	"fmt"
	"reflect"
)

func mapValue(s, f reflect.Value) reflect.Value {
	out := reflect.MakeSlice(reflect.SliceOf(f.Type().Out(0)), s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		out.Index(i).Set(f.Call([]reflect.Value{s.Index(i)})[0])
	}
	return out
}
func main() {
	out := mapValue(reflect.ValueOf([]int{1, 2}), reflect.ValueOf(func(n int) string { return fmt.Sprint(n) }))
	fmt.Printf("%T %v\n", out.Interface(), out.Interface())
}
