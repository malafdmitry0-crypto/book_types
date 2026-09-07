package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := reflect.ValueOf([]int32{7})
	f := reflect.ValueOf(func(n int64) int64 { return n + 1 })
	v := s.Index(0).Convert(f.Type().In(0))
	fmt.Println(f.Call([]reflect.Value{v})[0].Int())
}
