package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := 7
	p := &x
	json.Unmarshal([]byte(`null`), &p)
	fmt.Println(p == nil, x)
}
