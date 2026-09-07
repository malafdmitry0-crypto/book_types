package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]int{"x": 7}
	reflect.ValueOf(m).SetMapIndex(reflect.ValueOf("x"), reflect.Value{})
	fmt.Println(len(m))
}
