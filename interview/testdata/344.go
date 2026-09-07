package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var n int64
	reflect.ValueOf(&n).Elem().Set(reflect.ValueOf(int32(1)))
}
