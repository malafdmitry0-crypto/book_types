package main

import (
	"fmt"
	"reflect"
)

type I interface{ F() }

func main() {
	m, _ := reflect.TypeFor[I]().MethodByName("F")
	fmt.Println(m.Type.NumIn())
}
