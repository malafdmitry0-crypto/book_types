package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]int{"x": 7}
	v := reflect.ValueOf(m).MapIndex(reflect.ValueOf("x"))
	fmt.Println(v.CanAddr(), v.CanSet())
}
