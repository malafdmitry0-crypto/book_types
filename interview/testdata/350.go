package main

import (
	"fmt"
	"reflect"
)

type S struct{ x int }

func main() {
	s := S{x: 7}
	v := reflect.ValueOf(&s).Elem().Field(0)
	fmt.Println(v.CanAddr(), v.CanSet(), v.CanInterface(), v.Int())
}
