package main

import (
	"fmt"
	"reflect"
)

type N int

func main() {
	x := N(3)
	v := reflect.ValueOf(&x).Elem()
	v.Set(reflect.ValueOf(8).Convert(v.Type()))
	fmt.Println(x)
}
