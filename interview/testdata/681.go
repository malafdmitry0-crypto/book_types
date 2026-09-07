package main

import (
	"fmt"
	"reflect"
)

func main() {
	var p *int
	fmt.Println(reflect.TypeOf(nil) == nil, reflect.TypeOf(p).Kind())
}
