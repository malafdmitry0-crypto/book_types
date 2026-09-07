package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(map[string]int{}).MapIndex(reflect.ValueOf("x"))
	fmt.Println(v.IsValid(), v.Kind())
}
