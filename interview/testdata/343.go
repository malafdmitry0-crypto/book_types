package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int8
	v := reflect.ValueOf(&n).Elem()
	fmt.Println(v.OverflowInt(257), v.OverflowInt(127))
}
