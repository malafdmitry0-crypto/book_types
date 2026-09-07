package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1}
	v := reflect.ValueOf(s)
	fmt.Println(v.CanSet(), v.Index(0).CanSet())
	v.Index(0).SetInt(9)
	fmt.Println(s)
}
