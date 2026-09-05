package algorithms_test

import (
	"errors"
	"reflect"
	"testing"

	"gotypes/algorithms"
	"gotypes/fraction"
	"gotypes/money"
)

var errOperation = errors.New("operation failed")

type predicateProbe struct {
	results []bool
	failAt  int
	visited []int
}

func (p *predicateProbe) Len() int { return len(p.results) }
func (p *predicateProbe) Match(i int) (bool, error) {
	p.visited = append(p.visited, i)
	if i == p.failAt {
		return true, errOperation
	}
	return p.results[i], nil
}

func TestPredicateAlgorithms(t *testing.T) {
	cases := []struct {
		name                string
		values              []bool
		index               int
		all, any            bool
		findCalls, allCalls int
	}{
		{"empty", nil, -1, true, false, 0, 0},
		{"all false", []bool{false, false}, -1, false, false, 2, 1},
		{"all true", []bool{true, true}, 0, true, true, 1, 2},
		{"mixed", []bool{false, true, false}, 1, false, true, 2, 1},
		{"later false", []bool{true, false, true}, 0, false, true, 1, 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := &predicateProbe{results: tc.values, failAt: -1}
			index, err := algorithms.Find(p)
			if err != nil || index != tc.index || len(p.visited) != tc.findCalls {
				t.Fatalf("Find: %d, %v, visits %v", index, err, p.visited)
			}
			p.visited = nil
			all, err := algorithms.All(p)
			if err != nil || all != tc.all || len(p.visited) != tc.allCalls {
				t.Fatalf("All: %v, %v, visits %v", all, err, p.visited)
			}
			p.visited = nil
			any, err := algorithms.Any(p)
			if err != nil || any != tc.any || len(p.visited) != tc.findCalls {
				t.Fatalf("Any: %v, %v, visits %v", any, err, p.visited)
			}
		})
	}
}

func TestPredicateErrorsStopTraversal(t *testing.T) {
	p := &predicateProbe{results: []bool{false, true, true}, failAt: 1}
	if index, err := algorithms.Find(p); index != -1 || err != errOperation {
		t.Fatalf("Find: %d,%v", index, err)
	}
	if !reflect.DeepEqual(p.visited, []int{0, 1}) {
		t.Fatal(p.visited)
	}
	p.visited = nil
	if found, err := algorithms.Any(p); found || err != errOperation {
		t.Fatalf("Any: %v,%v", found, err)
	}
	p.results[0] = true
	p.visited = nil
	if all, err := algorithms.All(p); all || err != errOperation {
		t.Fatalf("All: %v,%v", all, err)
	}
	if !reflect.DeepEqual(p.visited, []int{0, 1}) {
		t.Fatal(p.visited)
	}
}

func TestMapStopsAtErrorAndRetainsPrefix(t *testing.T) {
	source := []fraction.Fraction{{Numerator: 2, Denominator: 4}, {}, {Numerator: 3, Denominator: 4}}
	out := []string{"old-0", "old-1", "old-2"}
	calls := 0
	err := algorithms.Map(fractionText{source, out, func(f fraction.Fraction) (string, error) {
		calls++
		if err := f.Validate(); err != nil {
			return "", err
		}
		return f.String(), nil
	}})
	if !errors.Is(err, fraction.ErrZeroDenominator) || calls != 2 {
		t.Fatalf("error %v, calls %d", err, calls)
	}
	if !reflect.DeepEqual(out, []string{"1/2", "old-1", "old-2"}) {
		t.Fatal(out)
	}
	var empty []string
	if err := algorithms.Map(fractionText{Source: nil, Target: empty}); err != nil || empty != nil {
		t.Fatal("empty mapping changed result")
	}
}

type fractionSum struct {
	Values []fraction.Fraction
	Total  fraction.Fraction
}

func (s *fractionSum) Len() int { return len(s.Values) }
func (s *fractionSum) Add(i int) error {
	next, err := s.Total.Add(s.Values[i])
	if err != nil {
		return err
	}
	s.Total = next
	return nil
}

