package generic_test

import (
	"fmt"
	"gotypes/lessons/generic"
)

type user struct {
	name string
	age  int
}

func userName(u user) string { return u.name }

func ExampleMap_inference() {
	users := []user{{"Аня", 30}, {"Борис", 17}}
	names := generic.Map(users, userName)              // E = user, R = string выведены
	same := generic.Map[user, string](users, userName) // то же самое явно
	fmt.Println(names, same)
	// Output:
	// [Аня Борис] [Аня Борис]
}

func ExampleZero() {
	// generic.Zero() // ошибка компиляции: cannot infer T
	fmt.Printf("%q %v\n", generic.Zero[string](), generic.Zero[int]())
	// Output:
	// "" 0
}

func ExampleConvert() {
	n := 42
	wide := generic.Convert[int64](n) // To задан явно, From = int выведен
	fmt.Printf("%T %v\n", wide, wide)
	// Output:
	// int64 42
}

func ExampleMax() {
	fmt.Printf("%T %v\n", generic.Max(1, 2), generic.Max(1, 2))
	fmt.Printf("%T %v\n", generic.Max(1, 2.5), generic.Max(1, 2.5)) // Go 1.21+: T = float64
	fmt.Printf("%T %v\n", generic.Max("go", "generic"), generic.Max("go", "generic"))
	// Output:
	// int 2
	// float64 2.5
	// string go
}

func ExampleIdentity() {
	var f func(int) int = generic.Identity // Go 1.21+: T = int из типа переменной
	doubled := generic.Map([]int{1, 2}, generic.Double[int])
	inferred := generic.Map([]int{1, 2}, generic.Double) // Go 1.21+: T выведен из []int
	fmt.Println(f(7), doubled, inferred)
	// Output:
	// 7 [2 4] [2 4]
}

type ages []int

func (a ages) Oldest() int { return generic.Max(generic.Max(a[0], a[1]), a[2]) }

func ExampleClone_inference() {
	copied := generic.Clone(ages{3, 9, 5}) // S = ages из аргумента, E = int из ограничения ~[]E
	fmt.Printf("%T %v %d\n", copied, copied, copied.Oldest())
	// Output:
	// generic_test.ages [3 9 5] 9
}
