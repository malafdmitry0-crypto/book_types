package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(nil)
	fmt.Println(v.IsValid(), v.Kind())
}
