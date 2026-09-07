package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int8
	reflect.ValueOf(&n).Elem().SetInt(257)
	fmt.Println(n)
}
