package defined

import (
	"fmt"
)

// import "fmt"
type UserID int64

func (id UserID) String() string { return fmt.Sprintf("user:%d", id) }

type OrderID UserID

func exampleIDs() {
	user := UserID(42)
	order := OrderID(user)
	_ = user.String()
	// _ = order.String() // у OrderID нет этого метода
	_ = order
}

type APIUser struct {
	ID int64 `json:"id"`
}
type DBUser struct {
	ID int64 `db:"user_id"`
}

func convertUser(api APIUser) DBUser {
	return DBUser(api) // Go 1.8+: различия тегов не мешают
}

type A int
type B int

type F func(int) int
type G func(int) int

type Vector [2]int
type SliceA []int
type SliceB []int
type MapA map[string]int
type MapB map[string]int

func exampleComposite() {
	a := A(1)
	b := (*B)(&a) // допустимая конверсия неименованных указателей
	*b = 7        // изменяет a: это та же память

	f := F(func(n int) int { return n + 1 })
	g := G(f)
	vector := Vector([2]int{1, 2})

	s := SliceA{1, 2}
	t := SliceB(s) // тот же backing array
	t[0] = 9

	m := MapA{"x": 1}
	convertedMap := MapB(m) // та же map
	convertedMap["x"] = 2

	_, _, _ = g, vector, b
}
