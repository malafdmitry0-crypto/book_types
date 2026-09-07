package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[int]int{}
	reflect.ValueOf(m).SetMapIndex(reflect.ValueOf(1), reflect.Zero(reflect.TypeFor[int]()))
	v, ok := m[1]
	fmt.Println(v, ok)
}
