package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	reflect.ValueOf(nil).IsNil()
}
