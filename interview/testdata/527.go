package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	f := reflect.ValueOf(func(any) {})
	f.Call([]reflect.Value{reflect.ValueOf(nil)})
}
