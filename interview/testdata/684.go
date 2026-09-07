package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x any = 7
	reflect.ValueOf(&x).Elem().Set(reflect.ValueOf("go"))
	fmt.Printf("%T %v", x, x)
}
