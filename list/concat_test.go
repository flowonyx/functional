package list_test

import (
	"fmt"
	"testing"

	"github.com/flowonyx/functional/list"
)

func ExampleConcat() {
	inputs := [][]int{{1, 2}, {3}, {4, 5}}
	r := list.Concat(inputs...)
	fmt.Println(r)
	// Output: [1 2 3 4 5]
}

func TestConcat(t *testing.T) {
	cases := []struct {
		s    [][]int
		want []int
	}{
		{
			s:    [][]int{nil},
			want: nil,
		},
		{
			s:    [][]int{{1}},
			want: []int{1},
		},
		{
			s:    [][]int{{1}, {2}},
			want: []int{1, 2},
		},
		{
			s:    [][]int{{1}, nil, {2}},
			want: []int{1, 2},
		},
	}
	for _, tc := range cases {
		got := list.Concat(tc.s...)
		if !list.Equal(tc.want, got) {
			t.Errorf("Concat(%v) = %v, want %v", tc.s, got, tc.want)
		}
		var sink []int
		allocs := testing.AllocsPerRun(5, func() {
			sink = list.Concat(tc.s...)
		})
		_ = sink
		if allocs > 1 {
			errorf := t.Errorf
			errorf("Concat(%v) allocated %v times; want 1", tc.s, allocs)
		}
	}
}
