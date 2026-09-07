package main

import (
	"fmt"
	"reflect"
)

func main() {
	var p *int
	t := reflect.TypeOf(p)
	fmt.Println(t == nil, t.Kind(), t.Elem())
}
