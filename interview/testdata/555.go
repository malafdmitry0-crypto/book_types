package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	a := make([]int64, 1)
	reflect.Copy(reflect.ValueOf(a), reflect.ValueOf([]int32{1}))
}
