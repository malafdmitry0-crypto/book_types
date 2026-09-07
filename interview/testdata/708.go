package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var p *int
	b, _ := json.Marshal(struct {
		X any `json:"x,omitempty"`
	}{p})
	fmt.Println(string(b))
}
