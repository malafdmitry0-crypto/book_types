package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	v := reflect.ValueOf(struct{ x int }{7}).Field(0)
	fmt.Println(v.Interface())
}
