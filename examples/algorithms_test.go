package algorithms

import (
	"reflect"
	"slices"
	"testing"
)

func boxedUsers(src []User) []interface{} {
	if src == nil {
		return nil
	}
	out := make([]interface{}, len(src))
	for i, item := range src {
		out[i] = item
	}
	return out
}

func TestSharedContract(t *testing.T) {
	cases := []struct {
		name  string
		users []User
		index int
		all   bool
		names []string
	}{
		{"nil", nil, -1, true, nil},
		{"empty", []User{}, -1, true, []string{}},
		{"mixed", []User{{"Аня", 17, true}, {"Борис", 31, true}, {"Вера", 25, false}}, 1, false, []string{"Аня", "Борис", "Вера"}},
		{"no match", []User{{"Аня", 17, true}}, -1, true, []string{"Аня"}},
		{"zero value", []User{{}}, -1, false, []string{""}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			adult := func(u User) bool { return u.Age >= 18 }
			active := func(u User) bool { return u.Active }
			name := func(u User) string { return u.Name }
			boxed := boxedUsers(tc.users)
			finds := map[string]int{
				"concrete":  FindUser(tc.users, adult),
				"closure":   FindIndex(len(tc.users), func(i int) bool { return adult(tc.users[i]) }),
				"interface": FindInterface(UserQuery{tc.users, adult}),
				"any":       FindAny(boxed, func(v interface{}) bool { return adult(v.(User)) }),
				"generated": FindUsersGenerated(tc.users, adult),
				"generic":   Find(tc.users, adult),
			}
			idx, err := FindReflect(tc.users, adult)
			if err != nil {
				t.Fatal(err)
			}
			finds["reflect"] = idx
			for variant, got := range finds {
				if got != tc.index {
					t.Errorf("Find %s: got %d, want %d", variant, got, tc.index)
				}
			}
			alls := map[string]bool{
				"concrete":  AllUsers(tc.users, active),
				"closure":   AllIndex(len(tc.users), func(i int) bool { return active(tc.users[i]) }),
				"interface": AllInterface(UserQuery{tc.users, active}),
				"any":       AllAny(boxed, func(v interface{}) bool { return active(v.(User)) }),
				"generated": AllUsersGenerated(tc.users, active),
				"generic":   All(tc.users, active),
				"sequence":  AllSeq(slices.Values(tc.users), active),
			}
			all, err := AllReflect(tc.users, active)
			if err != nil {
				t.Fatal(err)
			}
			alls["reflect"] = all
			for variant, got := range alls {
				if got != tc.all {
					t.Errorf("All %s: got %v, want %v", variant, got, tc.all)
				}
			}
			mapped := map[string][]string{"concrete": MapUserNames(tc.users, name), "generated": MapUsersGenerated(tc.users, name), "generic": Map(tc.users, name)}
			var closure, adapter []string
			if tc.users != nil {
				closure = make([]string, len(tc.users))
				adapter = make([]string, len(tc.users))
			}
			MapInto(len(tc.users), func(i int) { closure[i] = name(tc.users[i]) })
			mapped["closure"] = closure
			RunMap(UserNameJob{tc.users, adapter, name})
			mapped["interface"] = adapter
			raw := MapAny(boxed, func(v interface{}) interface{} { return name(v.(User)) })
			var names []string
			if raw != nil {
				names = make([]string, len(raw))
			}
			for i, v := range raw {
				names[i] = v.(string)
			}
			mapped["any"] = names
			result, err := MapReflect(tc.users, name)
			if err != nil {
				t.Fatal(err)
			}
			mapped["reflect"] = result.([]string)
			for variant, got := range mapped {
				if !reflect.DeepEqual(got, tc.names) {
					t.Errorf("Map %s: got %#v, want %#v", variant, got, tc.names)
				}
			}
			// Collect intentionally has a different empty-result contract.
			if got := slices.Collect(MapSeq(slices.Values(tc.users), name)); !slices.Equal(got, tc.names) {
				t.Errorf("MapSeq: %v", got)
			}
			found, ok := FindSeq(slices.Values(tc.users), adult)
			if ok != (tc.index >= 0) {
				t.Fatalf("FindSeq found=%v", ok)
			}
			if ok && found != tc.users[tc.index] {
				t.Errorf("FindSeq returned %v", found)
			}
		})
	}
}

