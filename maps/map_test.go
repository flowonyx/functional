package maps

import (
	"fmt"

	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/strings"
)

func ExampleFromSlice() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3)}
	m := FromSlice(keyvalues)
	fmt.Println(m["2"])
	// Output: 2
}

func ExampleFromSlices() {
	m := FromSlices([]string{"1", "2", "3"}, []int{1, 2, 3})
	fmt.Println(m["2"])
	// Output: 2
}

func ExampleIter() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3)}
	m := FromSlice(keyvalues)
	Iter(func(key string, value int) {
		if value == 2 {
			fmt.Println(key)
		}
	}, m)
	// Output: 2
}

func ExampleToSlice() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3)}
	m := FromSlice(keyvalues)
	r := ToSlice(m)
	fmt.Println(list.Equal(keyvalues, r))
	// Output: true
}

func ExampleCopyFrom() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3)}
	m := FromSlice(keyvalues)
	m2 := map[string]int{"3": 4, "4": 5}

	r := CopyFrom(m2, m)
	fmt.Println(r["1"], r["2"], r["3"], r["4"])
	// Output: 1 2 4 5
}

func ExampleRemoveBy() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := RemoveBy(func(key string, value int) bool { return key != strings.FromInt(value) }, m)
	fmt.Println(len(r), TryGet("3", r), TryGet("4", r))
	// Output: 3 Some(3) None
}

func ExampleFilter() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := Filter(func(key string, value int) bool { return key != strings.FromInt(value) }, m)
	fmt.Println(len(r), TryGet("3", r), TryGet("4", r))
	// Output: 1 None Some(5)
}

func ExampleTryFindKey() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := TryFindKey(func(key string, val int) bool { return key != strings.FromInt(val) }, m)
	fmt.Println(r)
	// Output: Some("4")
}

func ExampleExists() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := Exists(func(key string, val int) bool { return key != strings.FromInt(val) }, m)
	fmt.Println(r)
	// Output: true
}

func ExampleForAll() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := ForAll(func(key string, val int) bool { return key != strings.FromInt(val) }, m)
	fmt.Println(r)
	// Output: false
}

func ExamplePartition() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	rt, rf := Partition(func(key string, val int) bool { return key == strings.FromInt(val) }, m)
	fmt.Println(Contains("3", rt), Contains("4", rt), Contains("4", rf))
	// Output: true false true
}

func ExampleFoldMap() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := FoldMap(func(state int, key string, val int) int {
		i := strings.ToIntOpt(key)
		// ((((0 * 1 + 1) * 2 + 2) * 3 + 3) * 4 + 5) = 65
		return state*i.Value() + val
	}, 0, m)
	fmt.Println(r)
	// Output: 65
}

func ExampleFoldBackMap() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := FoldBackMap(func(key string, val int, state int) int {
		i := strings.ToIntOpt(key)
		// ((((0 * 4 + 5) * 3 + 3) * 2 + 2) * 1 + 1) = 39
		return state*i.Value() + val
	}, 0, m)
	fmt.Println(r)
	// Output: 39
}

func ExampleMapTo() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := MapTo(func(key string, val int) string {
		return key
	}, m)
	fmt.Println(r["4"])
	// Output: 4
}
