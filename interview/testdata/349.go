package main

import (
	"fmt"
	"reflect"
)

func main() {
	var p *int
	v := reflect.ValueOf(p).Elem()
	fmt.Println(v.IsValid())
}
