package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(struct{ x int }{7}).Field(0)
	fmt.Println(v.Int(), v.CanInterface())
}
