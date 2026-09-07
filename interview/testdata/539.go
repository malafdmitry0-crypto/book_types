package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var f func()
	v := reflect.ValueOf(f)
	fmt.Println(v.IsValid(), v.IsNil())
	v.Call(nil)
}
