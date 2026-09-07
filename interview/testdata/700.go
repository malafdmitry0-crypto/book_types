package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(struct{ X any }{[]int{1}})
	fmt.Println(v.Type().Comparable(), v.Comparable())
}
