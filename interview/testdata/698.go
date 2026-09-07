package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	t := reflect.TypeOf(func() int { return 0 })
	v := reflect.MakeFunc(t, func([]reflect.Value) []reflect.Value { return []reflect.Value{reflect.ValueOf(int64(3))} })
	v.Call(nil)
}
