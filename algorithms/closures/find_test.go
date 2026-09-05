package closures_test

import (
	"errors"
	"reflect"
	"testing"

	"gotypes/algorithms/closures"
)

func TestFindIndex(t *testing.T) {
	stop := errors.New("stop")
	cases := []struct {
		name    string
		results []bool
		failAt  int
		want    int
		visited []int
	}{
		{"empty", nil, -1, -1, nil},
		{"absent", []bool{false, false}, -1, -1, []int{0, 1}},
		{"first match stops", []bool{false, true, true}, -1, 1, []int{0, 1}},
		{"error overrides match", []bool{false, true, true}, 1, -1, []int{0, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var visited []int
			index, err := closures.FindIndex(len(tc.results), func(i int) (bool, error) {
				visited = append(visited, i)
				if i == tc.failAt {
					return tc.results[i], stop
				}
				return tc.results[i], nil
			})
			var wantErr error
			if tc.failAt >= 0 {
				wantErr = stop
			}
			if index != tc.want || !errors.Is(err, wantErr) || !reflect.DeepEqual(visited, tc.visited) {
				t.Fatalf("got (%d, %v), visited %v; want (%d, %v), visited %v", index, err, visited, tc.want, wantErr, tc.visited)
			}
		})
	}
}
