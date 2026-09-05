//go:build go1.24

package modern

import "fmt"

func Example_alias() {
	aliasSetExample()
	var s Set[string] = map[string]bool{"go": true}
	fmt.Println(s["go"])
	// Output: true
}
