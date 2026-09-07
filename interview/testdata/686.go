package main

import (
	"fmt"
	"reflect"
)

type N int

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	x := N(3)
	reflect.ValueOf(&x).Elem().Set(reflect.ValueOf(8))
}
