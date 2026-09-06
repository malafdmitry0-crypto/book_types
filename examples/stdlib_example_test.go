package algorithms

import (
	"cmp"
	"fmt"
	"maps"
	"math"
	"slices"
	"sort"
)

// Ratio повторяет закон Equal из главы о Fraction: равенство по значению,
// а не по полям. 2/4 и 1/2 — одно число, но разные структуры.
type Ratio struct{ Num, Den int64 }

func (r Ratio) Equal(other Ratio) bool { return r.Num*other.Den == other.Num*r.Den }

func Example_floatOrdering() {
	nan := math.NaN()
	values := []float64{2, nan, 1}
	slices.Sort(values)
	fmt.Println(values)
	legacy := []float64{2, nan, 1}
	sort.Float64s(legacy)
	fmt.Println(legacy)
	fmt.Println(cmp.Compare(nan, 1), cmp.Compare(nan, nan), cmp.Less(nan, math.Inf(-1)))
	fmt.Println(nan < 1, nan == nan)
	// Output:
	// [NaN 1 2]
	// [NaN 1 2]
	// -1 0 true
	// false false
}

func Example_binarySearchContract() {
	sorted := []int{1, 3, 5, 7}
	fmt.Println(slices.BinarySearch(sorted, 5))
	unsorted := []int{5, 1, 7, 3}
	position, found := slices.BinarySearch(unsorted, 5)
	fmt.Println(position, found, slices.Contains(unsorted, 5))
	// Output:
	// 2 true
	// 2 false true
}

func Example_mapKeysOrder() {
	ages := map[string]int{"Вера": 25, "Анна": 17, "Борис": 25}
	fmt.Println(slices.Sorted(maps.Keys(ages)))
	fmt.Println(slices.Sorted(maps.Values(ages)))
	// Output:
	// [Анна Борис Вера]
	// [17 25 25]
}

func Example_containsByValue() {
	values := []Ratio{{1, 3}, {2, 4}}
	half := Ratio{1, 2}
	fmt.Println(slices.Contains(values, half))
	fmt.Println(slices.ContainsFunc(values, half.Equal))
	fmt.Println(slices.Equal(values, []Ratio{{1, 3}, {1, 2}}))
	fmt.Println(slices.EqualFunc(values, []Ratio{{1, 3}, {1, 2}}, Ratio.Equal))
	// Output:
	// false
	// true
	// false
	// true
}

func Example_builtinMinMax() {
	fmt.Println(min(3, 1, 2), max(2.5, 1))
	fmt.Println(max(1.0, math.NaN()), min(math.NaN(), -1.0))
	fmt.Println(max("", "foo", "bar"), min("b", "ab"))
	// Output:
	// 1 2.5
	// NaN NaN
	// foo ab
}
