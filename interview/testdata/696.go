package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	f := reflect.ValueOf(func(x int64) {})
	f.Call([]reflect.Value{reflect.ValueOf(3)})
}
