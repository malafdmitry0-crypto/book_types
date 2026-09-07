package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	f := reflect.ValueOf(func(int64) {})
	f.Call([]reflect.Value{reflect.ValueOf(int32(1))})
}
