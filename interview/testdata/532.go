package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	v := reflect.MakeFunc(reflect.TypeFor[func() int64](), func([]reflect.Value) []reflect.Value { return []reflect.Value{reflect.ValueOf(int32(1))} })
	v.Call(nil)
}
