package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [1]int{1}
	v := reflect.ValueOf(a)
	p := reflect.ValueOf(&a).Elem()
	p.Index(0).SetInt(9)
	fmt.Println(v.Index(0).Int(), p.Index(0).Int(), v.Index(0).CanSet())
}
