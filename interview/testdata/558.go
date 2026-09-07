package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := reflect.ValueOf(map[string]int{})
	v := m.MapIndex(reflect.ValueOf("x"))
	fmt.Println(v.IsValid())
}
