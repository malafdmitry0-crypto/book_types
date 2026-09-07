package main

import (
	"reflect"
)

func id[T any](v T) T { return v }
func main() {
	_ = reflect.ValueOf(id)
}
