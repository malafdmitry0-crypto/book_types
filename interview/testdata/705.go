package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]int{"old": 1, "x": 2}
	json.Unmarshal([]byte(`{"x":9}`), &m)
	fmt.Println(m["old"], m["x"], len(m))
}
