package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	v := reflect.ValueOf([]int{1})
	_ = v.Convert(reflect.TypeFor[[2]int]())
}
