package defined

import (
	"fmt"
	"testing"
)

func Example() {
	exampleIDs()
	exampleComposite()
	fmt.Println(UserID(42), convertUser(APIUser{ID: 42}))
	// Output: user:42 {42}
}
func TestSharedStorage(t *testing.T) {
	a := A(1)
	b := (*B)(&a)
	*b = 7
	if a != 7 {
		t.Fatal(a)
	}
	s := SliceA{1, 2}
	converted := SliceB(s)
	converted[0] = 9
	if s[0] != 9 {
		t.Fatal(s)
	}
	m := MapA{"x": 1}
	MapB(m)["x"] = 2
	if m["x"] != 2 {
		t.Fatal(m)
	}
	if G(F(func(n int) int { return n + 1 }))(2) != 3 {
		t.Fatal("function conversion")
	}
}
