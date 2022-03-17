package orderedMap

import (
	"fmt"
	"strings"

	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

func ExampleOrderedMap_Set() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(1, "one")
	m.Set(2, "two")
	fmt.Println(m.Keys(), m.Values())
	// Output: [1 2] [one two]
}

func ExampleFromSlice() {
	m := FromSlice([]functional.Pair[string, int]{
		functional.PairOf("one", 1),
		functional.PairOf("two", 2),
		functional.PairOf("two", 3),
	})

	fmt.Println(m.Keys(), m.Values())
	// Output: [one two] [1 3]
}

func ExampleEqual() {
	m := NewOrderedMap[int, string]()
	m.Set(2, "two")
	m.Set(1, "one")
	m2 := NewOrderedMap[int, string]()
	m2.Set(1, "one")
	m2.Set(2, "two")
	m3 := NewOrderedMap[int, string]()
	m3.Set(1, "one")
	m3.Set(2, "two")
	m3.Set(3, "three")
	m4 := NewOrderedMap[int, string]()
	m4.Set(1, "One")
	m4.Set(2, "Two")
	fmt.Println(Equal(m, m2), Equal(m, m3), Equal(m, m4))
	// Output: true false false
}

func ExampleEqualBy() {
	type vt struct {
		Value string
	}
	m := NewOrderedMap[int, vt]()
	m.Set(1, vt{Value: "one"})
	m.Set(2, vt{Value: "two"})
	m.Set(3, vt{Value: "three"})

	m2 := NewOrderedMap[int, vt]()
	m2.Set(1, vt{Value: "one"})
	m2.Set(2, vt{Value: "two"})
	m2.Set(3, vt{Value: "three"})

	m3 := NewOrderedMap[int, vt]()
	m3.Set(1, vt{Value: "One"})
	m3.Set(2, vt{Value: "Two"})
	m3.Set(3, vt{Value: "Three"})

	pred := func(v1 vt, v2 vt) bool {
		return v1.Value == v2.Value
	}

	fmt.Println(EqualBy(pred, m, m2), EqualBy(pred, m, m3))
	// Output: true false
}

func ExampleOrderedMap_Get() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	fmt.Println(m.Get(1))
	// Output: one
}

func ExampleOrderedMap_Remove() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	m.Remove(1)

	fmt.Println(m.TryGet(1))
	// Output: None
}

func ExampleOrderedMap_Filter() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "Three")

	r := m.Filter(func(i int, s string) bool { return strings.HasPrefix(s, "T") })
	fmt.Println(r.Keys())
	// Output: [3]
}

func ExampleOrderedMap_Find() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	v, err := m.Find(1)
	if err != nil {
		panic(err)
	}

	_, err = m.Find(3)

	fmt.Println(v, errors.Is(err, errors.KeyNotFoundErr))
	// Output: one true
}

func ExampleOrderedMap_FindKey() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	v, err := m.FindKey(func(key int, value string) bool { return strings.HasPrefix(value, "t") })
	if err != nil {
		panic(err)
	}

	_, err = m.FindKey(func(key int, value string) bool { return strings.HasPrefix(value, "T") })

	fmt.Println(v, errors.Is(err, errors.KeyNotFoundErr))
	// Output: 2 true
}

func ExampleOrderedMap_TryFindKey() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	v := m.TryFindKey(func(_ int, value string) bool { return strings.HasPrefix(value, "t") })
	v2 := m.TryFindKey(func(_ int, value string) bool { return strings.HasPrefix(value, "T") })

	fmt.Println(v, v2)
	// Output: Some(2) None
}

func ExampleOrderedMap_ForAll() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	lessThan3 := func(key int, _ string) bool { return key < 3 }

	r1 := m.ForAll(lessThan3)
	m.Set(3, "three")
	r2 := m.ForAll(lessThan3)
	fmt.Println(r1, r2)
	// Output: true false
}

func ExampleOrderedMap_Iter() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	m.Iter(func(key int, value string) {
		fmt.Print(key, ":", value, ",")
	})

	// Output: 1:one,2:two,
}

func ExampleOrderedMap_Iteri() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	m.Iteri(func(i int, key int, value string) {
		fmt.Print(i, ":", key, ":", value, ",")
	})

	// Output: 0:1:one,1:2:two,
}

func ExampleOrderedMap_Keys() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	fmt.Println(m.Keys())
	// Output: [1 2]
}

func ExampleOrderedMap_Values() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	fmt.Println(m.Values())
	// Output: [one two]
}

func ExampleOrderedMap_Partition() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	t, f := m.Partition(func(i int, _ string) bool {
		return i < 3
	})

	fmt.Println(t.Contains(3), f.Contains(3))
	// Output: false true
}

func ExampleFold() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	r := Fold(func(state int, key int, value string) int {
		// (((((1 * 1) + 1) * 2) +1) * 3) = 15
		return (state + 1) * key
	}, 0, m)

	fmt.Println(r)
	// Output: 15
}

func ExampleFoldBack() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	r := FoldBack(func(key int, value string, state int) int {
		// (((((1 * 3) + 1) * 2) +1) * 1) = 9
		return (state + 1) * key
	}, m, 0)

	fmt.Println(r)
	// Output: 9
}

func ExampleMapTo() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	r := MapTo(func(key int, value string) (string, int) {
		return value, key
	}, m)

	fmt.Println(r.Keys())
	// Output: [one two three]
}

func ExampleMapValuesTo() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	r := MapValuesTo(func(key int, value string) int {
		return key + 1
	}, m)

	fmt.Println(r.Get(1))
	// Output: 2
}

func ExamplePick() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	r, err := Pick(func(key int, _ string) option.Option[float64] {
		if key > 1 {
			return option.Some(float64(key))
		}
		return option.None[float64]()
	}, m)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%.2f", r)
	// Output: 2.00
}

func ExampleTryPick() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	r := TryPick(func(key int, _ string) option.Option[float64] {
		if key > 1 {
			return option.Some(float64(key))
		}
		return option.None[float64]()
	}, m)

	fmt.Println(r)
	// Output: Some(2)
}

func ExampleSet() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	r := Set(m, 2, "Two")
	fmt.Println(r.Values())
	// Output: [one Two three]
}

func ExampleRemove() {
	m := NewOrderedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Set(3, "three")

	r := Remove(m, 2)
	fmt.Println(r.Keys())
	// Output: [1 3]
}
