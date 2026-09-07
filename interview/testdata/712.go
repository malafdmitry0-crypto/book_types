package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x struct{ N int }
	err := json.Unmarshal([]byte(`{"N":4,"extra":8}`), &x)
	fmt.Println(err == nil, x.N)
}
