package main

import (
	"fmt"
	"reflect"
)

func Map[A, B any](s []A, f func(A) B) (out []B) {
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}
func main() {
	v := reflect.MakeFunc(reflect.TypeFor[func(int) string](), func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(fmt.Sprint(in[0].Int()))}
	})
	out := Map([]int{1, 2}, v.Interface().(func(int) string))
	fmt.Println(out)
}
