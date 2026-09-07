package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[int]int{1: 2}
	v := reflect.ValueOf(m).MapIndex(reflect.ValueOf(1))
	fmt.Println(v.CanAddr(), v.CanSet())
}
