package main

import (
	"fmt"
	"reflect"
)

type N int

func main() {
	x := N(3)
	v := reflect.ValueOf(&x).Elem()
	v.SetInt(8)
	fmt.Printf("%T %v", x, x)
}
