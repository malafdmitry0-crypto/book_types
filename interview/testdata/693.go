package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1}
	v := reflect.ValueOf(s)
	v.Index(0).SetInt(9)
	fmt.Println(v.CanSet(), s)
}
