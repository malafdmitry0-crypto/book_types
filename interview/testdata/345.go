package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int64
	v := reflect.ValueOf(&n).Elem()
	v.Set(reflect.ValueOf(int32(7)).Convert(v.Type()))
	fmt.Println(n)
}
