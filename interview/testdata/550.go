package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	v := reflect.New(reflect.TypeFor[struct{ N int }]())
	e := json.Unmarshal([]byte(`{"N":7}`), v.Interface())
	fmt.Println(v.Elem().Field(0).Int(), e == nil)
}
