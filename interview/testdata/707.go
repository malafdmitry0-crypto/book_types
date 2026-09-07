package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	z := 0
	b, _ := json.Marshal(struct {
		A *int `json:"a,omitempty"`
		B *int `json:"b,omitempty"`
	}{nil, &z})
	fmt.Println(string(b))
}
