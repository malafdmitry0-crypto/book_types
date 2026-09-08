package typing_test

import "fmt"

type Speaker interface {
	Speak() string
}

type Cat struct{}

func (Cat) Speak() string { return "meow" }

type Robot struct{}

func (Robot) Speak() string { return "beep" }

func say(s Speaker) { fmt.Println(s.Speak()) }

func Example_interfaceTypes() {
	var x any = 42
	fmt.Printf("%T\n", x)
	if n, ok := x.(int); ok {
		fmt.Println(n + 1)
	}
	x = "hello"
	fmt.Printf("%T\n", x)
	n, ok := x.(int)
	fmt.Println(n, ok)
	// Output:
	// int
	// 43
	// string
	// 0 false
}

func Example_dispatch() {
	var s Speaker = Cat{}
	say(s)
	s = Robot{}
	say(s)
	// Output:
	// meow
	// beep
}

func Example_interfaceAssertion() {
	var x any = Robot{}
	s, ok := x.(Speaker)
	if ok {
		fmt.Println(s.Speak())
	}
	x = 42
	_, ok = x.(Speaker)
	fmt.Println(ok)
	// Output:
	// beep
	// false
}
