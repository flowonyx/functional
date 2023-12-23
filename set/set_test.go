package set

import (
	"fmt"
	"strconv"
	"strings"
)

func ExampleNewSet() {
	less := func(a, b int) int {
		if a > b {
			return -1
		}
		if b < a {
			return 1
		}
		return 0
	}
	s := NewSet(less)
	s.Add(1)
	s.Add(2)
	s.Add(4)
	s.Add(3)
	fmt.Println(s.Items())
	// Output: [4 3 2 1]
}

func ExampleSingleton() {
	s := Singleton(1)
	fmt.Println(s.Items())
	// Output: [1]
}

func ExampleFromSlice() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]int{1, 3, 4, 2, 4, 3, 2}, func(a, b int) int {
		if a < b {
			return -1
		}
		if b < a {
			return 1
		}
		return 0
	})
	fmt.Println(s.Items(), s2.Items())
	// Output: [one two] [1 2 3 4]
}

func ExampleSet_Equal() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]string{"two", "one", "one"})
	fmt.Println(s.Equal(s2))
	// Output: true
}

func ExampleSet_Contains() {
	s := FromSlice([]string{"one", "two", "one"})
	fmt.Println(s.Contains("two"), s.Contains("three"))
	// Output: true false
}

func ExampleSet_Exists() {
	s := FromSlice([]string{"one", "two", "one"})
	r := s.Exists(func(s string) bool { return strings.HasPrefix(s, "o") })
	r2 := s.Exists(func(s string) bool { return strings.HasPrefix(s, "s") })
	fmt.Println(r, r2)
	// Output: true false
}

func ExampleSet_IsSubsetOf() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]string{"one", "two", "one", "three"})
	fmt.Println(s.IsSubsetOf(s2), s2.IsSubsetOf(s))
	// Output: true false
}

func ExampleSet_IsSupersetOf() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]string{"one", "two", "one", "three"})
	fmt.Println(s.IsSupersetOf(s2), s2.IsSupersetOf(s))
	// Output: false true
}

func ExampleSet_IsProperSubsetOf() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]string{"one", "two", "one", "three"})
	s3 := FromSlice([]string{"one", "two", "one", "three"})
	fmt.Println(s.IsProperSubsetOf(s2), s2.IsProperSubsetOf(s), s2.IsProperSubsetOf(s3))
	// Output: true false false
}

func ExampleSet_IsProperSupersetOf() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]string{"one", "two", "one", "three"})
	s3 := FromSlice([]string{"one", "two", "one", "three"})
	fmt.Println(s.IsProperSupersetOf(s2), s2.IsProperSupersetOf(s), s2.IsProperSupersetOf(s3))
	// Output: false true false
}

func ExampleSet_IndexOf() {
	s := NewSet(func(a, b int) int {
		if a < b {
			return -1
		}
		if b < a {
			return 1
		}
		return 0
	})
	s.Add(1)
	s.Add(3)
	s.Add(2)
	fmt.Println(s.IndexOf(3))
	// Output: 2
}

func ExampleSet_Difference() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]string{"one", "two", "one", "three"})
	r := s.Difference(s2)
	fmt.Println(r.Items())
	// Output: [three]
}

func ExampleSet_Union() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]string{"one", "three"})
	r := s.Union(s2)
	fmt.Println(r.Items())
	// Output: [one two three]
}

func ExampleSet_Intersect() {
	s := FromSlice([]string{"one", "two", "one"})
	s2 := FromSlice([]string{"one", "three"})
	r := s.Intersect(s2)
	fmt.Println(r.Items())
	// Output: [one]
}

func ExampleSet_Filter() {
	s := FromSlice([]int{1, 2, 3, 4})
	r := s.Filter(func(i int) bool { return i%2 == 0 })
	fmt.Println(r.Items())
	// Output: [2 4]
}

func ExampleSet_Iter() {
	s := FromSlice([]int{1, 2, 3, 4})
	s.Iter(func(item int) { fmt.Print(item) })
	// Output: 1234
}

func ExampleSet_Iteri() {
	s := FromSlice([]int{1, 2, 3, 4})
	s.Iteri(func(i, item int) { fmt.Print(i, item) })
	// Output: 0 11 22 33 4
}

func ExampleSet_Partition() {
	s := FromSlice([]int{1, 2, 3, 4})
	t, f := s.Partition(func(i int) bool { return i%2 == 0 })
	fmt.Println(t.Items(), f.Items())
	// Output: [2 4] [1 3]
}

func ExampleDifferenceMany() {
	s := FromSlice([]int{1, 2, 3, 4})
	s2 := FromSlice([]int{1, 2, 3, 4, 5, 6})
	s3 := FromSlice([]int{6, 7})
	r := DifferenceMany(s, s2, s3)
	fmt.Println(r.Items())
	// Output: [5 7]
}

func ExampleUnionMany() {
	s := FromSlice([]int{1, 2, 3, 4})
	s2 := FromSlice([]int{1, 2, 3, 4, 5, 6})
	s3 := FromSlice([]int{6, 7})
	r := UnionMany(s, s2, s3)
	fmt.Println(r.Items())
	// Output: [1 2 3 4 5 6 7]
}

func ExampleIntersectMany() {
	s := FromSlice([]int{1, 2, 3, 4})
	s2 := FromSlice([]int{1, 2, 3, 4, 5, 6})
	s3 := FromSlice([]int{3, 6, 7})
	r := IntersectMany(s, s2, s3)
	fmt.Println(r.Items())
	// Output: [3]
}

func ExampleMap() {
	s := FromSlice([]int{1, 2, 3, 4})
	s2 := Map(func(v int) string { return strconv.Quote(strconv.Itoa(v)) }, s)
	fmt.Println(s2.Items())
	// Output: ["1" "2" "3" "4"]
}

func ExampleMaxElement() {
	s := FromSlice([]int{1, 2, 3, 4})
	r, _ := MaxElement(s)
	fmt.Println(r)
	// Output: 4
}

func ExampleMaxElementBy() {
	s := FromSlice([]int{1, 2, 3, 4})
	r, _ := MaxElementBy(func(v int) int {
		if v < 3 {
			return v * 10
		}
		return v
	}, s)
	fmt.Println(r)
	// Output: 2
}

func ExampleMinElement() {
	s := FromSlice([]int{1, 2, 3, 4})
	r, _ := MinElement(s)
	fmt.Println(r)
	// Output: 1
}

func ExampleMinElementBy() {
	s := FromSlice([]int{1, 2, 3, 4})
	r, _ := MinElementBy(func(v int) int {
		if v < 3 {
			return v * 10
		}
		return v
	}, s)
	fmt.Println(r)
	// Output: 3
}
