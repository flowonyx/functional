package simpleMap

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

func ExampleSimpleMap_Iter() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3)}
	m := FromSlice(keyvalues)
	m.Iter(func(key string, value int) {
		if value == 2 {
			fmt.Println(key)
		}
	})
	// Output: 2
}

func ExampleSimpleMap_ToSlice() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3)}
	m := FromSlice(keyvalues)
	r := m.ToSlice()
	fmt.Println(list.Equal(keyvalues, r))
	// Output: true
}

func ExampleSimpleMap_CopyFrom() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3)}
	m := FromSlice(keyvalues)
	m2 := SimpleMap[string, int]{"3": 4, "4": 5}

	r := m.CopyFrom(m2)
	fmt.Println(r["1"], r["2"], r["3"], r["4"])
	// Output: 1 2 4 5
}

func ExampleSimpleMap_RemoveBy() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := m.RemoveBy(func(key string, value int) bool { return key != strings.FromInt(value) })
	fmt.Println(len(r), r.TryGet("3"), r.TryGet("4"))
	// Output: 3 Some(3) None
}

func ExampleSimpleMap_Filter() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := m.Filter(func(key string, value int) bool { return key != strings.FromInt(value) })
	fmt.Println(len(r), r.TryGet("3"), r.TryGet("4"))
	// Output: 1 None Some(5)
}

func ExampleSimpleMap_TryFindKey() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := m.TryFindKey(func(key string, val int) bool { return key != strings.FromInt(val) })
	fmt.Println(r)
	// Output: Some("4")
}

func ExampleSimpleMap_Exists() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := m.Exists(func(key string, val int) bool { return key != strings.FromInt(val) })
	fmt.Println(r)
	// Output: true
}

func ExampleSimpleMap_ForAll() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	r := m.ForAll(func(key string, val int) bool { return key != strings.FromInt(val) })
	fmt.Println(r)
	// Output: false
}

func ExampleSimpleMap_Partition() {
	keyvalues := []Pair[string, int]{PairOf("1", 1), PairOf("2", 2), PairOf("3", 3), PairOf("4", 5)}
	m := FromSlice(keyvalues)
	rt, rf := m.Partition(func(key string, val int) bool { return key == strings.FromInt(val) })
	fmt.Println(rt.Contains("3"), rt.Contains("4"), rf.Contains("4"))
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