func TestSumDifferentTypesAndErrors(t *testing.T) {
	fractions := fractionSum{
		Values: []fraction.Fraction{{Numerator: 1, Denominator: 2}, {Numerator: 1, Denominator: 3}},
		Total:  fraction.Fraction{Denominator: 1},
	}
	if err := algorithms.Sum(&fractions); err != nil || fractions.Total.String() != "5/6" {
		t.Fatal(fractions.Total, err)
	}
	sums := moneySum{
		Values: []money.Money{{Amount: 100, Currency: "USD"}, {Amount: 200, Currency: "EUR"}, {Amount: 300, Currency: "USD"}},
		Total:  money.Money{Amount: 50, Currency: "USD"},
	}
	if err := algorithms.Sum(&sums); !errors.Is(err, money.ErrCurrencyMismatch) {
		t.Fatal(err)
	}
	if sums.Total.Amount != 150 || sums.Total.Currency != "USD" {
		t.Fatal("partial accumulator", sums.Total)
	}
	sums.Values = nil
	if err := algorithms.Sum(&sums); err != nil || sums.Total.Amount != 150 {
		t.Fatal("empty sum reset the initial value")
	}
}

type item struct {
	key   int
	label string
}
type orderProbe struct {
	items           []item
	calls, failCall int
}

func (p *orderProbe) Len() int { return len(p.items) }
func (p *orderProbe) Compare(i, j int) (int, error) {
	p.calls++
	if p.calls == p.failCall {
		return 0, errOperation
	}
	if p.items[i].key < p.items[j].key {
		return -1, nil
	}
	if p.items[i].key > p.items[j].key {
		return 1, nil
	}
	return 0, nil
}
func (p *orderProbe) Swap(i, j int) { p.items[i], p.items[j] = p.items[j], p.items[i] }

func TestSortIsStable(t *testing.T) {
	p := &orderProbe{items: []item{{2, "a"}, {1, "b"}, {2, "c"}, {1, "d"}, {3, "e"}}}
	if err := algorithms.Sort(p); err != nil {
		t.Fatal(err)
	}
	want := []item{{1, "b"}, {1, "d"}, {2, "a"}, {2, "c"}, {3, "e"}}
	if !reflect.DeepEqual(p.items, want) {
		t.Fatal(p.items)
	}
}

func TestOrderEdgeCases(t *testing.T) {
	for _, items := range [][]item{nil, {{4, "only"}}} {
		p := &orderProbe{items: items}
		want := -1
		if len(items) > 0 {
			want = 0
		}
		if got, err := algorithms.Min(p); err != nil || got != want {
			t.Fatal("Min", got, err)
		}
		if got, err := algorithms.Max(p); err != nil || got != want {
			t.Fatal("Max", got, err)
		}
		if err := algorithms.Sort(p); err != nil || p.calls != 0 {
			t.Fatal("unexpected comparison", p.calls, err)
		}
	}
	p := &orderProbe{items: []item{{3, "a"}, {1, "b"}, {3, "c"}, {1, "d"}}}
	if got, err := algorithms.Min(p); err != nil || got != 1 {
		t.Fatal("first minimum", got, err)
	}
	if got, err := algorithms.Max(p); err != nil || got != 0 {
		t.Fatal("first maximum", got, err)
	}
}

func TestOrderErrors(t *testing.T) {
	for _, run := range []func(algorithms.Ordered) (int, error){algorithms.Min, algorithms.Max} {
		p := &orderProbe{items: []item{{3, "a"}, {2, "b"}, {1, "c"}}, failCall: 1}
		if index, err := run(p); index != -1 || err != errOperation || p.calls != 1 {
			t.Fatal(index, err, p.calls)
		}
	}
	p := &orderProbe{items: []item{{3, "a"}, {2, "b"}, {1, "c"}}, failCall: 2}
	if err := algorithms.Sort(p); err != errOperation || p.calls != 2 {
		t.Fatal(err, p.calls)
	}
	want := []item{{2, "b"}, {3, "a"}, {1, "c"}}
	if !reflect.DeepEqual(p.items, want) {
		t.Fatal("partial ordering", p.items)
	}
}
