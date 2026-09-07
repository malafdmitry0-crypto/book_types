package main

import (
	"fmt"
	"reflect"
)

type N int

func (n *N) Inc() { *n++ }
func main() {
	n := N(1)
	v := reflect.ValueOf(&n).Elem()
	fmt.Println(v.CanAddr(), v.MethodByName("Inc").IsValid(), v.Addr().MethodByName("Inc").IsValid())
}
