package cppmemory_test

import "fmt"

type Counter struct{ N int }

func (c *Counter) Inc() { c.N++ }

type Incrementer interface{ Inc() }

func makeNumber() *int {
	x := 42
	return &x
}

func Example_memory() {
	x := 10
	y := x
	p := &x
	*p = 20
	fmt.Println(x, y, *p)
	fmt.Println(*makeNumber())

	c := Counter{N: 10}
	var i Incrementer = &c
	i.Inc()
	fmt.Println(c.N)

	c.N = 10
	var a any = c
	var b any = &c
	c.N = 99
	fmt.Println(a.(Counter).N, b.(*Counter).N)

	var absent *Counter
	var typedNil Incrementer = absent
	var empty Incrementer
	fmt.Println(absent == nil, typedNil == nil, empty == nil)

	const n = 42
	var small int32 = n
	var large int64 = n
	defaultInt := n
	fmt.Printf("%T %T %T\n", small, large, defaultInt)
	// Output:
	// 20 10 20
	// 42
	// 11
	// 10 99
	// true false true
	// int32 int64 int
}
