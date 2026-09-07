package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	s := reflect.ValueOf([]int32{1})
	f := reflect.ValueOf(func(int64) int { return 1 })
	f.Call([]reflect.Value{s.Index(0)})
}
