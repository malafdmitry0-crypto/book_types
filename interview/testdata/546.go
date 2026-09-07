package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := reflect.ValueOf([]any{nil})
	v := s.Index(0)
	fmt.Println(v.IsNil(), v.Elem().IsValid())
}
