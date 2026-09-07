package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := reflect.ValueOf(map[string]any{"x": nil})
	v := m.MapIndex(reflect.ValueOf("x"))
	fmt.Println(v.IsValid(), v.IsNil())
}
