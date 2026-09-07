package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[int]int{1: 2}
	reflect.ValueOf(m).SetMapIndex(reflect.ValueOf(1), reflect.Value{})
	fmt.Println(len(m))
}
