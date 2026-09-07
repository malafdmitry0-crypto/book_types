package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeFor[error]()
	fmt.Println(t.Kind(), t.NumMethod())
}
