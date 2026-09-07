package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x struct {
		N int `json:"n,string"`
	}
	err := json.Unmarshal([]byte(`{"n":12}`), &x)
	fmt.Println(err != nil, x.N)
}