func TestShortCircuitAcrossApproaches(t *testing.T) {
	users := []User{{Age: 1}, {Age: 2}, {Age: 3}, {Age: 4}}
	variants := map[string]struct {
		find func(func(User) bool) int
		all  func(func(User) bool) bool
	}{
		"concrete":  {func(p func(User) bool) int { return FindUser(users, p) }, func(p func(User) bool) bool { return AllUsers(users, p) }},
		"closure":   {func(p func(User) bool) int { return FindIndex(len(users), func(i int) bool { return p(users[i]) }) }, func(p func(User) bool) bool { return AllIndex(len(users), func(i int) bool { return p(users[i]) }) }},
		"interface": {func(p func(User) bool) int { return FindInterface(UserQuery{users, p}) }, func(p func(User) bool) bool { return AllInterface(UserQuery{users, p}) }},
		"any": {func(p func(User) bool) int {
			return FindAny(boxedUsers(users), func(v interface{}) bool { return p(v.(User)) })
		}, func(p func(User) bool) bool {
			return AllAny(boxedUsers(users), func(v interface{}) bool { return p(v.(User)) })
		}},
		"reflect": {func(p func(User) bool) int {
			v, e := FindReflect(users, p)
			if e != nil {
				t.Fatal(e)
			}
			return v
		}, func(p func(User) bool) bool {
			v, e := AllReflect(users, p)
			if e != nil {
				t.Fatal(e)
			}
			return v
		}},
		"generated": {func(p func(User) bool) int { return FindUsersGenerated(users, p) }, func(p func(User) bool) bool { return AllUsersGenerated(users, p) }},
		"generic":   {func(p func(User) bool) int { return Find(users, p) }, func(p func(User) bool) bool { return All(users, p) }},
	}
	for name, v := range variants {
		t.Run(name, func(t *testing.T) {
			calls := 0
			got := v.find(func(u User) bool { calls++; return u.Age == 2 })
			if got != 1 || calls != 2 {
				t.Fatalf("Find: index=%d calls=%d", got, calls)
			}
			calls = 0
			all := v.all(func(u User) bool { calls++; return u.Age != 2 })
			if all || calls != 2 {
				t.Fatalf("All: result=%v calls=%d", all, calls)
			}
		})
	}
}

func TestReflectionRejectsInvalidAPIs(t *testing.T) {
	var nilFn func(User) string
	cases := []struct{ src, fn interface{} }{
		{nil, func(User) string { return "" }},
		{123, func(User) string { return "" }},
		{[]User{}, nil},
		{[]User{}, nilFn},
		{[]User{}, 17},
		{[]User{}, func(string) string { return "" }},
		{[]User{}, func(User, User) string { return "" }},
		{[]User{}, func(User) {}},
		{[]User{}, func(...User) string { return "" }},
	}
	for i, tc := range cases {
		if _, err := MapReflect(tc.src, tc.fn); err == nil {
			t.Errorf("case %d accepted invalid API", i)
		}
	}
	if _, err := FindReflect([]User{}, func(User) int { return 1 }); err == nil {
		t.Error("non-bool predicate accepted")
	}
	if _, err := AllReflect([]User{}, func(User) int { return 1 }); err == nil {
		t.Error("non-bool predicate accepted")
	}
}

func TestCallbackPanicPropagates(t *testing.T) {
	defer func() {
		if got := recover(); got != "callback failure" {
			t.Errorf("panic=%v", got)
		}
	}()
	_, _ = MapReflect([]User{{}}, func(User) string { panic("callback failure") })
}

func TestSequenceIsLazyAndStops(t *testing.T) {
	produced, transformed := 0, 0
	source := func(yield func(int) bool) {
		for i := 1; i <= 100; i++ {
			produced++
			if !yield(i) {
				return
			}
		}
	}
	mapped := MapSeq(source, func(v int) int { transformed++; return v * 2 })
	if produced != 0 || transformed != 0 {
		t.Fatal("MapSeq ran before consumption")
	}
	if AllSeq(mapped, func(v int) bool { return v < 6 }) {
		t.Fatal("AllSeq accepted 6")
	}
	if produced != 3 || transformed != 3 {
		t.Fatalf("produced=%d transformed=%d", produced, transformed)
	}
	got, ok := FindSeq(slices.Values([]int{0, 1}), func(v int) bool { return v == 0 })
	if got != 0 || !ok {
		t.Fatalf("zero value match: (%d,%v)", got, ok)
	}
}

func TestDifferentElementAndResultTypes(t *testing.T) {
	got := Map([]int{1, 2}, func(v int) bool { return v%2 == 0 })
	if !slices.Equal(got, []bool{false, true}) {
		t.Fatal(got)
	}
	if Find([][]int{{1}, {2, 3}}, func(v []int) bool { return len(v) == 2 }) != 1 {
		t.Fatal("Find incorrectly requires comparable")
	}
	type IDs []int
	original := IDs{1, 2}
	cloned := Clone(original)
	var _ IDs = cloned
	cloned[0] = 9
	if original[0] != 1 {
		t.Fatal("Clone shares element storage")
	}
}
