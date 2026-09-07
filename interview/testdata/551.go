package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	v := reflect.Zero(reflect.TypeFor[struct{ N int }]())
	e := json.Unmarshal([]byte(`{"N":7}`), v.Interface())
	fmt.Println(e != nil)
}
