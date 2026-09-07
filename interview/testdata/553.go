package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1}
	v := reflect.ValueOf(&s).Elem()
	v.Set(reflect.Append(v, reflect.ValueOf(2)))
	fmt.Println(s)
}
