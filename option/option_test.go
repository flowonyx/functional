package option_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/math"
	"github.com/flowonyx/functional/option"
)

func ExampleHandleOption() {
	input := option.Some(1)
	err := option.HandleOption(input, func(i int) error {
		fmt.Printf("%d", i)
		return nil
	}, func() error {
		fmt.Println("None")
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Output: 1
}

func ExampleHandleOptionIgnoreNone() {
	input := option.None[int]()
	err := option.HandleOptionIgnoreNone(input, func(i int) error {
		fmt.Printf("%d", i)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Output:
}

func ExampleMap() {
	f := func(i int) string { return "i:" + strconv.Itoa(i) }
	input := []option.Option[int]{option.None[int](), option.Some(1)}
	r := list.Map(functional.Curry2To1(option.Map[int, string], f), input)
	fmt.Println(r)
	// Output: [None Some(i:1)]
}

func ExampleBind() {
	input := option.Some(2)
	r := option.Bind(func(i int) option.Option[int] { return option.Some(i * 2) }, input)
	fmt.Println(r)
	// Some(4)
}

func TestBind(t *testing.T) {
	r := option.Bind(math.TryParseInt[int], option.None[string]())
	if !r.IsNone() {
		t.Errorf("Binding None should result in None: got %v", r.Value())
	}
	r = option.Bind(math.TryParseInt[int], option.Some("42"))
	if !r.IsSome() || r.Value() != 42 {
		t.Errorf("Binding string(42) should result in Some(42): got %v", r.Value())
	}
	r = option.Bind(math.TryParseInt[int], option.Some("Forty-two"))
	if !r.IsNone() {
		t.Errorf("Binding string(Forty-two) should result in None: got %v", r.Value())
	}
}

func TestContains(t *testing.T) {
	if r := option.Contains(99, option.None[int]()); r {
		t.Error("Contains should return false for None")
	}
	if r := option.Contains(99, option.Some(99)); !r {
		t.Error("Contains should return true for Some(99)")
	}
	if r := option.Contains(99, option.Some(100)); r {
		t.Error("Contains should return false for Some(100)")
	}
}

func TestCount(t *testing.T) {
	if r := option.Count(option.None[int]()); r != 0 {
		t.Error("Count should have returned 0 for None")
	}

	if r := option.Count(option.Some(99)); r != 1 {
		t.Error("Count should have returned 1 for Some(99)")
	}
}

func TestDefaultValue(t *testing.T) {
	if r := option.DefaultValue(99, option.None[int]()); r != 99 {
		t.Error("DefaultValue should return the default when given None")
	}
	if r := option.DefaultValue(99, option.Some(42)); r != 42 {
		t.Error("DefaultValue should return the value when given Some")
	}
}

func TestDefaultWith(t *testing.T) {
	defthunk := func() int { return 99 }
	if r := option.DefaultWith(defthunk, option.None[int]()); r != 99 {
		t.Error("DefaultValue should return the default when given None")
	}
	if r := option.DefaultWith(defthunk, option.Some(42)); r != 42 {
		t.Error("DefaultValue should return the value when given Some")
	}
}

func TestExists(t *testing.T) {
	check := func(x int) bool { return x >= 5 }
	if r := option.Exists(check, option.None[int]()); r {
		t.Error("Exists should return false for None")
	}
	if r := option.Exists(check, option.Some(42)); !r {
		t.Error("Exists should return true for Some(42)")
	}
	if r := option.Exists(check, option.Some(4)); r {
		t.Error("Exists should return false for Some(4)")
	}
}

func TestFilter(t *testing.T) {
	check := func(x int) bool { return x >= 5 }
	if r := option.Filter(check, option.None[int]()); !r.IsNone() {
		t.Error("Filter should return None when given None")
	}
	if r := option.Filter(check, option.Some(42)); r.IsNone() || r.Value() != 42 {
		t.Error("Filter should return Some(42) when given Some(42)")
	}
	if r := option.Filter(check, option.Some(4)); !r.IsNone() {
		t.Error("Filter should return None when given Some(4)")
	}
}

func TestFlatten(t *testing.T) {
	r := option.Flatten(option.None[option.Option[int]]())
	if !r.IsNone() {
		t.Error("Flatten should return None when given None")
	}
	r = option.Flatten(option.Some(option.None[int]()))
	if !r.IsNone() {
		t.Error("Flatten should return None when given Some(None)")
	}
	r = option.Flatten(option.Some(option.Some(42)))
	if r.IsNone() || r.Value() != 42 {
		t.Error("Flatten should return Some(42) when given Some(Some(42))")
	}
}

func TestFold(t *testing.T) {
	folder := func(acc, x int) int { return acc + x*2 }
	r := option.Fold(folder, 0, option.None[int]())
	if r != 0 {
		t.Errorf("Fold should return 0: got %v", r)
	}
	r = option.Fold(folder, 0, option.Some(1))
	if r != 2 {
		t.Errorf("Fold should return 2: got %v", r)
	}
	r = option.Fold(folder, 10, option.Some(1))
	if r != 12 {
		t.Errorf("Fold should return 12: got %v", r)
	}
}

func TestFoldBack(t *testing.T) {
	folder := func(x, acc int) int { return acc + x*2 }
	r := option.FoldBack(folder, option.None[int](), 0)
	if r != 0 {
		t.Errorf("Fold should return 0: got %v", r)
	}
	r = option.FoldBack(folder, option.Some(1), 0)
	if r != 2 {
		t.Errorf("Fold should return 2: got %v", r)
	}
	r = option.FoldBack(folder, option.Some(1), 10)
	if r != 12 {
		t.Errorf("Fold should return 12: got %v", r)
	}
}

func TestForAll(t *testing.T) {
	check := func(x int) bool { return x >= 5 }
	if r := option.ForAll(check, option.None[int]()); !r {
		t.Error("ForAll should return true for None")
	}
	if r := option.ForAll(check, option.Some(42)); !r {
		t.Error("ForAll should return true for Some(42)")
	}
	if r := option.ForAll(check, option.Some(4)); r {
		t.Error("ForAll should return false for Some(4)")
	}
}

func TestMap2(t *testing.T) {
	mapping := func(x, y int) int { return x + y }
	r := option.Map2(mapping, option.None[int](), option.None[int]())
	if !r.IsNone() {
		t.Error("(None, None) should result in None")
	}
	r = option.Map2(mapping, option.Some(5), option.None[int]())
	if !r.IsNone() {
		t.Error("(Some(5), None) should result in None")
	}
	r = option.Map2(mapping, option.None[int](), option.Some(10))
	if !r.IsNone() {
		t.Error("(None, Some(10)) should result in None")
	}
	r = option.Map2(mapping, option.Some(5), option.Some(10))
	if !r.IsSome() || r.Value() != 15 {
		t.Error("(Some(5), Some(10)) should result in Some(15)")
	}
}

func TestMap3(t *testing.T) {
	mapping := func(x, y, z int) int { return x + y + z }
	r := option.Map3(mapping, option.None[int](), option.None[int](), option.None[int]())
	if !r.IsNone() {
		t.Error("(None, None, None) should result in None")
	}
	r = option.Map3(mapping, option.Some(100), option.None[int](), option.None[int]())
	if !r.IsNone() {
		t.Error("(Some(100), None, None) should result in None")
	}
	r = option.Map3(mapping, option.None[int](), option.Some(100), option.None[int]())
	if !r.IsNone() {
		t.Error("(None, Some(100), None) should result in None")
	}
	r = option.Map3(mapping, option.None[int](), option.None[int](), option.Some(100))
	if !r.IsNone() {
		t.Error("(None, None, Some(100)) should result in None")
	}
	r = option.Map3(mapping, option.Some(5), option.Some(100), option.Some(10))
	if !r.IsSome() || r.Value() != 115 {
		t.Error("(Some(5), Some(100), Some(10)) should result in Some(115)")
	}
}

func TestOfNullable(t *testing.T) {
	var i *int
	if !option.OfNullable(i).IsNone() {
		t.Error("OfNullable should be None for nil pointer")
	}
	var j int = 42
	if option.Get(option.OfNullable(&j)) != 42 {
		t.Error("OfNullable should be Some(42)")
	}
}
