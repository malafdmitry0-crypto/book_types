package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(struct{ X any }{[]int{}})
	fmt.Println(v.Type().Comparable(), v.Comparable())
}
